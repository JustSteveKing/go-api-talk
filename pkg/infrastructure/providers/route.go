package providers

import "github.com/gorilla/mux"

// RouteProvider will create a new HTTP Router Mux to use within our application
func RouteProvider() *mux.Router {
	return mux.NewRouter()
}
