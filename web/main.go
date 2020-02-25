package web

import "github.com/lualfe/casamento/app"

// Web instance
type Web struct {
	App *app.App
}

// New creates a new Web instance
func New(a *app.App) (*Web, error) {
	return &Web{App: a}, nil
}
