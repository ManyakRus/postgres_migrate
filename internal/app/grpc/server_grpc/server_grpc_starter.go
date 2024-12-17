package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/stopapp"
	"google.golang.org/grpc"
	"net"
	"os"
)

// ServerGRPC is used to implement UnimplementedPostgresMigrateServer.
type ServerGRPC struct {
	grpc_proto.UnimplementedPostgresMigrateServer
}

type SettingsINI struct {
	GRPC_PORT string
}

var Settings SettingsINI

var Conn *grpc.Server

func Connect() {
	//var err error

	if Settings.GRPC_PORT == "" {
		FillSettings()
	}

	Conn = grpc.NewServer()
	grpc_proto.RegisterPostgresMigrateServer(Conn, &ServerGRPC{})

	stopapp.GetWaitGroup_Main().Add(1)
	go serve_go()

}

func serve_go() {
	defer stopapp.GetWaitGroup_Main().Done()

	lis, err := net.Listen("tcp", ":"+Settings.GRPC_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server GRPC listening at %v", lis.Addr())
	if err := Conn.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
}

func FillSettings() {
	Settings = SettingsINI{}
	Settings.GRPC_PORT = os.Getenv("GRPC_PORT")
	if Settings.GRPC_PORT == "" {
		log.Panic("Need fill GRPC_PORT ! in OS Environment ")
	}

}

// WaitStop - ожидает отмену глобального контекста
func WaitStop() {
	defer stopapp.GetWaitGroup_Main().Done()

	select {
	case <-contextmain.GetContext().Done():
		log.Warn("Context app is canceled. grpc_connect")
	}

	//ждём пока отправляемых сейчас сообщений будет =0
	stopapp.WaitTotalMessagesSendingNow("grpc_connect")

	//закрываем соединение
	CloseConnection()

}

// Start - необходимые процедуры для запуска сервера GRPC
func Start() {
	Connect()

	stopapp.GetWaitGroup_Main().Add(1)
	go WaitStop()

}

func CloseConnection() {
	Conn.Stop()
}
