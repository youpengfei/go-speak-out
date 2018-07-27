package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"io"
	"github.com/qiniu/api.v7/auth/qbox"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()

	route.LoadHTMLGlob("templates/**/*")
	route.GET("/", toIndex)
	route.GET("/manager", toManager)
	route.GET("/api/users", getAllUsers)
	route.POST("/files", updateImages)
	route.Run(":8888")
}

func getAllUsers(c *gin.Context) {

}

func toIndex(c *gin.Context) {
	log.Print("get index")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "尤鹏飞的博客",
	})
}

func toManager(c *gin.Context) {
	log.Print("to manager")
	c.HTML(http.StatusOK, "manager/index.tmpl", gin.H{
		"title": "博客后台",
	})
}

func updateImages(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	accessKey := "LYgEIxfrKt7h6tf2ero1VrtwxxqUr1qmIuqhiV2n"
	secretKey := "PTihHOOPlnNDgoi25fo4x7JvSZ8lX2zAw3cqXuUX"
	mac := qbox.NewMac(accessKey, secretKey)
}
