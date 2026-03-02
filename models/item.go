package models

import (
	"Warehouse/errors"
)

type Item struct {
	Id     string
	Title  string
	Price  float64
	Weight float64
}

func NewItem() Item {
	return Item{
		Title:  "",
		Price:  0.0,
		Weight: 0,
	}
}

func (i *Item) ChangeTitle(title string) error {
	if title == "" {
		return errors.ErrFieldAreEmpty
	}
	i.Title = title
	return nil
}
