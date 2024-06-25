package dao

import (
	"context"
	"errors"

	"github.com/baking-code/recipeasy-go/internal/recipe"
)

type InMemoryDao struct {
	ById map[recipe.Id]recipe.Recipe
}

// CreateRecipe implements RecipeDao.
func (d *InMemoryDao) CreateRecipe(ctx context.Context, recipe recipe.Recipe) (recipe.Recipe, error) {
	found, ok := d.ById[recipe.Id]
	if !ok {
		return found, errors.New("already exists")
	} else {
		d.ById[recipe.Id] = recipe
		return recipe, nil
	}
}

// DeleteRecipe implements RecipeDao.
func (d *InMemoryDao) DeleteRecipe(ctx context.Context, id recipe.Id) (bool, error) {
	_, ok := d.ById[id]
	if !ok {
		return false, errors.New("not found")
	} else {
		delete(d.ById, id)
		return true, nil
	}
}

// GetRecipe implements RecipeDao.
func (d *InMemoryDao) GetRecipe(ctx context.Context, id recipe.Id) (recipe.Recipe, error) {
	recipe, ok := d.ById[id]
	if !ok {
		return recipe, errors.New("not found")
	} else {
		return recipe, nil
	}
}

// ListRecipes implements RecipeDao.
func (d *InMemoryDao) ListRecipes(ctx context.Context) ([]recipe.Recipe, error) {
	arr := []recipe.Recipe{}
	for _, r := range d.ById {
		arr = append(arr, r)
	}
	return arr, nil
}

// UpdateRecipe implements RecipeDao.
func (d *InMemoryDao) UpdateRecipe(ctx context.Context, id recipe.Id, update recipe.Recipe) (recipe.Recipe, error) {
	recipe, ok := d.ById[id]
	if !ok {
		return recipe, errors.New("not found")
	} else {
		recipe = update
		return recipe, nil
	}
}

func NewInMemoryDao() RecipeDao {
	return &InMemoryDao{}
}
