package main
import (
	"fmt"
	pb "github.com/jamwujustyle/gogrpc/greet/proto" 
)

func (s *Server) (in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("greetmanytimes invoked with %v", in)


	for i:= 0; i < 10; i++ {
		res:= fmt.Sprintf("greeted %s %d many tiems", in.FirstName, i)
	}
	return nil
}

