package grpcserver

import (
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	srv     *grpc.Server
	address string
	notify  chan error
}

// Регистрирует grpc сервер и запускает прослушивание порта
func New(grpcServer *grpc.Server, address string) *GrpcServer {
	s := &GrpcServer{
		srv:     grpcServer,
		address: address,
		notify:  make(chan error, 1),
	}

	s.start()

	return s
}

func (s *GrpcServer) start() {
	go func() {
		s.notify <- s.serveAndListen(s.address)
		close(s.notify)
	}()
}

func (s *GrpcServer) Notify() <-chan error {
	return s.notify
}

func (s *GrpcServer) serveAndListen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "failed to listen address")
	}

	err = s.srv.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "failed to serve listener")
	}

	return nil
}

func (s *GrpcServer) Stop() {
	s.srv.Stop()
}
