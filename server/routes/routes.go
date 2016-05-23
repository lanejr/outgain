package routes

import (
	"controller"
	"net/http"
)

//GetHandler returns a mux that mapps routes to controller actions
func GetHandler(static string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(static)))
	mux.HandleFunc("/ping", controller.PingHandler)

	return mux
}
