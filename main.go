package main

import (
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"pregen/api"
	"pregen/db"
	"strings"
	"time"
)

var htmlSitePath string
var assetsSitePath, assetsSiteRootPath string

func init() {
	InitServerPathVars(true)
	db.PingMongoDB()
}

//	func TestInsert() {
//		client := db.ConnectToDB()
//		coll := client.Database("data").Collection("classes")
//
//		var class classes.ClassWriteToBD
//		json.Unmarshal(db.JsonTemp, &class)
//
//		docs := []interface{}{}
//
//		for _, cl := range class {
//			fmt.Println(cl)
//			docs = append(docs, cl)
//		}
//
//		result, err := coll.InsertMany(context.TODO(), docs)
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
//		for _, id := range result.InsertedIDs {
//			fmt.Printf("Inserted document with _id: %v\n", id)
//		}
//	}
//
// да уберу я
func InitServerPathVars(status bool) {
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

	// This makes it so each ip can only make 3 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 3,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// api method
	v1 := router.Group("api/v1/")
	{
		//v1.GET("/dice:number", func(c *gin.Context) {
		//	number := c.Param("number")
		//	api.GetDice(c, number)
		//})
		v1.GET("/roll", mw, func(c *gin.Context) {
			api.GetMultiRoll(c)
		})

		v1.GET("/get-character", mw, func(c *gin.Context) {
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
		c.HTML(http.StatusOK, "character_box_ru.html", gin.H{
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

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
	time.Sleep(1 * time.Second)
}
