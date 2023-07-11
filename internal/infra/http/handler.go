package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

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
		svc service.RecipeService
	}
)

const (
	cookHandlerName = "cook-handler"
)

func NewCookHandler(svc service.RecipeService, opts ...sys.Option) *CookHandler {
	return &CookHandler{
		SimpleCore: sys.NewCore(cookHandlerName, opts...),
		svc:        svc,
	}
}

func (h *CookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetBooks not implemented yet"))
	if err != nil {
		h.Log().Error(err)
	}
}

func (h *CookHandler) PostBook(w http.ResponseWriter, r *http.Request) {
	// Session
	userID, err := h.User(r)
	if err != nil {
		err = errors.Wrap(err, "post book error")
		h.handleError(w, err)
	}

	// Request to Transport
	defer h.closeBody(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, InvalidRequestErr)
	}

	var book service.CreateBookReq
	err = json.Unmarshal(body, &book)
	if err != nil {
		h.handleError(w, InvalidRequestDataErr)
	}
	book.UserID = userID

	// Use svc to save the model
	ctx := r.Context()
	res := h.Service().CreateBook(ctx, book)
	if err = res.Err(); err != nil {
		err = errors.Wrap(err, "post book error")
		h.handleError(w, err)
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

func (h *CookHandler) PostRecipe(w http.ResponseWriter, r *http.Request, bookID string) {
	// Session
	userID, err := h.User(r)
	if err != nil {
		err = errors.Wrap(err, "post book error")
		h.handleError(w, err)
	}

	// Request to Transport
	defer h.closeBody(r.Body)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, InvalidRequestErr)
	}

	var recipe service.CreateRecipeReq
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		h.handleError(w, InvalidRequestDataErr)
	}
	recipe.UserID = userID
	recipe.BookID = bookID

	// Use svc to save the model
	ctx := r.Context()
	res := h.Service().CreateRecipe(ctx, recipe)
	if err = res.Err(); err != nil {
		err = errors.Wrap(err, "post book error")
		h.handleError(w, err)
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

func (h *CookHandler) User(r *http.Request) (userID string, err error) {
	// Authentication mechanism not yet established.
	// WIP: A hardcoded value is returned for now.
	uid := "c4c109ad-f178-400a-b86d-3b0d548d852c"

	_, err = uuid.Parse(uid)
	if err != nil {
		return "", NoUserErr
	}

	return uid, nil
}

// closeBody close the body and log errors if happened.
func (h *CookHandler) closeBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		h.Log().Error(errors.Wrap(err, "failed to close body"))
	}
}

// Handler interface

// Service returns a recipe svc implementation.
func (h *CookHandler) Service() service.RecipeService {
	return h.svc
}
