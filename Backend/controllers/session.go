package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Puppylove-IITK/puppylove/config"
	"github.com/Puppylove-IITK/puppylove/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginInfo struct {
	Username string `json:"username" xml:"username" form:"username"`
	Passhash string `json:"password" xml:"password" form:"password"`
}

func SessionLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("Status") != nil {
		session.Clear()
		session.Save()
	}

	u := new(LoginInfo)
	if err := c.BindJSON(u); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
		return
	}

	// @TODO @IMPORTANT Move password to env variable
	if u.Username == "admin" {
		if u.Passhash == config.CfgAdminPass {
			session.Set("Status", "login")
			session.Set("id", u.Username)
			session.Save()
			c.String(http.StatusOK,
				fmt.Sprintf("Logged in: %s", u.Username))
		} else {
			SessionLogout(c)
			c.String(http.StatusOK, "Invalid username or password")
		}
		return
	}

	// Fetch user
	var user models.User
	if err := Db.GetCollection("user").FindOne(context.Background(), bson.M{"_id": u.Username}).Decode(&user); err != nil {
		SessionLogout(c)
		c.String(http.StatusNotFound, "Invalid user")
		log.Println("Invalid user: " + u.Username)
		return
	}
	// Check login
	if user.Pass == u.Passhash {
		session.Set("Status", "login")
		session.Set("id", u.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
		})
	} else {
		SessionLogout(c)
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func SessionLogout(c *gin.Context) {
	sessions.Default(c).Clear()
	sessions.Default(c).Save()
}

func SessionId(c *gin.Context) (string, error) {
	id := sessions.Default(c).Get("id")
	if id != nil {
		return id.(string), nil
	}
	return "", errors.New("No such session")
}
