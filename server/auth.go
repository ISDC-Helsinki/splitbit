package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"isdc-helsinki.fi/splitbit/server/models"
)

type loginDTO struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// User demo
type User struct {
	id       int
	username string
}

func createAuthMiddleware() *jwt.GinJWTMiddleware {

	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginDTO
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
                        user, err := models.Members(models.MemberWhere.Username.EQ(loginVals.Username),
				models.MemberWhere.Password.EQ(loginVals.Password),
			).OneG(c)
                        if err != nil {
                                return nil, jwt.ErrFailedAuthentication
                        }
			return user, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Member); ok {
                                println(v.Username)
				return jwt.MapClaims{
					identityKey: v.ID,
					"username": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
                LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
                },
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return int64(claims[identityKey].(float64))
                },
		Authorizator: func(data interface{}, c *gin.Context) bool {
                        return true
                        if v, ok := data.(int64); ok && v == 1 {
                                return true
                        }
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("Unable to initialize JWT Auth Middleware")
	}
	return middleware

}
func AuthPing(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"text":     "Pong",
		"userID":   claims[identityKey],
		"userName": user.(*User).id,
	})
}
