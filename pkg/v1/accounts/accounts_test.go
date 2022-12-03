package accounts_test

import (
	"log"
	"testing"

	accountpkg "github.com/AndrewAlizaga/-grpc_basic_example_grpc_client/pkg/v1/accounts"
	accountv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/account"
	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
)

func TestAccountSignUp(t *testing.T) {

	log.Println("INVOKE - TestAccountSignUp")

	request := accountapiv1.SignUpRequest{
		Account: &accountv1.Account{
			Name:     "and",
			Email:    "and1@gmail.com",
			Password: "qwerty",
		},
	}

	res, err := accountpkg.AccountSignUp(&request)

	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	} else {
		if res.GetError() != "" {
			t.Logf(res.GetError())
			t.FailNow()
		}
	}

	t.Log("test completed successfully")
	t.Log(res.GetJwt())

}

func TestAccountLogin(t *testing.T) {

	log.Println("INVOKE - TestAccountSignUp")

	request := accountapiv1.LoginRequest{
		Email:    "day1@gmail.com",
		Password: "qwerty",
	}

	res, err := accountpkg.AccountLogin(&request)

	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	} else {
		if res.GetError() != "" {
			t.Logf(res.GetError())
			t.FailNow()
		}
	}

	t.Log("test completed successfully")
	t.Log(res.GetJwt())

}
