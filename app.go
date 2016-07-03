import (
  "net/http"
  "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome!")
}

func main() {
  http.Handle("/", handler)
  http.ListenAndServe(":8080",nil)
}
