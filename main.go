package main

import (
  "net/http"
)

func main() {
  const port = "8080"

	serveMux := http.NewServeMux()
  httpServer := &http.Server{
    Addr: ":" + port,
    Handler: serveMux,
  }
  httpServer.ListenAndServe()
}
