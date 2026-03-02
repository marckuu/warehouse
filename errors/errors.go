package errors

import "errors"

var ErrItemAlreadyExist = errors.New("товар уже существует")
var ErrFieldAreEmpty = errors.New("передано пустое значение")
var ErrItemNotFound = errors.New("товар не найден")
