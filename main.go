package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Users []User = []User{}

func handleusers(c *gin.Context) {
	var user = User{}

	b, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(b, &user)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("user", user)
	Users = append(Users, user)
	fmt.Println("users", Users)

	c.JSON(http.StatusCreated, user)

}

func GetUser(c *gin.Context) {
	u := User{}
	id := c.Param("id")

	for i := 0; i < len(Users); i++ {
		if id == Users[i].Id {
			u = Users[i]
			break
		}
	}

	if u.Id == "" && u.Name == "" {
		c.JSON(http.StatusNotFound, "")
		return
	}

	c.JSON(http.StatusOK, u)
}

func main() {
	fmt.Println("server")

	r.POST("/users", handleusers)

	r.GET("/users/:id", GetUser)

	r.Run()
}
