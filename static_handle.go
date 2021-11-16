package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//go:embed static/*
var staticFS embed.FS

var contentType map[string]string = map[string]string{
	"html":  "text/html; charset=utf-8",
	"css":   "text/css",
	"js":    "application/x-javascript",
	"png":   "image/png",
	"jpg":   "image/jpeg",
	"gif":   "image/gif",
	"ico":   "image/x-icon",
	"svg":   "image/svg+xml",
	"bmp":   "image/bmp",
	"webp":  "image/webp",
	"jpeg":  "image/jpeg",
	"eot":   "application/octet-stream",
	"ttf":   "application/octet-stream",
	"woff":  "application/octet-stream",
	"woff2": "application/octet-stream",
}

func staticHandler(context *gin.Context) {
	path := context.Request.URL.Path
	if path[len(path)-1:] == "/" {
		path += "index.html"
	}
	b, err := staticFS.ReadFile("static" + path)
	if err != nil {
		context.Data(http.StatusOK, "text/html; charset=utf-8", []byte{})
		return
	}
	ext := path[strings.LastIndex(path, ".")+1:]
	ct, ok := contentType[ext]
	if !ok {
		ct = "text/html; charset=utf-8"
	}
	context.Data(http.StatusOK, ct, b)
}
