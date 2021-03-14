package usersUsecase

import (
	"errors"
	"manajemen-komponen-api/middlewares/jwtToken"
	"manajemen-komponen-api/middlewares/validation"
	"manajemen-komponen-api/models/usersModel"
	"manajemen-komponen-api/modules/users"
)

type usersUsecase struct {
	usersRepository users.UsersRepository
}

//  NewUsersUsecase - create an object representation users.UsersUsecase interface
func NewUsersUsecase(usersRepository users.UsersRepository) users.UsersUsecase {
	return &usersUsecase{usersRepository: usersRepository}
}

// LoginUser - usecase login user
func (userUC *usersUsecase) LoginUser(data *usersModel.ReqLogin) (*usersModel.ResLogin, error) {
	// cek valid email
	isValidEmail := validation.IsEmailValid(data.Email)
	if !isValidEmail {
		// return error string "1" for is not valid email
		return nil, errors.New("1")
	}
	// get user by email
	dataUser, errGetUser := userUC.usersRepository.RetrieveUserByEmail(data.Email)
	if errGetUser != nil {
		if errGetUser.Error() == "2" {
			// return error string "2" for is not exist email
			return nil, errors.New("2")
		}
		return nil, errors.New("usersUsecase.LoginUser err = " + errGetUser.Error())
	}

	// check password
	checkPwdOK := validation.CheckPasswordHash(data.Password, dataUser.Password)

	// generate token jwt
	jwtTokenString, errGenerateToken := jwtToken.GenerateToken(int(dataUser.ID))
	if errGenerateToken != nil {
		return nil, errors.New("usersUsecase.LoginUser err = " + errGenerateToken.Error())
	}

	if !checkPwdOK {
		// return error string "3" for incorect password
		return nil, errors.New("3")
	}

	dataResp := usersModel.ResLogin{
		ID:     dataUser.ID,
		RoleID: dataUser.RoleID,
		Token:  jwtTokenString,
	}

	return &dataResp, nil

}

// CreateUser - usecase create user
func (userUC *usersUsecase) CreateUser(data *usersModel.ReqInsertUser) (uint, error) {
	// cek valid email
	isValidEmail := validation.IsEmailValid(data.Email)
	if !isValidEmail {
		// return error string "1" for is not valid email
		return 0, errors.New("1")
	}

	// check max length an min length password
	isOkLength := validation.MoreThanMaxLength(data.Password)
	if !isOkLength {
		// return error string "2" for password cant more than 15 char & min than 5 char
		return 0, errors.New("2")
	}
	// cek email exist
	isEmailExist, errExistEmail := userUC.usersRepository.IsEmailExist(data.Email)
	if errExistEmail != nil {
		return 0, errors.New("usersUsecase.CreateUser err = " + errExistEmail.Error())
	}

	if isEmailExist {
		// return error string "3" for email is exist
		return 0, errors.New("3")
	}

	// hash password
	hashString, errHash := validation.HashPassword(data.Password)
	if errHash != nil {
		return 0, errors.New("usersUsecase.CreateUser err = " + errHash.Error())
	}

	dataInsert := usersModel.ReqInsertUser{
		Email:    data.Email,
		Username: data.Username,
		Password: hashString,
		RoleID:   data.RoleID,
	}

	// insert to table user
	ID, errInsert := userUC.usersRepository.InsertUser(&dataInsert)
	if errInsert != nil {
		return 0, errors.New("usersUsecase.CreateUser err = " + errInsert.Error())
	}

	return ID, nil

}

// GetAllUsers - usecase get all data users
func (userUC *usersUsecase) GetAllUsers() (*[]usersModel.ResGetUsers, error) {
	dataUsers, err := userUC.usersRepository.RetrieveUsers()

	if err != nil {
		return nil, errors.New("usersUsecase.GetAllUsers err = " + err.Error())
	}

	return dataUsers, nil
}

// GetUserByID - usecase get user by ID
func (userUC *usersUsecase) GetUserByID(ID int) (*usersModel.ResGetUsers, error) {
	data, err := userUC.usersRepository.RetrieveUserByID(ID)

	if err != nil {
		if err.Error() == "1" {
			// return error string "1" for ID is not Exist
			return nil, errors.New("1")
		}
		return nil, errors.New("usersUsecase.GetUserByID err = " + err.Error())
	}

	return data, nil
}

// EditUser - usecase update user by id
func (userUC *usersUsecase) EditUser(data *usersModel.ReqUpdateUser) error {
	err := userUC.usersRepository.UpdateUser(data)

	if err != nil {
		if err.Error() == "1" {
			// return error string "1" for ID is not Exist
			return errors.New("1")
		}
		return errors.New("usersUsecase.EditUser err = " + err.Error())
	}

	return nil
}

// DropUser - usecase delete user by ID
func (userUC *usersUsecase) DropUser(ID int) error {
	err := userUC.usersRepository.DeleteUser(ID)

	if err != nil {
		if err.Error() == "1" {
			// return error string "1" for ID is not Exist
			return errors.New("1")
		}
	}

	return nil
}

// GetRoles - usecase get roles
func (userUC *usersUsecase) GetRoles() (*[]usersModel.Roles, error) {
	data, err := userUC.usersRepository.RetrieveRoles()

	if err != nil {
		return nil, errors.New("usersUsecase.GetRoles err = " + err.Error())
	}

	return data, nil
}

// GetRoleByID - usecase get role by id
func (userUC *usersUsecase) GetRoleByID(ID int) (*usersModel.Roles, error) {
	data, err := userUC.usersRepository.RetrieveRoleByID(ID)

	if err != nil {
		if err.Error() == "1" {
			// return error string "1" for ID not exist
			return nil, errors.New("1")
		}
	}

	return data, nil

}

// UpdateCredential - usecase change password
func (userUC *usersUsecase) UpdateCredential(ID int, data *usersModel.ReqChangePassword) error {
	// get credential by id
	passHash, errGetCredential := userUC.usersRepository.RetrieveCredentialByID(ID)
	if errGetCredential != nil {
		return errors.New("usersUsecase.UpdateCredential err = " + errGetCredential.Error())
	}

	// check old password is matches
	isMatches := validation.CheckPasswordHash(data.OldPassword, passHash)

	if !isMatches {
		// return error string "1" for old password is not matches
		return errors.New("1")
	}

	// hashing new password
	newPass, errHash := validation.HashPassword(data.NewPassword)
	if errHash != nil {
		return errors.New("usersUsecase.UpdateCredential err = " + errHash.Error())
	}

	// update password
	errUpdate := userUC.usersRepository.UpdateCredentialByID(ID, newPass)
	if errUpdate != nil {
		return errors.New("usersUsecase.UpdateCredential err = " + errUpdate.Error())
	}

	return nil
}
