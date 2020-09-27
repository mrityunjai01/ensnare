package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/", entry).Methods("GET")
	return router
}

func entry (writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("/ hit")
	fmt.Fprint(writer, "Welcome to Ensnare")
}
func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
