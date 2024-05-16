package user

import (
	userModel "github.com/triwira-joel/simple-wallet/private/adapter/repository/db/user"
)

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

var users = []userModel.UserDB{{Id: 1, Name: "Ana"}, {Id: 2, Name: "Bobby"}}
