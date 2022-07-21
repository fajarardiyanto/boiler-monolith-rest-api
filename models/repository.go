package models

import "privy/interfaces"

type PostRepository interface {
	GetPost() ([]interfaces.Post, error)
}
