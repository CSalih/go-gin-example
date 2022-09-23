package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r routes) addUsers(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	authorizedUsers := users.Group("/", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))

	users.GET("/:name", getUserByUsername)
	authorizedUsers.PUT("/:name", updateUserStatus)
}

var db = make(map[string]string)

func getUserByUsername(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func updateUserStatus(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Status string `json:"status" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Status
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
