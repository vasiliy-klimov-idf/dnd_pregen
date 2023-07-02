package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"pregen/backend/backgr"
	"pregen/db"
	"strings"
)

var htmlSitePath string
var assetsSitePath, assetsSiteRootPath string

func init() {
	InitServerPathVars(true)
	db.CheckPostgresDB()
	backgr.GenerateBackground("")
}

func InitServerPathVars(status bool) {
	db.InitPostgresENV(status)

	if status == true {
		assetsSitePath = "/usr/share/nginx/html/assets"
		assetsSiteRootPath = "./usr/share/nginx/html/assets"
		htmlSitePath = "/usr/share/nginx/html/*.html"
	} else {
		assetsSitePath = "frontend/assets"
		assetsSiteRootPath = "./frontend/assets"
		htmlSitePath = "frontend/html/*.html"
	}
}

func main() {
	router := gin.Default()

	// api method
	v1 := router.Group("api/v1/")
	{
		//v1.GET("/dice:number", func(c *gin.Context) {
		//	number := c.Param("number")
		//	api.GetDice(c, number)
		//})
		v1.GET("/roll", func(c *gin.Context) {
			api.GetMultiRoll(c)
		})

		v1.GET("/get-character", func(c *gin.Context) {
			api.GetRandomCharacter(c)
		})

	}

	//site pages
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static(assetsSitePath, assetsSiteRootPath)
	router.LoadHTMLGlob(htmlSitePath)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dices.html", gin.H{
			"title": "Dice Roller",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "character_box.html", gin.H{
			"title": "Character Box",
		})
	})
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"content": "This is an about page...",
		})
	})
	router.GET("/version", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(Version+" VK_RED23"+"\n"))
	})

	//router.Run(":848") //local
	router.RunTLS(":444", "/etc/letsencrypt/live/diceroll.swn.by/fullchain.pem", "/etc/letsencrypt/live/diceroll.swn.by/privkey.pem") //prod
}
