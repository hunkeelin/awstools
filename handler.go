package main
import(
    "github.com/json-iterator/go"
    "net/http"
    "fmt"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func mainHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        err := post(w, r)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(500)
            w.Write([]byte(err.Error()))
        }
    default:
        fmt.Println("invalid method")
        w.WriteHeader(500)
    }
}

