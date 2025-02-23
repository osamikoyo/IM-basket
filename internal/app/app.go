package app

import (
	"fmt"
	"net"

	"github.com/osamikoyo/IM-basket/internal/config"
	"github.com/osamikoyo/IM-basket/internal/data"
	"github.com/osamikoyo/IM-basket/internal/rpc"
	"github.com/osamikoyo/IM-basket/internal/server"
	"github.com/osamikoyo/IM-basket/pkg/loger"
	"github.com/osamikoyo/IM-basket/pkg/proto/pb"
	"google.golang.org/grpc"
)

type App struct{
	rpc *rpc.RpcServer
	gRPC *grpc.Server
	cfg *config.Config
	loger loger.Logger
	server *server.Server
}

func Init() (*App, error) {
	cfg, err := config.Load("config.yml")
	if err != nil{
		return nil, err
	}

	st, err := data.New(cfg)
	if err != nil{
		return nil,err
	}

	grpcServer := grpc.NewServer()
	server := &server.Server{
		Storage: st,
	}
	rpc, err := rpc.New(cfg)
	if err != nil{
		return nil, err
	}

	return &App{
		rpc: rpc,
		gRPC: grpcServer,
		server: server,
		cfg: cfg,
		loger: loger.New(),
	}, nil
}

func (a *App) Run() error {
	a.loger.Info().Msg("Starting Basket Service...")
	var ch chan error

	go a.rpc.Listen(ch)

	go func ()  {
		err := <- ch
		a.loger.Error().Err(err)	
	}()

	pb.RegisterBasketServiceServer(a.gRPC, a.server)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port))
	if err != nil{
		return err
	}

	a.loger.Info().Str("addr", lis.Addr().String()).Msg("Server started!")
	return a.gRPC.Serve(lis)
}