package repository

import (
	"fmt"

	"github.com/Coldwws/306/internal/models"
	"github.com/jmoiron/sqlx"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{
		db: db,
	}
}

func (r *RoomPostgres) GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	query := `SELECT * FROM rooms`
	err := r.db.Select(&rooms, query)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomPostgres) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	query := `SELECT * FROM rooms WHERE id = $1`
	err := r.db.Get(&room, query, id)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func (r *RoomPostgres) CreateRoom(room models.Room) (int, error) {
	var id int
	query := `INSERT INTO rooms(number,type,description) VALUES ($1,$2,$3) RETURNING id`
	err := r.db.QueryRow(query, room.Number, room.Type, room.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RoomPostgres) DeleteRoom(id int) (int, error) {
	query := `DELETE FROM rooms WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("Failed to delete room: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Could not get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("Room with id: %d not found", id)
	}
	return id, nil
}

func (r *RoomPostgres) UpdateRoom(id int, updateRoom models.UpdateRoom) error {
	query := `UPDATE rooms SET number = $1, type = $2, description = $3 WHERE id = $4`
	_, err := r.db.Exec(query, updateRoom.Number, updateRoom.Type, updateRoom.Description, id)
	if err != nil {
		return fmt.Errorf("Failed to update room: %w", err)
	}
	return nil
}
