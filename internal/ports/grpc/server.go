package grpc

import (
	"fmt"
	"net"

	pb "github.com/CafeKetab/PBs/golang/auth"
	"github.com/CafeKetab/auth/pkg/crypto"
	"github.com/CafeKetab/auth/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	config *Config
	logger *zap.Logger
	crypto crypto.Crypto
	token  token.Token

	api *grpc.Server
	pb.UnimplementedAuthServer
}

func New(cfg *Config, log *zap.Logger, c crypto.Crypto, t token.Token) *server {
	s := &server{config: cfg, logger: log, crypto: c, token: t}

	s.api = grpc.NewServer()
	pb.RegisterAuthServer(s.api, s)

	return s
}

func (s *server) Serve() {
	address := fmt.Sprintf(":%d", s.config.ListenPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Panic("Error listening on tcp address", zap.Int("port", s.config.ListenPort), zap.Error(err))
	}

	if err := s.api.Serve(listener); err != nil {
		s.logger.Fatal("Error serving gRPC server", zap.Error(err))
	}
}
