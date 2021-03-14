package users

import "manajemen-komponen-api/models/usersModel"

// UsersUsecase - usecase users interface
type UsersUsecase interface {
	LoginUser(data *usersModel.ReqLogin) (*usersModel.ResLogin, error)
	CreateUser(data *usersModel.ReqInsertUser) (uint, error)
	GetAllUsers() (*[]usersModel.ResGetUsers, error)
	GetUserByID(ID int) (*usersModel.ResGetUsers, error)
	EditUser(data *usersModel.ReqUpdateUser) error
	DropUser(ID int) error
	GetRoles() (*[]usersModel.Roles, error)
	GetRoleByID(ID int) (*usersModel.Roles, error)
	UpdateCredential(ID int, data *usersModel.ReqChangePassword) error
}
