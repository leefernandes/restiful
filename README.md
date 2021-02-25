# restiful
Super Straightforward Golang Middleware.  
  
There are some sweet middleware packages available from the go community, but I think there is room for more of less. Less additional code, and less of a learning curve for those to follow.

```go
go get github.com/itsleeowen/restiful
```

```go
import "github.com/itsleeowen/restiful"

router.Handler("POST", "/thing", restiful.Handle(
  acl.ValidToken,
  logger.RestApi,
  thing.Post,
))
```


Having middleware control-flow managed by a single http.Handler takes away the need to pass around “next”, and provides an expectation that middleware only needs to return an error, else the flow shall continue. Who doesn’t love returning errors in Go?  
  
Try it out, to me it really just feels right. Our middleware interface is basically ServeHTTP with an error return, and our route handler is an http.Handler processing middleware handlers looking for an error. We can use an http.Request based context package (such as http://www.gorillatoolkit.org/pkg/context), to pass contextual info along the middleware chain.

```go
// http://laicos.com/writing-handsome-golang-middleware/
 
package main
 
import (
  "net/http"
  "log"

	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/context"
	"github.com/itsleeowen/restiful/v1"
)
 
func main() {
	router := httprouter.New()
	router.Handler("GET", "/thing", restiful.Handle(
			MiddleOne,
			MiddleTwo,
			GetThing,
	))
	log.Fatal(http.ListenAndServe(":8080", router))
}
 
 
func MiddleOne(w http.ResponseWriter, r *http.Request) error {
	context.Set(r, "rideTheMiddlewareChain", "like a wave")
	return nil
}
 
func MiddleTwo(w http.ResponseWriter, r *http.Request) error {
	value := context.Get(r, "rideTheMiddlewareChain").(string)
	context.Set(r, "rideTheMiddlewareChain", value+" from the depths below.")
	return nil
}
 
func GetThing(w http.ResponseWriter, r *http.Request) error {
	value := context.Get(r, "rideTheMiddlewareChain").(string)
      context.Clear(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "`+value+`"}`))
	return nil
}
```
