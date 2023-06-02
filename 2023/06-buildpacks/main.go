package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func hello(w http.ResponseWriter, req *http.Request) {
	log.Info("called hello endpoint")
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	log = logger.Sugar()

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":1337", nil)
}
