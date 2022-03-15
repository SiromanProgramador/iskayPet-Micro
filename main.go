package main

import (
	"context"
	"fmt"
	PetsRepo "iskayPetMicro/Domains/pets/entity/repository/mongodb"
	PetsInterface "iskayPetMicro/Domains/pets/interface"
	PetsUsecase "iskayPetMicro/Domains/pets/usecase"
	SpeciesRepo "iskayPetMicro/Domains/species/entity/repository/mongodb"
	"log"
	"net"

	// gin-swagger middleware
	// swagger embed files

	pb "iskayPetMicro/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/mgo.v2"
)

type server struct {
	pb.UnimplementedCreatePetServiceServer
	pb.UnimplementedGetPetsServiceServer
	pb.UnimplementedGetStatisticsServiceServer
}

var db *mgo.Database

func loadPetsRepositories() PetsInterface.InterfaceInterface {
	petsRepo := PetsRepo.NewMongoDBRepository(db.Session)
	speciesRepo := SpeciesRepo.NewMongoDBRepository(db.Session)
	petsUsecase := PetsUsecase.NewUsecase(petsRepo, speciesRepo)
	petsInterface := PetsInterface.PetsInterface(petsUsecase)

	return petsInterface
}

func (s *server) CreatePet(ctx context.Context, in *pb.CreatePetRequest) (*pb.CreatePetReply, error) {

	petsInterface := loadPetsRepositories()

	response, err := petsInterface.CreatePet(*in.Pet)

	if err != nil {
		return &pb.CreatePetReply{Pet: response}, err
	} else {
		return &pb.CreatePetReply{Pet: response}, nil
	}
}

func (s *server) GetPets(ctx context.Context, in *pb.GetPetsRequest) (*pb.GetPetsReply, error) {

	petsInterface := loadPetsRepositories()

	response, err := petsInterface.GetAllPets(in.Filter)

	if err != nil {
		return &pb.GetPetsReply{Pets: response}, err
	} else {
		return &pb.GetPetsReply{Pets: response}, nil
	}
}

func (s *server) GetStatistics(ctx context.Context, in *pb.GetStatisticsRequest) (*pb.GetStatisticsReply, error) {

	petsInterface := loadPetsRepositories()

	response, err := petsInterface.GetStatistics(in.PetName)

	if err != nil {
		return &pb.GetStatisticsReply{Statistics: response}, err
	} else {
		return &pb.GetStatisticsReply{Statistics: response}, nil
	}
}

func main() {
	db = MongoStart()
	//Init DataBase
	//Boot.Boot(db.Session)
	// create a listener on TCP port 7777
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

	//pets
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
