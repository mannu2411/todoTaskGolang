package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/todoTask/handler/sessionHandler"
	"github.com/todoTask/handler/taskHandler"
	"github.com/todoTask/handler/userHandler"
	"github.com/todoTask/middlewareUtils"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()

	router.Route("/api", func(api chi.Router) {
		/*test APIs
		api.Get("/test", userHandler.Test)
		api.Get("/all-users", userHandler.AllUsers)*/

		api.Route("/Auth", func(auth chi.Router) {
			auth.Post("/sign-in", sessionHandler.SignInUser)
			auth.Post("/sign-up", userHandler.AddUser)
		})

		api.Route("/", func(r chi.Router) {
			r.Use(middlewareUtils.AuthMiddleware)
			r.Use(middlewareUtils.JWTAuthMiddleware)
			r.Use(middlewareUtils.GetUserContext)
			r.Delete("/sign-out", sessionHandler.SignOut)

			r.Route("/user", func(user chi.Router) {
				user.Put("/update-user", userHandler.UpdateUser)
				user.Delete("/delete-user", userHandler.DeleteUser)
			})

			r.Route("/task", func(task chi.Router) {
				task.Get("/get-tasks", taskHandler.GetTasks)
				task.Post("/add-task", taskHandler.AddTask)
				task.Put("/complete-task", taskHandler.CompleteTask)
				task.Delete("/delete-task", taskHandler.DeleteTask)
			})
		})

	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
