package application

import (
	"jwt-auth-with-gorm/database"
	"jwt-auth-with-gorm/models"
	"jwt-auth-with-gorm/models/managers"
)

func Get(id int) (*models.Book, error) {
	conn, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	manager := managers.NewBookManager(conn)
	return manager.Get(id)
}
