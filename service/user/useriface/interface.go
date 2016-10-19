package useriface

import "../../user"

type UserAPI interface {
	GetUsers(input *user.GetUsersInput) (output *user.UserOutputList)
	GetUser(id int) (output *user.UserOutput)
	CreateUser(input *user.CreateUserInput) (output *user.UserOutput)
	UpdateUser(input *user.UpdateUserInput) (output *user.UserOutput)
	DeleteUser(id int) (output bool)
	test()
}

var _ UserAPI = (*user.Service)(nil)
