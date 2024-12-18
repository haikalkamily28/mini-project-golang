package repository

import (
	"mini/entity"
)

type ItemRepository interface {
    GetAllItems() ([]entity.Item, error)
    CreateItem(item *entity.Item) error
    UpdateItem(item *entity.Item) error 
    DeleteItem(id uint) error 
}
