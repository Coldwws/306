package repository

import "github.com/Coldwws/306/internal/models"

type RoomRepository interface {
	GetAllRooms() ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
	CreateRoom(room models.Room) (int, error)
	DeleteRoom(id int) (int, error)
	UpdateRoom(id int, updateRoom models.UpdateRoom) error
}
