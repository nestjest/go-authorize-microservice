package grpcapp

import "google.golang.org/grpc"

type App struct {
	log         *slog.Logger
	gRRPCServer *grpc.Server
	port        int
}

func New(log *slog.Logger, authService authgrpc.Auth, port int) *App {
	// TODO: Создать GRPC сервер и подключить к нему интерсепторы

	// TODO: зарегистрировать у сервера наш gRPC-сервис Auth

	// TODO: вернуть обьект App со всеми необходимыми полями
}
