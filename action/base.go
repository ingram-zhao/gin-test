package action

import (
	"log"
	"time"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "Your url path: /someGet",
	})
}

func Gettingname(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func Gettingmsg(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

// curl -X POST -d "message=hello_world" http://127.0.0.1:8080/form_post
func Form_post(c *gin.Context) {
	// message := c.PostForm("message")
	message := c.PostFormMap("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	fmt.Printf("message: %s; nick: %s", message, nick)
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

// curl -X POST -d "name=zhangsan&message=hello" http://127.0.0.1:8080/post?id=1&page=1
func Postting(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}


type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

// curl -X POST -H "Content-Type:application/json" -d '{"name":"zhangsan","address":"shanghai"}'  http://127.0.0.1:8080/testing
func StartPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}

	// https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292
	if c.BindJSON(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
	}

	// c.String(200, "Success")
	c.JSON(200, gin.H{
		"name": person.Name,
		"address": person.Address,
	})
}