package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"isdc-helsinki.fi/splitbit/server/data"
	// "io"
	// "github.com/otiai10/gosseract/v2"
)

//go:embed schema.sql
var ddl string

func main() {

	db := setupDB()

	qs := data.New(db) // qs stands for queries

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow cookies to be sent
		MaxAge:           12 * time.Hour, // How long the results of a preflight request can be cached
	}))
	authMiddleware := createAuthMiddleware(qs)
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
		id, err := qs.AddUser(c, data.AddUserParams{Username: dto.Username, Displayname: dto.Username, Password: dto.Password})

		if err != nil {
			println(err)
			c.JSON(400, gin.H{"error": "Error inserting user to db"})
			return
		}
		c.JSON(200, id)
	})

	// Grouping routes that now use authentication
	a := r.Group("/", authMiddleware.MiddlewareFunc())

	a.GET("/groups", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		g, _ := qs.GetGroupsOfMember(c, claims["id"].(int64))
		c.JSON(http.StatusOK, g)
	})

	r.GET("/groups-nonauthed", func(c *gin.Context) {
		fmt.Print(c.Cookie("jwt"))
		g, err := qs.GetGroupsAll(c)
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

		g, _ := qs.AddGroup(c, dto.Name)
		c.JSON(http.StatusOK, g)
	})

	r.GET("/groups/:id/items", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		g, _ := qs.GetItemsOfGroup(c, int64(id))

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

		item_id, err := qs.AddItemToGroup(c, data.AddItemToGroupParams{
			Name:      dto.Name,
			Timestamp: dto.Timestamp,
			Price:     dto.Price,
			GroupID:   int64(id),
			AuthorID:  dto.Member_id,
		})

		if err != nil {
			print(err)
			c.JSON(400, gin.H{"error": "Error inserting to db"})
		}
		c.JSON(200, item_id)
	})

	r.GET("/groups/:id/members", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		m, err := qs.GetMembersOfGroup(c, int64(id))
		if err != nil {
			println(err)
		}
		c.JSON(200, m)
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
		err := qs.AddMemberToGroup(c, data.AddMemberToGroupParams{GroupID: int64(id), MemberID: int64(dto.MemberID)})
		if err != nil {
			c.JSON(404, gin.H{"error": "Could not add member"})
			return
		}
		c.Status(200)
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
