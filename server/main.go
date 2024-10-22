package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"isdc-helsinki.fi/splitbit/server/.gen/model"
	. "isdc-helsinki.fi/splitbit/server/.gen/table"
	"isdc-helsinki.fi/splitbit/server/models"
	// "io"
	// "github.com/otiai10/gosseract/v2"
)

func main() {

	db := setupDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow cookies to be sent
		MaxAge:           12 * time.Hour, // How long the results of a preflight request can be cached
	}))
	authMiddleware := createAuthMiddleware()
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", func(c *gin.Context) {
		var dto struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		member := models.Member{
			Username:    dto.Username,
			DisplayName: dto.Username,
			Password:    dto.Password,
		}
		err := member.InsertG(c, boil.Infer())

		if err != nil {
			println(err)
			c.JSON(400, gin.H{"error": "Error inserting user to db"})
			return
		}
		c.JSON(200, member.ID)
	})

	// Grouping routes that now use authentication
	a := r.Group("/", authMiddleware.MiddlewareFunc())

	a.GET("/groups", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		var dest []model.Groups
		err := SELECT(Groups.AllColumns).
			FROM(Groups.
				LEFT_JOIN(MemberGroups, Groups.ID.EQ(MemberGroups.GroupID))).
			WHERE(MemberGroups.MemberID.EQ(Int(int64(claims["id"].(float64))))).Query(db, &dest)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching groups from db"})
			return
		}
		c.JSON(http.StatusOK, dest)
	})

	r.GET("/groups-nonauthed", func(c *gin.Context) {
		fmt.Print(c.Cookie("jwt"))
		var dest []model.Groups
		err := SELECT(Groups.AllColumns).FROM(Groups).Query(db, &dest)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching groups from db"})
			return
		}
		c.JSON(http.StatusOK, dest)
	})

	r.POST("/groups", func(c *gin.Context) {
		var dto struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		var group_id []int64

		err := Groups.INSERT(Groups.Name).VALUES(dto.Name).RETURNING(Groups.ID).Query(db, &group_id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching groups from db"})
			return
		}
		c.JSON(200, group_id[0])

	})

	r.GET("/groups/:id/items", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var dest []model.Items
		err := SELECT(Items.AllColumns).WHERE(Items.GroupID.EQ(Int(int64(id)))).ORDER_BY(Items.Timestamp.DESC()).FROM(Items).Query(db, &dest)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": "Error fetching items from db"})
			return
		}
		c.JSON(http.StatusOK, dest)
	})

	r.POST("/groups/:id/items", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var dto struct {
			Name      string  `json:"name"`
			Timestamp int64   `json:"timestamp"`
			Price     float64 `json:"price"`
			Member_id int64   `json:"member_id"`
		}

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		item := model.Items{
			Name:      dto.Name,
			Timestamp: int32(dto.Timestamp),
			Price:     float32(dto.Price),
			GroupID:   int32(id),
			AuthorID:  int32(dto.Member_id),
		}

		var item_id []int64 // for whatever reason this has to be a slice
		err := Items.INSERT(Items.AllColumns.Except(Items.ID)).MODEL(item).RETURNING(Items.ID).Query(db, &item_id)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"error": "Error inserting to db"})
			return
		}
		c.JSON(200, item_id)
	})

	r.GET("/groups/:id/members", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		m, err := models.Members(models.MemberWhere.ID.EQ(int64(id)), Load(models.MemberRels.Groups)).OneG(c) // eager loading
		if err != nil {
			println(err)
		}
		c.JSON(200, m.R.GetGroups())
	})

	r.POST("/groups/:id/members", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var dto struct {
			MemberID int `json:"member_id"`
		}
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		member, err := models.FindMemberG(c, int64(dto.MemberID))
		if err != nil {
			c.JSON(404, gin.H{"error": "Member not found"})
			return
		}
		group, err := models.FindGroupG(c, int64(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Group to add the user to not found"})
			return
		}
		member.AddGroupsG(c, false, group)
	})

	// Demo of the ocr functionality
	// r.POST("/receipt", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)
	// 	file_handle, _ := file.Open()
	// 	file_bytes, _ := io.ReadAll(file_handle)
	// 	client := gosseract.NewClient()
	// 	defer client.Close()
	// 	client.SetImageFromBytes(file_bytes)
	// 	text, _ := client.Text()
	// 	c.JSON(200, text)
	// })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
