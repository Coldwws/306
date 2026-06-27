package service

import "github.com/Coldwws/306/internal/models"

type RoomService interface {
	GetAllRooms() ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
	CreateRoom(room models.Room) (int, error)
	DeleteRoom(id int) (int, error)
	UpdateRoom(id int, updateRoom models.UpdateRoom) error
}

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
}
