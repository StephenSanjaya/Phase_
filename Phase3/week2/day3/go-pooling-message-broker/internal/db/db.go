package db

import (
	"Phase3/week2/day3/go-pooling-message-broker/internal/app/interfaces"
	"Phase3/week2/day3/go-pooling-message-broker/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

var _ interfaces.Datastore = &Database{}

// Init Database
func NewDatabase(dsn string) (*Database, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
}

func (db *Database) GetAllItems() ([]models.Item, error) {
	items := []models.Item{}
	rows, err := db.Pool.Query(context.Background(), "SELECT id, name, price FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {

			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (db *Database) GetItemByID(id int) (*models.Item, error) {
	item := &models.Item{}
	err := db.Pool.QueryRow(context.Background(), "SELECT id, name, price FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (db *Database) InsertItem(item *models.Item) error {
	err := db.Pool.QueryRow(context.Background(), "INSERT INTO items (name, price) VALUES ($1, $2) RETURNING id", item.Name, item.Price).Scan(&item.ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (db *Database) UpdateItem(id int, item *models.Item) error {
	commandTag, err := db.Pool.Exec(context.Background(), "UPDATE items SET name = $1, price = $2 WHERE id = $3", item.Name, item.Price, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (db *Database) DeleteItem(id int) error {
	commandTag, err := db.Pool.Exec(context.Background(), "DELETE FROM items WHERE id = $1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
