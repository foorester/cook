package service

import (
	"errors"

	"github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/domain/model"
)

type (
	BookValidator struct {
		domain.Validator
		Model model.Book
	}
)

func NewBookValidator(m model.Book) BookValidator {
	return BookValidator{
		Validator: domain.NewValidator(),
		Model:     m,
	}
}

func (v BookValidator) ValidateForCreate() error {
	// Username
	ok0 := v.ValidateRequiredName()
	ok1 := v.ValidateMinLengthName(2)

	if ok0 && ok1 {
		return nil
	}

	return errors.New("recipe book has errors")
}

func (v BookValidator) ValidateForUpdate() error {
	return errors.New("not implemented yet")
}

func (v BookValidator) ValidateRequiredName(errMsg ...string) (ok bool) {
	book := v.Model

	ok = v.ValidateRequired(book.Name)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.RequiredErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}

func (v BookValidator) ValidateMinLengthName(min int, errMsg ...string) (ok bool) {
	m := v.Model

	ok = v.ValidateMinLength(m.Name, min)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.MinLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}

func (v BookValidator) ValidateMaxLengthName(max int, errMsg ...string) (ok bool) {
	m := v.Model

	ok = v.ValidateMaxLength(m.Name, max)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.MaxLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}

type (
	RecipeValidator struct {
		domain.Validator
		Model model.Recipe
	}
)

func NewRecipeValidator(m model.Recipe) RecipeValidator {
	return RecipeValidator{
		Validator: domain.NewValidator(),
		Model:     m,
	}
}

func (v RecipeValidator) ValidateForCreate() error {
	ok0 := v.ValidateRequiredName()
	ok1 := v.ValidateMinLengthName(4)

	if ok0 && ok1 {
		return nil
	}

	return errors.New("recipe book has errors")
}

func (v RecipeValidator) ValidateForUpdate() error {
	return errors.New("not implemented yet")
}

func (v RecipeValidator) ValidateRequiredName(errMsg ...string) (ok bool) {
	book := v.Model

	ok = v.ValidateRequired(book.Name)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.RequiredErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}

func (v RecipeValidator) ValidateMinLengthName(min int, errMsg ...string) (ok bool) {
	m := v.Model

	ok = v.ValidateMinLength(m.Name, min)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.MinLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}

func (v RecipeValidator) ValidateMaxLengthName(max int, errMsg ...string) (ok bool) {
	m := v.Model

	ok = v.ValidateMaxLength(m.Name, max)
	if ok {
		return true
	}

	msg := domain.ValidatorMsg.MaxLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	v.Errors["Name"] = append(v.Errors["Name"], msg)
	return false
}
