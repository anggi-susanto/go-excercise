package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

//ErrSkuNotFound sku not found
var ErrSkuNotFound = errors.New("sku not found")

//ErrPromoRulesNotFound promo rule not found
var ErrPromoRulesNotFound = errors.New("promo rule not found")
