package usersModel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	// Users - object table users
	Users struct {
		gorm.Model
		Email    string `gorm:"column:email"`
		Username string `gorm:"column:username"`
		Password string `gorm:"column:password"`
		RoleID   int    `gorm:"column:role_id"`
	}

	// ResGetAllUsers - object response get all users
	ResGetUsers struct {
		ID        uint      `json:"user_id" example:"1"`
		Email     string    `json:"email" example:"bla@bla.com"`
		Username  string    `json:"username" example:"your username"`
		RoleID    int       `json:"role_id" example:"1"`
		CreatedAt time.Time `json:"created_at" example:"2021-03-13T07:15:11+07:00"`
		UpdatedAt time.Time `json:"updated_at" example:"2021-03-13T07:15:11+07:00"`
	}

	// ReqInsertUser - object request insert users
	ReqInsertUser struct {
		Email    string `json:"email" binding:"required" example:"bla@bla.com"`
		Username string `json:"username" binding:"required" example:"your username"`
		Password string `json:"password" binding:"required" example:"your password"`
		RoleID   int    `json:"role_id" binding:"required" example:"1"`
	}

	// ReqUpdateUser - object request update users
	ReqUpdateUser struct {
		ID       uint   `json:"user_id" binding:"required" example:"1"`
		Email    string `json:"email" binding:"required" example:"bla@bla.com"`
		Username string `json:"username" binding:"required" example:"your username"`
		RoleID   int    `json:"role_id" binding:"required" example:"1"`
	}

	// ReqLogin - object request login user
	ReqLogin struct {
		Email    string `json:"email" binding:"required" example:"bla@bla.com"`
		Password string `json:"password" binding:"required" example:"your password"`
	}

	// ResLogin - response login user
	ResLogin struct {
		ID     uint   `json:"user_id" example:"1"`
		RoleID int    `json:"role_id" example:"1"`
		Token  string `json:"token" example:"asvvasvdavvdhavbhdhabvhdas.sabfhbhasb.ajsfbhbashb"`
	}

	// ReqChangePassword - request change password
	ReqChangePassword struct {
		OldPassword string `json:"old_password" example:"123456789"`
		NewPassword string `json:"new_password" example:"987654321"`
	}
)

// role object manajemen
type (
	// Roles - object role data
	Roles struct {
		ID       int    `json:"role_id" example:"1"`
		RoleName string `json:"role_name" example:"Admin"`
	}
)
