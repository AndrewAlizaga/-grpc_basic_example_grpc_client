package accounts

import (
	"context"
	"log"
	"time"

	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
)


func AccountLogin(request *accountapiv1.LoginRequest) (response *accountapiv1.LoginResponse, err error){
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	
	conn := v1.AccountConnection()

	response, err = conn.AccountLogin(ctx, request)

	log.Println("Time for AccountLogin - ", time.Since(now))
	return
}

func AccountSignUp(request *accountapiv1.SignUpRequest) (response *accountapiv1.SignUpResponse, err error){
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	
	conn := v1.AccountConnection()

	response, err = con.AccountLogin(ctx, request)

	log.Println("Time for AccountLogin - ", time.Since(now))
	return
}
