package march

import (
    "net/http"
    "fmt"
    "strings"
)

var ms = NewMarchStore()

func KeyFromURL(URL string) string {
    pieces := strings.Split(URL, "/")
    if len(pieces) > 1 {
        return strings.Split(URL, "/")[1]
    } else {
        return ""
    } 
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
    key := KeyFromURL(r.URL.Path)
    w.Write([]byte(ms.GetKey(key))) 
}

func DelHandler(w http.ResponseWriter, r *http.Request) {
    key := KeyFromURL(r.URL.Path)
    ms.DeleteKey(key)
    w.Write([]byte("SUCCESS"))
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
    key := KeyFromURL(r.URL.Path)
    val := make([]byte, r.ContentLength)
    _, err := r.Body.Read(val)

    if err != nil {
        ms.PutKey(key, string(val))
    }
    w.Write([]byte("SUCCESS"))
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
    method := r.Method

    switch method {
    case "GET":
        GetHandler(w, r)
        break
    case "DELETE":
        DelHandler(w, r)
        break
    case "PUT":
        PutHandler(w, r)
        break
    }

    w.(http.Flusher).Flush()
}

func Listen(port int) {
    http.HandleFunc("/", HttpHandler)
    http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
