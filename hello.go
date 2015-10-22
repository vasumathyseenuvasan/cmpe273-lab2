package main
import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "io/ioutil"
    "io"
    m "encoding/json"
  )

type Req struct{
  Name string `json:"name"`
}
type Resp struct{
  Name string `json:"greeting"`
}

func hellopost(rw http.ResponseWriter, req *http.Request , p httprouter.Params) {
    var response Resp
    body,_ := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
    request:=&Req{}
    m.Unmarshal(body, &request)
    response.Name = "Hello, "+request.Name + "!"
    rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
    rw.WriteHeader(http.StatusCreated)
    if err := m.NewEncoder(rw).Encode(response); err != nil {
       panic(err)
   }
}


func main() {
    mux := httprouter.New()
    mux.POST("/hello", hellopost)
    server := http.Server{
            Addr:        ":8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
