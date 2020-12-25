package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"
)

var bot *linebot.Client

type T struct {
	ChannelSecret      string
	ChannelAccessToken string
	Port               string
	Certfile           string
	Keyfile            string
}
type packageInfo struct {
	Recipient     string `form:"recipient"`
	RecipientAddr string `form:"recipientaddr"`
	Sender        string `form:"sender"`
}

func main() {
	var (
		err error
	)
	yamlPath := flag.String("yaml", "../../env.yaml", "environment yaml file path")
	flag.Parse()
	data, err := ioutil.ReadFile(*yamlPath)
	if err != nil {
		log.Fatal(err)
	}
	t := T{}
	err = yaml.Unmarshal(data, &t)
	// line init
	bot, err = linebot.New(t.ChannelSecret, t.ChannelAccessToken)
	if err != nil {
		fmt.Println(t)
		log.Fatal("linebot.New的問題\n", err)

	}

	// gin區
	router := gin.Default()
	router.POST("/line/callback", lineHandler)
	router.GET("/postsystem", postsystemHandler)
	// router.Run(":8001")
	err = router.RunTLS(t.Port, t.Certfile, t.Keyfile)
	log.Fatal(err)
}

func lineHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}
func postsystemHandler(c *gin.Context) {
	var p packageInfo
	if err := c.ShouldBind(&p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	c.String(http.StatusOK, fmt.Sprint(rand.Int()))
}
