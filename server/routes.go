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
		router.Route("/user", func(user chi.Router) {
			user.Get("/test", userHandler.Test)
			user.Get("/all-users", userHandler.AllUsers)
			user.Post("/add-user", userHandler.AddUser)
			user.Put("/update-user", userHandler.UpdateUser)
			user.Delete("/delete-user", userHandler.DeleteUser)
		})
		router.Route("/task", func(task chi.Router) {
			task.Get("/get-tasks", taskHandler.GetTasks)
			task.Post("/add-task", taskHandler.AddTask)
			task.Put("/complete-task", taskHandler.CompleteTask)
			task.Delete("/delete-task", taskHandler.DeleteTask)
		})
		api.Post("/signin", sessionHandler.SignInUser)
		api.Post("/signout", sessionHandler.SignOut)
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
