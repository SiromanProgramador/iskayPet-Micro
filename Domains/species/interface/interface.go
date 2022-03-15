package _interface

import (
	Usecase "iskayPetMicro/Domains/species/usecase"
)

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
