package main

import (
	"fmt"
	"net/http"

	"github.com/Narasimha1997/httppool"
)

/*
	The handler will be passed with *http.ResponseWriter, *http.Request pointers,
	channel is a notifier which notifies the dispatcher about the routine completion,
	Pass channel<-true after completion
*/
func testRoute(w *http.ResponseWriter, r *http.Request, channel chan<- bool) {

	//use the r : http.Request

	fmt.Fprintf(*w, "<html></head>")
	fmt.Fprintf(*w, "<title>Sample test page</title>")
	fmt.Fprintf(*w, "</head>")
	fmt.Fprintf(*w, "<body>")
	fmt.Fprintf(*w, "<h4>Welcome to sample page</h4>")
	fmt.Fprintf(*w, "</body>")
	fmt.Fprintf(*w, "</html>")

	channel <- true
}

func main() {
	//Create a pool with 300 workers and 100 max queue size
	pool := httppool.NewRouteHandler(300, 100)

	//Register the route /test
	pool.RegisterRoute("/test", testRoute)

	//Register the route handler and redirect the root requests to our pool
	//This is how we glue http and pool modules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//dispatch
		pool.Dispatch(w, r)
	})

	//Start the server on :9000 port
	http.ListenAndServe(":9000", nil)
}
