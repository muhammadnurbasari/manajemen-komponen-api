package usersRepository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"manajemen-komponen-api/models/usersModel"
	"manajemen-komponen-api/modules/users"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type sqlRepository struct {
	Conn *gorm.DB
}

// NewUsersRepository - create an object representation users.UsersRepository interface
func NewUsersRepository(Conn *gorm.DB) users.UsersRepository {
	return &sqlRepository{Conn}
}

// RetrieveUsers - get all data users where deleted_at is NULL
func (db *sqlRepository) RetrieveUsers() (*[]usersModel.ResGetUsers, error) {
	var users []usersModel.Users

	data := db.Conn.Find(&users)

	if data.Error != nil {
		return nil, errors.New("usersRepository.RetrieveUsers err = " + data.Error.Error())
	}

	var dataResp []usersModel.ResGetUsers
	for _, list := range users {
		result := usersModel.ResGetUsers{
			ID:        list.ID,
			Email:     list.Email,
			Username:  list.Username,
			RoleID:    list.RoleID,
			CreatedAt: list.CreatedAt,
			UpdatedAt: list.UpdatedAt,
		}

		dataResp = append(dataResp, result)
	}

	return &dataResp, nil
}

// RetrieveUserByEmail - get users by email
func (db *sqlRepository) RetrieveUserByEmail(email string) (*usersModel.Users, error) {
	var users usersModel.Users

	// cek email is exist
	var count int
	isExist := db.Conn.Model(&usersModel.Users{}).Where("email = ?", email).Count(&count)
	if isExist.Error != nil {
		return nil, errors.New("usersRepository.RetrieveUserByEmail err = " + isExist.Error.Error())
	}

	if count == 0 {
		// return error string "2" for email not exist
		return nil, errors.New("2")
	}

	data := db.Conn.Where("email = ?", email).Find(&users)

	if data.Error != nil {
		return nil, errors.New("usersRepository.RetrieveUserByEmail err = " + data.Error.Error())
	}

	return &users, nil

}

// InsertUser - insert to table users
func (db *sqlRepository) InsertUser(data *usersModel.ReqInsertUser) (uint, error) {
	users := usersModel.Users{
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
		RoleID:   data.RoleID,
	}

	queryInsert := db.Conn.Create(&users)

	if queryInsert.Error != nil {
		return 0, errors.New("usersRepository.InsertUser err = " + queryInsert.Error.Error())
	}

	return users.ID, nil
}

// UpdateUser - update users by ID
func (db *sqlRepository) UpdateUser(data *usersModel.ReqUpdateUser) error {
	var users usersModel.Users

	// cek id is exist
	var count int
	isExist := db.Conn.Model(&usersModel.Users{}).Where("id = ?", data.ID).Count(&count)

	if count == 0 {
		// return error string "1" for ID not exist
		return errors.New("1")
	}

	if isExist.Error != nil {
		return errors.New("usersRepository.UpdateUser err = " + isExist.Error.Error())
	}

	queryUpdate := db.Conn.Model(&users).Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"email":    data.Email,
			"username": data.Username,
			"role_id":  data.RoleID,
		})

	if queryUpdate.Error != nil {
		return errors.New("usersRepository.UpdateUser err = " + queryUpdate.Error.Error())
	}

	return nil
}

// RetrieveUserByID -get users by id
func (db *sqlRepository) RetrieveUserByID(ID int) (*usersModel.ResGetUsers, error) {
	var users usersModel.Users

	// cek id is exist
	var count int
	isExist := db.Conn.Model(&usersModel.Users{}).Where("id = ?", ID).Count(&count)

	if count == 0 {
		// return error string "1" for ID not exist
		return nil, errors.New("1")
	}

	if isExist.Error != nil {
		return nil, errors.New("usersRepository.RetrieveUserByID err = " + isExist.Error.Error())
	}

	// get data users by id
	dataUsers := db.Conn.Where("id = ?", ID).Find(&users)
	if dataUsers.Error != nil {
		return nil, errors.New("usersRepository.RetrieveUserByID err = " + dataUsers.Error.Error())
	}
	dataResp := usersModel.ResGetUsers{
		ID:        users.ID,
		Email:     users.Email,
		Username:  users.Username,
		RoleID:    users.RoleID,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}

	return &dataResp, nil

}

