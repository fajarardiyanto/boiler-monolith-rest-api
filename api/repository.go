package api

import "net/http"

type PostControllerRepository interface {
	GetPost(w http.ResponseWriter, r *http.Request)
}
