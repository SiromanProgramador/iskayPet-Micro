package _interface

import (
	Usecase "iskayPetMicro/Domains/pets/usecase"
	pb "iskayPetMicro/api"
	"iskayPetMicro/model"
)

type InterfaceInterface interface {
	GetAllPets(filter string) ([]*pb.Pet, error)
	GetStatistics(petName string) (*pb.ResponseStatistics, error)
	CreatePet(pb.Pet) (*pb.Pet, error)
}

type Interface struct {
	usecase Usecase.UsecaseInterface
}

func PetsInterface(usecase Usecase.UsecaseInterface) InterfaceInterface {
	return &Interface{
		usecase: usecase,
	}
}

func (ui *Interface) GetAllPets(filter string) ([]*pb.Pet, error) {

	var qfilter model.QueryFilters

	pets, errPets := ui.usecase.GetAllPets(qfilter)

	if errPets != nil {
		return pets, errPets
	}

	return pets, errPets
}

func (ui *Interface) GetStatistics(petName string) (*pb.ResponseStatistics, error) {

	//SET filter
	var qfilter model.QueryFilters

	objectToReturn, errorFind := ui.usecase.GetStatistics(qfilter)
	if errorFind != nil {
		return &objectToReturn, errorFind
	}

	return &objectToReturn, errorFind
}

func (ui *Interface) CreatePet(objectToCreate pb.Pet) (*pb.Pet, error) {

	response, errCreate := ui.usecase.CreatePet(objectToCreate)
	if errCreate != nil {

		return &response, errCreate
	}

	return &response, errCreate
}
