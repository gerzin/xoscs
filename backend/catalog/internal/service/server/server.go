package server

import (
	"context"

	pb "github.com/gerzin/xoscs/backend/catalog/proto"
)

type Server struct{}

func (s *Server) GetAllScenarios(ctx context.Context, in *pb.GetAllScenariosRequest) (*pb.GetAllScenariosResponse, error) {

}

func (s *Server) DeleteScenario(ctx context.Context, in *pb.DeleteScenarioRequest) (*pb.DeleteScenarioResponse, error) {

}