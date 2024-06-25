package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/baking-code/recipeasy-go/internal/recipe"
	"github.com/baking-code/recipeasy-go/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// Register connects the handlers to the router.
func (t *Handler) Register(r chi.Router) {
	r.Get("/{id}", t.getRecipe)
}

func (t *Handler) getRecipe(w http.ResponseWriter, r *http.Request) {
	idRaw := chi.URLParam(r, "id")
	ctx := r.Context()
	recipe, err := t.service.GetRecipe(ctx, recipe.Id(idRaw))
	if err != nil {
		message := "cannot get recipe"
		slog.Error(message, err)
		http.Error(w, message, 404)
	} else {
		// Marshal data into JSON
		jsonData, err := json.Marshal(recipe)
		if err != nil {
			message := "error marshalling JSON"
			slog.Error(message, err)
			http.Error(w, message, 400)
		} else {
			w.Write(jsonData)
		}
	}

}
