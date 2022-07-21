package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"privy/api"
	"privy/internal/response"
	"privy/middleware"
)

type Routes struct {
	postController api.PostControllerRepository
}

func NewRoutes() Routes {
	return Routes{
		postController: api.NewPostControllers(),
	}
}

func (c *Routes) Run(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusNotFound, map[string]interface{}{
			"message": "Not Found",
		})
	})

	r.HandleFunc("/post", middleware.SetMiddlewareJSON(c.postController.GetPost)).Methods("GET")
}
