package service

import (
	"github.com/Coldwws/306/internal/models"
	"github.com/Coldwws/306/internal/repository"
)

type roomService struct {
	storage repository.RoomRepository
}

func NewRoomImplementation(storage repository.RoomRepository) *roomService {
	return &roomService{storage: storage}
}

func (service *roomService) GetAllRooms() ([]models.Room, error) {
	return service.storage.GetAllRooms()
}

func (service *roomService) CreateRoom(room models.Room) (int, error) {
	return service.storage.CreateRoom(room)
}

func (service *roomService) UpdateRoom(id int, updateRoom models.UpdateRoom) error {
	return service.storage.UpdateRoom(id, updateRoom)
}

func (service *roomService) DeleteRoom(id int) (int, error) {
	return service.storage.DeleteRoom(id)
}

func (service *roomService) GetRoomById(id int) (models.Room, error) {
	return service.storage.GetRoomById(id)
}