// DeleteUser - delete user by ID
func (db *sqlRepository) DeleteUser(ID int) error {
	var users usersModel.Users

	// cek id is exist
	var count int
	isExist := db.Conn.Model(&usersModel.Users{}).Where("id = ?", ID).Count(&count)

	if count == 0 {
		// return error string "1" for ID not exist
		return errors.New("1")
	}

	if isExist.Error != nil {
		return errors.New("usersRepository.DeleteUser err = " + isExist.Error.Error())
	}

	// delete users
	queryDelete := db.Conn.Where("id = ?", ID).Delete(&users)

	if queryDelete.Error != nil {
		return errors.New("usersRepository.DeleteUser err = " + queryDelete.Error.Error())
	}

	return nil
}

// IsEmailExist - validation email is exist
func (db *sqlRepository) IsEmailExist(email string) (bool, error) {
	// cek email is exist
	var count int
	isExist := db.Conn.Model(&usersModel.Users{}).Where("email = ?", email).Count(&count)
	if isExist.Error != nil {
		return false, errors.New("usersRepository.IsEmailExist err = " + isExist.Error.Error())
	}

	if count == 0 {
		// return false for email not exist
		return false, nil
	}

	return true, nil
}

// RetrieveRoles - get roles json
func (db *sqlRepository) RetrieveRoles() (*[]usersModel.Roles, error) {
	// open json file
	jsonFile, err := os.Open("modules/users/usersDTO/role.json")
	if err != nil {
		return nil, errors.New("usersRepository.RetrieveRoles err = " + err.Error())
	}

	log.Info().Msg("Successfully Opened role.json")

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var role []usersModel.Roles

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'role' which we defined above
	json.Unmarshal(byteValue, &role)

	return &role, nil
}

// RetrieveRoleByID - get roles json by id
func (db *sqlRepository) RetrieveRoleByID(ID int) (*usersModel.Roles, error) {
	// open json file
	jsonFile, err := os.Open("modules/users/usersDTO/role.json")
	if err != nil {
		return nil, errors.New("usersRepository.RetrieveRoleByID err = " + err.Error())
	}

	log.Info().Msg("Successfully Opened role.json")

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var role []usersModel.Roles

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'role' which we defined above
	json.Unmarshal(byteValue, &role)
	var roleID int
	var result usersModel.Roles
	for _, list := range role {
		if ID == list.ID {
			result.ID = list.ID
			result.RoleName = list.RoleName
			roleID = list.ID
			break
		} else {
			roleID = 0
		}
	}

	if roleID == 0 {
		// return error string "1" for ID not exist
		return nil, errors.New("1")
	}

	return &result, nil
}

// UpdateCredentialByID - update credential by ID
func (db *sqlRepository) UpdateCredentialByID(ID int, password string) error {
	queryUpdateCredential := db.Conn.Model(&usersModel.Users{}).Where("id = ?", ID).
		Update(map[string]interface{}{
			"password": password,
		})

	if queryUpdateCredential.Error != nil {
		return errors.New("usersRepository.UpdateCredentialByID err = " + queryUpdateCredential.Error.Error())
	}

	return nil
}

// RetrieveCredentialByID - get credential by ID
func (db *sqlRepository) RetrieveCredentialByID(ID int) (string, error) {
	var users usersModel.Users
	credential := db.Conn.Select("password").Where("id = ?", ID).Find(&users)

	if credential.Error != nil {
		return "", errors.New("usersRepository.RetrieveCredentialByID err = " + credential.Error.Error())
	}

	return users.Password, nil
}
