package dao

import (
	"context"

	recipe "github.com/baking-code/recipeasy-go/internal/recipe"
)

type RecipeDao interface {
	GetRecipe(ctx context.Context, id recipe.Id) (recipe.Recipe, error)
	CreateRecipe(ctx context.Context, recipe recipe.Recipe) (recipe.Recipe, error)
	UpdateRecipe(ctx context.Context, id recipe.Id, update recipe.Recipe) (recipe.Recipe, error)
	ListRecipes(ctx context.Context) ([]recipe.Recipe, error)
	DeleteRecipe(ctx context.Context, id recipe.Id) (bool, error)
}
