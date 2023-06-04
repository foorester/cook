package http

import (
	"net/http"

	"github.com/foorester/cook/internal/sys"
)

type (
	CookHandler struct {
		sys.Worker
	}
)

const (
	cookHandlerName = "cook-handler"
)

func NewCookHandler(opts ...sys.Option) *CookHandler {
	return &CookHandler{
		Worker: sys.NewWorker(cookHandlerName, opts...),
	}
}

func (h *CookHandler) GetRecipeBooks(w http.ResponseWriter, r *http.Request) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipeBooks not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PostRecipeBook(w http.ResponseWriter, r *http.Request) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostRecipeBook not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteBook not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetBook not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PutBook(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutBook not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetRecipes(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipes not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PostRecipe(w http.ResponseWriter, r *http.Request, bookId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PutRecipe(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutRecipe not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetDirectionSteps(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetDirectionSteps not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PostDirectionStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostDirectionStep not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, stepId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetStep not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PutStep(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, stepId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutStep not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetIngredients(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetIngredients not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PostIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PostIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("DeleteIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) GetIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("GetIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}

func (h *CookHandler) PutIngredient(w http.ResponseWriter, r *http.Request, bookId string, recipeId string, ingredientId string) {
	//TODO not implemented yet
	_, err := w.Write([]byte("PutIngredient not implemented yet"))
	if err != nil {
		h.Log().Error(err.Error())
	}
}
