package interfaces

import "Phase3/week2/day3/go-pooling-message-broker/models"

type Datastore interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id int) (*models.Item, error)
	InsertItem(item *models.Item) error
	UpdateItem(id int, item *models.Item) error
	DeleteItem(id int) error
}

type MessageQueue interface {
	PublishItemCreated(message []byte) error
	PublishItemUpdated(message []byte) error
	PublishItemDeleted(id int) error
}
