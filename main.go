package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ingramzhao/gin-test/action"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// This handler will match /someGet
	// router.GET("/someGet", action.Getting)

	// This handler will match /user/xx but will not match /user/ or /user
	router.GET("/someGet/:name",action.Gettingname)
	
	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/someGet/:name/*action", action.Gettingmsg)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", action.Welcome)

	// Multipart/Urlencoded Form
	router.POST("/form_post", action.Form_post)

	//query + post form
	router.POST("/post", action.Postting)

	// Bind Query String or Post Data
	router.Any("/testing", action.StartPage)
	


	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}