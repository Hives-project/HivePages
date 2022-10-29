package pb

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *SongServerRequest) (*SongServerRequest, error) {
	log.Printf(" Received  message body from client: %s", message.SongName)
	return &SongServerRequest{SongName: "name of thing"}, nil
}
