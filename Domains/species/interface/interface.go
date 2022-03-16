package _interface

import (
	Usecase "iskayPetMicro/Domains/species/usecase"
)

//This interface is empty because  in this challenge we dont use species interface layer
//But in this layer we recive client data, prepare and send this data to dhe next layer (Usecase layer)
type InterfaceInterface interface {
}

type Interface struct {
	usecase Usecase.UsecaseInterface
}

func SpeciesInterface(usecase Usecase.UsecaseInterface) InterfaceInterface {
	return &Interface{
		usecase: usecase,
	}
}
