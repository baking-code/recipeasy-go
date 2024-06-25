package recipe

import (
	"encoding/json"
	"time"
)

type Id string

// Duration is a custom type for time.Duration with JSON marshalling support
type Duration struct {
	time.Duration
}

// MarshalJSON implements the json.Marshaler interface
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

type Recipe struct {
	Id          Id       `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Method      []string `json:"method"`
	Ingredients []string `json:"ingredients"`
	Tags        []string `json:"tags"`
	Duration    Duration `json:"duration"`
}

type RecipeOption func(*Recipe)

func WithName(n string) RecipeOption {
	return func(r *Recipe) {
		r.Name = n
	}
}

func WithDescription(d string) RecipeOption {
	return func(r *Recipe) {
		r.Description = d
	}
}

func WithMethod(m []string) RecipeOption {
	return func(r *Recipe) {
		r.Method = m
	}
}

func WithIngredients(i []string) RecipeOption {
	return func(r *Recipe) {
		r.Ingredients = i
	}
}

func WithTags(t []string) RecipeOption {
	return func(r *Recipe) {
		r.Tags = t
	}
}

func WithDuration(d time.Duration) RecipeOption {
	return func(r *Recipe) {
		r.Duration = Duration{d}
	}
}

func RecipeUpdate(orig Recipe, opts ...RecipeOption) *Recipe {
	updatedRecipe := orig
	for _, opt := range opts {
		opt(&updatedRecipe)
	}
	return &updatedRecipe
}
