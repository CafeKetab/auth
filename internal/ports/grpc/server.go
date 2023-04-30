package grpc

import (
	"context"
	"errors"
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
	crypto crypto.Crypto
	token  token.Token

	server *grpc.Server
	pb.UnimplementedAuthServer
}

func New(logger *zap.Logger, crypto crypto.Crypto, token token.Token) ports.Server {
	s := &server{logger: logger, crypto: crypto, token: token}

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
	if pbId.Value == 0 {
		errString := "Invalid id provider"
		s.logger.Error(errString)
		return nil, errors.New(errString)
	}

	payload, err := s.crypto.Encrypt(fmt.Sprint(pbId.Value))
	if err != nil {
		errString := "Error while encrypting the id"
		s.logger.Error(errString, zap.Error(err))
		return nil, errors.New(errString)
	}

	token, err := s.token.CreateTokenString(payload)
	if err != nil {
		errString := "Error while creating token from the payload"
		s.logger.Error(errString, zap.Error(err))
		return nil, errors.New(errString)
	}

	return &pb.Token{Value: token}, nil
}

func (s *server) GetIdFromToken(ctx context.Context, pbToken *pb.Token) (*pb.Id, error) {
	var payload string
	if err := s.token.ExtractTokenData(pbToken.Value, &payload); err != nil {
		errString := "Error while extracting payload from token"
		s.logger.Error(errString, zap.String("token", pbToken.Value), zap.Error(err))
		return nil, errors.New(errString)
	}

	stringId, err := s.crypto.Decrypt(payload)
	if err != nil {
		errString := "Error while decrypting the payload"
		s.logger.Error(errString, zap.String("payload", payload), zap.Error(err))
		return nil, errors.New(errString)
	}

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		errString := "Invalid id has been given"
		s.logger.Error(errString, zap.String("id", stringId), zap.Error(err))
		return nil, errors.New(errString)
	}

	return &pb.Id{Value: id}, nil
}
