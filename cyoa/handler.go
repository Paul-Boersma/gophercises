package main

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/Paul-Boersma/gophercises/cyoa/model"
)

type Handler struct {
	story model.Story
	tmpl  *template.Template
}

type HandlerOption func(h *Handler)

// Handler constructor
func NewHandler(story model.Story, opts ...func(h *Handler)) Handler {
	defaultTemplate := template.Must(template.ParseFiles("./templates/cyoa.gohtml"))
	handler := Handler{
		story: story,
		tmpl:  defaultTemplate,
	}

	for _, opt := range opts {
		opt(&handler)
	}

	return handler
}

// Implement handler interface
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]
	if chapter, ok := h.story[path]; ok {
		err := h.tmpl.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

// Handler options
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *Handler) {
		h.tmpl = t
	}
}
