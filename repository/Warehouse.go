package repository

import (
	"Warehouse/errors"
	"Warehouse/models"
	"sync"
)

type Warehouse struct {
	items map[string]models.Item
	mtx   sync.RWMutex
}

func NewWarehouse() Warehouse {
	return Warehouse{
		items: make(map[string]models.Item),
	}
}

func (w *Warehouse) AddItem(item models.Item) error {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	_, ok := w.items[item.Id]
	if ok {
		return errors.ErrItemAlreadyExist
	}

	w.items[item.Id] = item

	return nil
}

func (w *Warehouse) DeleteItem(itemId string) error {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	_, ok := w.items[itemId]
	if !ok {
		return errors.ErrItemNotFound
	}

	delete(w.items, itemId)

	return nil
}

func (w *Warehouse) GetItem(itemId string) (models.Item, error) {
	w.mtx.RLock()
	defer w.mtx.RUnlock()
	item, ok := w.items[itemId]
	if !ok {
		return models.Item{}, errors.ErrItemNotFound
	}

	return item, nil
}

func (w *Warehouse) GetAllTItems() map[string]models.Item {
	w.mtx.RLock()
	defer w.mtx.RUnlock()
	return w.items
}

func (w *Warehouse) GetItemLighterThan(weight float64) map[string]models.Item {
	w.mtx.RLock()
	defer w.mtx.RUnlock()
	filteredItems := make(map[string]models.Item)

	for _, item := range w.items {
		if item.Weight <= weight {
			filteredItems[item.Id] = item
		}
	}

	return filteredItems
}

func (w *Warehouse) ChangeItemTitle(itemId string, title string) (models.Item, error) {
	w.mtx.RLock()
	defer w.mtx.RUnlock()
	item, ok := w.items[itemId]
	if !ok {
		return models.Item{}, errors.ErrItemNotFound
	}

	if err := item.ChangeTitle(title); err != nil {
		return models.Item{}, err
	}

	w.items[itemId] = item

	return item, nil
}
