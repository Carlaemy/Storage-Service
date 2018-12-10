package app

import (
	"log"
	"net/http"

	"./Handler"
	"github.com/gorilla/mux"
)

// App has router
type App struct {
	Router *mux.Router
}

// App initialize
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	a.Get("/users", a.GetUsers)
	a.Post("/create/user", a.CreateUser)
	a.Get("/user/{id}", a.GetUserById)
	a.Delete("/delete/{id}", a.DeleteUser)
	a.Post("/login", a.LoginUser)

	a.Get("/files", a.GetFiles)
	a.Post("/create/file", a.CreateFile)
	a.Get("/file/{id}", a.GetFileById)
	a.Delete("/delete/file/{id}", a.DeleteFile)
}

// GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST", "OPTIONS")
}

//  DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Users Data
func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetUsers(w, r)
}
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(w, r)
}
func (a *App) GetUserById(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(w, r)
}
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(w, r)
}
func (a *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	handler.LoginUser(w, r)
}

// Handlers to manage Files Data
func (a *App) GetFiles(w http.ResponseWriter, r *http.Request) {
	handler.GetFiles(w, r)
}
func (a *App) CreateFile(w http.ResponseWriter, r *http.Request) {
	handler.CreateFile(w, r)
}
func (a *App) GetFileById(w http.ResponseWriter, r *http.Request) {
	handler.GetFile(w, r)
}
func (a *App) DeleteFile(w http.ResponseWriter, r *http.Request) {
	handler.DeleteFile(w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Println("Server Listening http://localhost:3000")
	log.Fatal(http.ListenAndServe(host, a.Router))
}
