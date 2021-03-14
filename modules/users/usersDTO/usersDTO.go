package usersDTO

import "manajemen-komponen-api/models/usersModel"

// ResLoginResult - response example login module
type ResLoginResult struct {
	Status    int                 `json:"status" example:"200"`
	Messagges string              `json:"messages" example:"Login Success"`
	Result    usersModel.ResLogin `json:"result"`
}

//ReqLoginDTO - request json login users
type ReqLoginDTO struct {
	Email    string `json:"email" example:"xxxxxx@xxxxxx.com"`
	Password string `json:"password" example:"123456123"`
}

// ResInsertUser - response insert users
type ResInsertUser struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"User has been Created"`
	ID        int    `json:"user_id" example:"1"`
}

// ResUpdateUser - response update users
type ResUpdateUser struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"User has been Updated"`
}

// ResDeleteUserByID - response update users
type ResDeleteUserByID struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"User has been Deleted"`
}

// ResGetAllUsers - response get users
type ResGetAllUsers struct {
	Status    int                      `json:"status" example:"200"`
	Messagges string                   `json:"messages" example:"Success"`
	Result    []usersModel.ResGetUsers `json:"result"`
}

// ResGetUserByID - response get user by id
type ResGetUserByID struct {
	Status    int                    `json:"status" example:"200"`
	Messagges string                 `json:"messages" example:"Success"`
	Result    usersModel.ResGetUsers `json:"result"`
}

// ResGetRoleByID - response get role by id
type ResGetRoleByID struct {
	Status    int              `json:"status" example:"200"`
	Messagges string           `json:"messages" example:"Success"`
	Result    usersModel.Roles `json:"result"`
}

// ResGetRoles - response get roles
type ResGetRoles struct {
	Status    int                `json:"status" example:"200"`
	Messagges string             `json:"messages" example:"Success"`
	Result    []usersModel.Roles `json:"result"`
}

// ResChangePassword - response change password
type ResChangePassword struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"Change Password Success"`
}
