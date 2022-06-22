package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/todoTask/handler/sessionHandler"
	"github.com/todoTask/handler/taskHandler"
	"github.com/todoTask/handler/userHandler"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Get("/test", userHandler.Greet)
		api.Get("/allUsers", userHandler.AllUsers)
		api.Post("/addUser", userHandler.AddRow)
		api.Put("/updateUser", userHandler.UpdateRow)
		api.Delete("/deleteUser", userHandler.DeleteRow)
		api.Get("/getTasks", taskHandler.GetTasks)
		api.Post("/addTask", taskHandler.AddTask)
		api.Put("/completeTask", taskHandler.CompleteTask)
		api.Delete("/deleteTask", taskHandler.DeleteTask)
		api.Post("/signin", sessionHandler.SignInUser)
		api.Post("/signout", sessionHandler.SignOut)
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
