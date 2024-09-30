package main

import (
	"log"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"isdc-helsinki.fi/splitbit/server/models"
	// "io"
	// "github.com/otiai10/gosseract/v2"
)

func main() {

	setupDB()
	r := gin.Default()
	r.Use(cors.Default())

	authMiddleware := createAuthMiddleware()
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.GET("/meow", func(c *gin.Context) {
		c.String(200, "oink oink oink")
	})

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
		g, _ := models.Groups(models.MemberWhere.ID.EQ(claims["id"].(int64))).AllG(c)
		c.JSON(http.StatusOK, g)
	})

	r.GET("/groups-nonauthed", func(c *gin.Context) {
		g, err := models.Groups().AllG(c)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching groups from db"})
		}
		c.JSON(http.StatusOK, g)
	})

	r.POST("/groups", func(c *gin.Context) {
		var dto struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		group := models.Group{
			Name: dto.Name,
		}
		group.InsertG(c, boil.Infer())
	})

	r.GET("/groups/:id/items", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		g, _ := models.Items(models.ItemWhere.GroupID.EQ(int64(id)), OrderBy(models.ItemColumns.Timestamp+" desc")).AllG(c)
		c.JSON(http.StatusOK, g)
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

		item := models.Item{
			Name:      dto.Name,
			Timestamp: dto.Timestamp,
			Price:     dto.Price,
			GroupID:   int64(id),
			AuthorID:  dto.Member_id,
		}

		if err := item.InsertG(c, boil.Infer()); err != nil {
			print(err)
			c.JSON(400, gin.H{"error": "Error inserting to db"})
		}
		c.JSON(200, item.ID)
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
