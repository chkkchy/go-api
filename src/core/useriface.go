package core

type User interface {
	GetUsers(input *GetUsersInput) (output *UserOutputList)

	GetUser(id int) (output *UserOutput)

	CreateUser(input *CreateUserInput) (output *UserOutput)

	UpdateUser(input *UpdateUserInput) (output *UserOutput)

	DeleteUser(id int) (output bool)
}
