package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "time"
)

func main() {
  go func() {
    for {
      log.Println("Starting to ping...")
      err := pinger()
      if err != nil {
        log.Println(err)
      }
      log.Println("Sleeping for 30mins before starting again...")
      time.Sleep(30 * time.Minute)
    }
  }()

  http.HandleFunc("/", webFunc)
  port := os.Getenv("PORT")
  log.Printf("Serving on port %v\n", port)
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    panic(err)
  }
}


func webFunc(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "Pingin' them all")
}

func pinger() error {
  return nil;
}
