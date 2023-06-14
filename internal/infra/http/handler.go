package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/domain/service"
	"github.com/foorester/cook/internal/infra/openapi"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	CookHTTPHandler interface {
		sys.Core
		openapi.ServerInterface
		Service() service.RecipeService
	}

	CookHandler struct {
		*sys.SimpleCore
		service service.RecipeService
	}
)

const (
	cookHandlerName = "cook-handler"
)

func NewCookHandler(opts ...sys.Option) *CookHandler {
	return &CookHandler{
		SimpleCore: sys.NewCore(cookHandlerName, opts...),
	}
}

func (h *CookHandler) GetRecipeBooks(w http.ResponseWriter, r *http.Request) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipeBooks not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PostRecipeBook(w http.ResponseWriter, r *http.Request) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostRecipeBook not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteBook not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetBook not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PutBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutBook not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetRecipes(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipes not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PostRecipe(w http.ResponseWriter, r *http.Request, bookId string) {
	// Request to model
	// Here for now but will be wrapped in am externalized function
	var recipe model.Recipe
	defer func() {
		if err := r.Body.Close(); err != nil {
			// Handle the error here, such as logging or returning an error response.
		}
	}()

	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(InvalidJSONBodyErr.Error()))
		if err != nil {
			h.Log().Errorf("post recipe error: %w", err)
		}
		return
	}

	defer h.closeBody(r.Body)

	// Basic validation

	// Use service to save the model
	ctx := r.Context()
	err = h.Service().CreateRecipe(ctx, recipe)
	if err != nil {

	}
}

func (h *CookHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PutRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetSteps(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetSteps not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PostStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostStep not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, stepId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetStep not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PutStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, stepId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutStep not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetIngredients(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetIngredients not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PostIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) GetIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PutIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

// Helpers

// closeBody close the body and log errors if happened.
func (h *CookHandler) closeBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		h.Log().Error(errors.Wrap("failed to close body", err))
	}
}

// Handler interface

// Service returns a recipe service implementation.
func (h *CookHandler) Service() service.RecipeService {
	return h.service
}
