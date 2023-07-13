package service

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type DBService struct {
	DB *pgx.Conn
}

func NewDBService(db *pgx.Conn) *DBService {
	return &DBService{
		DB: db,
	}
}

func (d *DBService) SavePaymentInfo(paymentInfo interface{}) error {
	// Insert the payment information into the database.
	// This is just a placeholder - you'll need to implement this based on your database schema.
	_, err := d.DB.Exec(context.Background(), "INSERT INTO payments VALUES ($1)", paymentInfo)
	return err
}
