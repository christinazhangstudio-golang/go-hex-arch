package main

//https://www.youtube.com/watch?v=MpFog2kZsHk

import (
	"os"
	"log"
	"github.tesla.com/chrzhang/go-hex-arch/internal/adapters/app/api"
	"github.tesla.com/chrzhang/go-hex-arch/internal/adapters/core/arithmetic"
	"github.tesla.com/chrzhang/go-hex-arch/internal/adapters/framework/right/db"
	"github.tesla.com/chrzhang/go-hex-arch/internal/ports"

	gRPC "github.tesla.com/chrzhang/go-hex-arch/internal/adapters/framework/left/grpc"
)

func main() {
	var err error
	//ports
	var core ports.ArithmeticPort
	var dbaseAdapter ports.DbPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	//plugging in adapters to ports
	//database adapter needs to be injected to application layer, 
	//and app layer needs to be injected to grpc adapter

	//arguments to new adapter for db will come from ENV
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	//plug in arithmetic adapter
	core = arithmetic.NewAdapter()

	//plug in app adapter
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	//plug in grpc adapter
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	//start up grpc server
	gRPCAdapter.Run()


	// arithAdapter := arithmetic.NewAdapter()
	// result, err := arithAdapter.Addition(1, 3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
}