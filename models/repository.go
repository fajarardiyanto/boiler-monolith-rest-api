package models

import "github.com/fajarardiyanto/boiler-monolith-rest-api/interfaces"

type PostRepository interface {
	GetPost() ([]interfaces.Post, error)
}
