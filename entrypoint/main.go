package entrypoint

import (
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ybrbnf2080/genStateRand/iternal/render"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func Init() {

}
func Main() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Static("/", "./static")

	//r.GET("/exit", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//	os.Exit(1)
	//})
	//r.GET("/test", func(c *gin.Context) {
	//	//var file, _ = os.Open("./tmp/image.jpeg")
	//	//c.Data(200, "text/plain; charset=utf-8", []byte(render.RenderPict(file)))
	//})
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
		c.Data(200, "text/plain; charset=utf-8", []byte(render.RenderPict(openedFile, int(height), width)))
	})

	return r // listen and serve on 0.0.0.0:8080
}
