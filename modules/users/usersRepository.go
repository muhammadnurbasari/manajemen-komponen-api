package users

import "manajemen-komponen-api/models/usersModel"

// UsersRepository - repository users interface
type UsersRepository interface {
	RetrieveUsers() (*[]usersModel.ResGetUsers, error)
	RetrieveUserByEmail(email string) (*usersModel.Users, error)
	InsertUser(data *usersModel.ReqInsertUser) (uint, error)
	UpdateUser(data *usersModel.ReqUpdateUser) error
	RetrieveUserByID(ID int) (*usersModel.ResGetUsers, error)
	DeleteUser(ID int) error
	IsEmailExist(email string) (bool, error)
	RetrieveRoles() (*[]usersModel.Roles, error)
	RetrieveRoleByID(ID int) (*usersModel.Roles, error)
	UpdateCredentialByID(ID int, password string) error
	RetrieveCredentialByID(ID int) (string, error)
}
