// Fetch prints the content found at a URL.
// Excercise Use io.Copy(dst, src) instead of ioutil.ReadAll to copy the
// response body to os.Stdout without requireing a buffer large enough to hold
// the entire stream. Check the error result of io.Copy

package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
)

func main() {
  for _, url := range os.Args[1:] {
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    b, err := io.Copy(os.Stdout, resp.Body)

    if err != nil {
      log.Fatal(err)
    }
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", b)
  }
}
