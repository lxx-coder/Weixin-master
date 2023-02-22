package main

import (
	"crypto/sha1"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)

const (
	token = "weixin"
)

func checkSignature(c *gin.Context) {
	signature := c.Param("signature")
	timestamp := c.Param("timestamp")
	nonce := c.Param("nonce")
	echostr := c.Param("echostr")

	if check(signature, timestamp, nonce) {
		c.String(200, echostr)
	}
	c.String(500, "Bad Signature")

}

func check(signature, timestamp, nonce string) bool {
	arr := []string{token, timestamp, nonce}

	sort.Strings(arr)

	data := sha1.Sum([]byte(strings.Join(arr, "")))

	return string(data[:]) == signature
}

func handleMsg(c *gin.Context) {

}

func main() {
	g := gin.Default()
	g.GET("/", func(context *gin.Context) {
		checkSignature(context)
	})
	http.ListenAndServe(":8080", nil)
}
