package usecase

import (
	"iskayPetMicro/Domains/pets/entity/repository"
	SpeciesRepo "iskayPetMicro/Domains/species/entity/repository"
	pb "iskayPetMicro/api"
	"iskayPetMicro/model"
	"iskayPetMicro/presenters"
	"math"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type UsecaseInterface interface {
	CreatePet(pet pb.Pet) (pb.Pet, error)
	GetStatistics(queryFilters model.QueryFilters) (pb.ResponseStatistics, error)
	GetAllPets(filter model.QueryFilters) ([]*pb.Pet, error)
}

//import both repositories, pets and species to be able to use species methods
type Usecase struct {
	repo        repository.RepositoryInterface
	speciesRepo SpeciesRepo.RepositoryInterface
}

func NewUsecase(
	repo repository.RepositoryInterface,
	speciesRepo SpeciesRepo.RepositoryInterface,
) UsecaseInterface {
	return &Usecase{
		repo:        repo,
		speciesRepo: speciesRepo,
	}
}

func (u *Usecase) CreatePet(protoObjectToCreate pb.Pet) (pb.Pet, error) {
	var objectToCreate model.Pet
	var qFilter model.QueryFilters
	var species model.Species
	var errSpecies error

	//parsed name to lower case to no duplicate information. Example: cat - Cat - cAt - caT
	petNameParsed := strings.ToLower(protoObjectToCreate.Species)

	//check if species is in db
	qFilter.Filter = bson.M{"name": petNameParsed}
	species, errSpecies = u.speciesRepo.GetOne(&qFilter)

	//if errSpecies has a different error that not found, return
	if errSpecies != nil && errSpecies.Error() != "not found" {
		return protoObjectToCreate, errSpecies
	}

	//if species is not found in db, create them
	if errSpecies != nil && errSpecies.Error() == "not found" {

		species.Id = bson.NewObjectId()
		species.Instance = presenters.CreateInstance()
		species.Name = petNameParsed

		errCreateSpecies := u.speciesRepo.Create(&species)
		if errCreateSpecies != nil {
			return protoObjectToCreate, errCreateSpecies
		}
	}

	//parse object to insert in DB
	objectToCreate.Age = int(protoObjectToCreate.Age)
	objectToCreate.DateOfBird = protoObjectToCreate.DateOfBirth
	objectToCreate.Gender = protoObjectToCreate.Gender
	objectToCreate.Name = protoObjectToCreate.Name

	objectToCreate.SpeciesId = species.Id

	objectToCreate.Id = bson.NewObjectId()
	objectToCreate.Instance = presenters.CreateInstance()

	//create pet
	errCreate := u.repo.Create(&objectToCreate)

	return protoObjectToCreate, errCreate
}

func (u *Usecase) GetStatistics(queryFilters model.QueryFilters) (pb.ResponseStatistics, error) {
	var response pb.ResponseStatistics
	var filter model.QueryFilters
	var maxPetsBySpeciesId bson.ObjectId
	var maxPetsBySpeciesName string
	var maxPetsNumber int
	var ages []int
	var sumAges int
	var sd float64

	//get All species calling species repository
	species, errSpecies := u.speciesRepo.GetAll(filter)
	if errSpecies != nil {
		return response, errSpecies
	}

	//init counters
	maxPetsNumber = 0

	//find what species has more pets in DB, for each specie get count
	for _, specie := range species {

		filter.Filter = bson.M{"speciesId": specie.Id}

		numPets, errPet := u.repo.Count(filter)
		if errPet != nil {
			return response, errPet
		}

		//if count pets by species is higher than  before count by species, set new value
		if numPets > maxPetsNumber {
			maxPetsNumber = numPets
			maxPetsBySpeciesId = specie.Id
			maxPetsBySpeciesName = specie.Name
		}
	}

	//get all pets of maximum species
	filter.Filter = bson.M{"speciesId": maxPetsBySpeciesId}
	pets, errPets := u.repo.GetAll(filter)

	//assign the maximum species name to response
	response.Species = maxPetsBySpeciesName

	//calculate  average age for this species
	for _, pet := range pets {
		ages = append(ages, pet.Age)
		sumAges += pet.Age
	}

	response.AverageAge = float64(sumAges) / float64(len(ages))

	//calculate standard deviation
	for _, age := range ages {
		sd += math.Pow(float64(age)-response.AverageAge, 2)
	}

	sd = math.Sqrt(sd / float64(len(ages)))

	response.StandardDeviation = sd

	return response, errPets
}

func (u *Usecase) GetAllPets(filter model.QueryFilters) ([]*pb.Pet, error) {

	var response []*pb.Pet
	pets, errPets := u.repo.GetAll(filter)

	//parse mongodb object to protobuf object because  in protobuf file we can't use ObjectId
	//and we need the species name in the response
	for _, pet := range pets {

		var protoPet pb.Pet

		protoPet.Age = int32(pet.Age)
		protoPet.DateOfBirth = pet.DateOfBird
		protoPet.Species = pet.SpeciesInfo.Name
		protoPet.Gender = pet.Gender
		protoPet.Name = pet.Name
		response = append(response, &protoPet)
	}

	return response, errPets
}
