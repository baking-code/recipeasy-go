package service

import (
	"context"

	"github.com/baking-code/recipeasy-go/internal/dao"
	recipe "github.com/baking-code/recipeasy-go/internal/recipe"
)

type Service interface {
	GetRecipe(ctx context.Context, id recipe.Id) (recipe.Recipe, error)
	CreateRecipe(ctx context.Context, recipe recipe.Recipe) (recipe.Recipe, error)
	UpdateRecipe(ctx context.Context, id recipe.Id, update []recipe.RecipeOption) (recipe.Recipe, error)
	ListRecipes(ctx context.Context) ([]recipe.Recipe, error)
	DeleteRecipe(ctx context.Context, id recipe.Id) (bool, error)
}

type SimpleRecipeService struct {
	dao dao.RecipeDao
}

// CreateRecipe implements Service.
func (s *SimpleRecipeService) CreateRecipe(ctx context.Context, recipe recipe.Recipe) (recipe.Recipe, error) {
	return s.dao.CreateRecipe(ctx, recipe)
}

// DeleteRecipe implements Service.
func (s *SimpleRecipeService) DeleteRecipe(ctx context.Context, id recipe.Id) (bool, error) {
	return s.dao.DeleteRecipe(ctx, id)
}

// GetRecipe implements Service.
func (s *SimpleRecipeService) GetRecipe(ctx context.Context, id recipe.Id) (recipe.Recipe, error) {
	return s.dao.GetRecipe(ctx, id)
}

// ListRecipes implements Service.
func (s *SimpleRecipeService) ListRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	return s.dao.ListRecipes(ctx)
}

// UpdateRecipe implements Service.
func (s *SimpleRecipeService) UpdateRecipe(ctx context.Context, id recipe.Id, update []recipe.RecipeOption) (recipe.Recipe, error) {
	r, err := s.dao.GetRecipe(ctx, id)
	if err != nil {
		return r, err
	} else {
		recipe.RecipeUpdate(r, update...)
		return r, nil
	}
}

func NewSimpleService(d dao.RecipeDao) Service {
	return &SimpleRecipeService{d}
}

func NewSimpleServiceWithInMemoryDao() Service {
	return &SimpleRecipeService{dao.NewInMemoryDao()}
}
