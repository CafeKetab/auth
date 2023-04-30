package grpc

import (
	"context"
	"fmt"
	"net"
	"strconv"

	pb "github.com/CafeKetab/PBs/golang/auth"
	"github.com/CafeKetab/auth-go/internal/ports"
	"github.com/CafeKetab/auth-go/pkg/crypto"
	"github.com/CafeKetab/auth-go/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	logger *zap.Logger
	token  token.Token
	crypto crypto.Crypto

	server *grpc.Server
}

func NewServer(logger *zap.Logger, token token.Token, crypto crypto.Crypto) ports.Server {
	s := &server{logger: logger, token: token, crypto: crypto}

	s.server = grpc.NewServer()
	pb.RegisterAuthServer(s.server, s)

	return s
}

func (s *server) Serve(port int) {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Panic("Error listening for gRPC server", zap.Int("port", port), zap.Error(err))
	}

	if err := s.server.Serve(listener); err != nil {
		s.logger.Panic("Error serving gRPC server", zap.Error(err))
	}
}

func (s *server) CreateTokenFromId(ctx context.Context, pbId *pb.Id) (*pb.Token, error) {
	payload, err := s.crypto.Encrypt(fmt.Sprint(pbId.Value))
	if err != nil {

	}

	token, err := s.token.CreateTokenString(payload)
	if err != nil {

	}

	return &pb.Token{Value: token}, nil
}

func (s *server) GetIdFromToken(ctx context.Context, pbToken *pb.Token) (*pb.Id, error) {
	var payload string
	if err := s.token.ExtractTokenData(pbToken.Value, &payload); err != nil {

	}

	stringId, err := s.crypto.Decrypt(payload)
	if err != nil {

	}

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {

	}

	return &pb.Id{Value: id}, nil
}
