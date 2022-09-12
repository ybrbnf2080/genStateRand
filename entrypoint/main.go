package entrypoint

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ybrbnf2080/genStateRand/iternal/render"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

var LatestPict string

func Init() {

}
func Main() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.StaticFile("/", "./static/index.html")

	r.GET("/exit", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		os.Exit(1)
	})

	r.GET("/latest", func(c *gin.Context) {
		if _, ok := c.Request.Header["Connection"]; ok {
			c.Data(200, "text/html; charset=utf-8", []byte("<p >"+strings.Replace(LatestPict, "\n", "<br>", -1)+"</p> <style>body {   font-family: Courier, monospace; } </style>"))
		} else {
			c.Data(200, "text/plain; charset=utf-8", []byte(LatestPict))

		}
	})

	r.POST("/pict", func(c *gin.Context) {

		width, err := strconv.Atoi(c.PostForm("width"))
		if err != nil {
			width = 100
		}
		height, err := strconv.Atoi(c.PostForm("height"))
		if err != nil {
			height = 100
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		log.Println(file.Filename)
		openedFile, _ := file.Open()
		pict := render.RenderPict(openedFile, int(height), width)
		LatestPict = pict

		if c.PostForm("formatting") == "" {
			c.Data(200, "text/plain; charset=utf-8", []byte(pict))
		} else {
			c.Data(200, "text/html; charset=utf-8", []byte("<p >"+strings.Replace(pict, "\n", "<br>", -1)+"</p> <style> @import url('http://fonts.cdnfonts.com/css/terminus'); 			body {   font-family: Courier, monospace; } </style>"))

		}
	})

	return r // listen and serve on 0.0.0.0:8080
}
