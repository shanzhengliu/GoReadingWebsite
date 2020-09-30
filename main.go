package main

import (
	"ToolWebsite/router"
	"github.com/gin-gonic/gin"

	//"fmt"
	//"net/http"
	//"os"

)

func main()  {

	r :=gin.Default()
	r.GET("/bookclass",router.GetBookClass)
	r.GET("/classbooklist",router.GetClassBookList)
	r.GET("/chartherlist",router.GetCharther)
	r.GET("/bookcontent",router.GetBookContent)
	r.Run(":10000")

}

