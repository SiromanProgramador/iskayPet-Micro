package main

import (
	"context"
	"fmt"

	//PETS domain
	PetsRepo "iskayPetMicro/Domains/pets/entity/repository/mongodb"
	PetsInterface "iskayPetMicro/Domains/pets/interface"
	PetsUsecase "iskayPetMicro/Domains/pets/usecase"

	//SPECIES domain
	SpeciesRepo "iskayPetMicro/Domains/species/entity/repository/mongodb"

	"log"
	"net"

	pb "iskayPetMicro/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/mgo.v2"
)

//server struct
type server struct {
	pb.UnimplementedCreatePetServiceServer
	pb.UnimplementedGetPetsServiceServer
	pb.UnimplementedGetStatisticsServiceServer
}

var db *mgo.Database

//function to load Pet repositories to be called by all microservices methods
func loadPetsRepositories() PetsInterface.InterfaceInterface {
	petsRepo := PetsRepo.NewMongoDBRepository(db.Session)
	speciesRepo := SpeciesRepo.NewMongoDBRepository(db.Session)
	petsUsecase := PetsUsecase.NewUsecase(petsRepo, speciesRepo)
	petsInterface := PetsInterface.PetsInterface(petsUsecase)

	return petsInterface
}

//CreatePet server Funciton
func (s *server) CreatePet(ctx context.Context, in *pb.CreatePetRequest) (*pb.CreatePetReply, error) {

	//charge Pet Interface with repositories
	petsInterface := loadPetsRepositories()

	response, err := petsInterface.CreatePet(*in.Pet)

	if err != nil {
		return &pb.CreatePetReply{Pet: response}, err
	} else {
		return &pb.CreatePetReply{Pet: response}, nil
	}
}

//GetPet Server Funciton
func (s *server) GetPets(ctx context.Context, in *pb.GetPetsRequest) (*pb.GetPetsReply, error) {

	//charge Pet Interface with repositories
	petsInterface := loadPetsRepositories()

	response, err := petsInterface.GetAllPets(in.Filter)

	if err != nil {
		return &pb.GetPetsReply{Pets: response}, err
	} else {
		return &pb.GetPetsReply{Pets: response}, nil
	}
}

//GetStatistics server Funcion
func (s *server) GetStatistics(ctx context.Context, in *pb.GetStatisticsRequest) (*pb.GetStatisticsReply, error) {

	//charge Pet Interface with repositories
	petsInterface := loadPetsRepositories()

	response, err := petsInterface.GetStatistics(in.PetName)

	if err != nil {
		return &pb.GetStatisticsReply{Statistics: response}, err
	} else {
		return &pb.GetStatisticsReply{Statistics: response}, nil
	}
}

func main() {
	//Init DataBase
	db = MongoStart()

	// create a listener on TCP port 7770
	log.Println("[START] ISKAYPET Micro")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7770))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	// start the server

	//pets register services
	pb.RegisterCreatePetServiceServer(grpcServer, &s)
	pb.RegisterGetPetsServiceServer(grpcServer, &s)
	pb.RegisterGetStatisticsServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func MongoStart() *mgo.Database {
	session, err :=
		mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("challengeDB")

	return db

}
