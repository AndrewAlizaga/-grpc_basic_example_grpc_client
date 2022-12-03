package connection

import (
	"log"

	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
	"google.golang.org/grpc"
)


var agentServerConnection accountapiv1.AccountServiceClient


func GetAccountConnection() (accountapiv1.AccountServiceClient, error){
	
	var err error

	if agentServerConnection == nil {

		// Set up a connection to the server.
		var conn *grpc.ClientConn
		conn, err = grpc.Dial("localhost:8080", grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Error connecting with Configuration service: %v", err)
		}

		agentServerConnection = accountapiv1.NewAccountServiceClient(conn)

	}


	return agentServerConnection, err
}