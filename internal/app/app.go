package app

import (
	"encoding/json"
	"go-assessment/internal/config"
	"go-assessment/internal/userrepository"
	"go-assessment/internal/web"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type db interface {
	GetUser(userId string) (userrepository.User, error)
	DeleteUser(userId string) error
	UpdateUser(user userrepository.User) error
	CreateUser(firstName, lastName string) error
}

type App struct {
	Config *config.Config
	DB     db
}

// New creates a new App
func New(cfg config.Config, db db) App {
	app := App{
		Config: &cfg,
		DB:     db,
	}

	return app
}

// router function handles assignment of routes to handlers
// define your paths and middleware here
func (a *App) router() http.Handler {
	r := mux.NewRouter()

	// health check end point, initialized without validation middleware
	r.HandleFunc("/health-check", a.HealthCheckHandler).Methods(http.MethodGet)

	// user endpoints
	r.HandleFunc("/user", a.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{userId}", a.UpdateUserHandler).Methods(http.MethodPatch)
	r.HandleFunc("/user/{userId}", a.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{userId}", a.DeleteUserHandler).Methods(http.MethodDelete)

	return r
}

// HealthCheckHandler should be used to check the health of the application
func (a *App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData userrepository.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		web.RespondError("error creating user: "+err.Error(), w, http.StatusBadRequest)
	}

	err = json.Unmarshal(b, &userData)
	if err != nil {
		web.RespondError("error creating user: "+err.Error(), w, http.StatusBadRequest)
	}

	err = a.DB.CreateUser(userData.FirstName, userData.LastName)
	if err != nil {
		web.RespondError("error creating user: "+err.Error(), w, http.StatusInternalServerError)
	}

	web.Respond(web.Response{Message: "user created successfully"}, w, http.StatusCreated)
}

func (a *App) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData userrepository.User
	params := mux.Vars(r)
	userData.Id = params["userId"]

	b, err := io.ReadAll(r.Body)
	if err != nil {
		web.RespondError("error updating user: "+err.Error(), w, http.StatusBadRequest)
	}

	err = json.Unmarshal(b, &userData)
	if err != nil {
		web.RespondError("error updating user: "+err.Error(), w, http.StatusBadRequest)
	}

	err = a.DB.UpdateUser(userData)
	if err != nil {
		web.RespondError("error updating user: "+err.Error(), w, http.StatusInternalServerError)
	}

	web.Respond(web.Response{Message: "user updated successfully"}, w, http.StatusOK)
}

func (a *App) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData userrepository.User
	params := mux.Vars(r)
	userData.Id = params["userId"]

	user, err := a.DB.GetUser(userData.Id)
	if err != nil {
		web.RespondError("error getting user: "+err.Error(), w, http.StatusInternalServerError)
	}

	userResponse := web.DataResponse{
		Data: []userrepository.User{user},
	}

	web.Respond(userResponse, w, http.StatusOK)
}

func (a *App) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData userrepository.User

	params := mux.Vars(r)
	userData.Id = params["userId"]

	err := a.DB.DeleteUser(userData.Id)
	if err != nil {
		web.RespondError("error deleting user: "+err.Error(), w, http.StatusInternalServerError)
	}

	web.Respond("user deleted successfully", w, http.StatusNoContent)
}
