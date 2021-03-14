package usersDelivery

import (
	"manajemen-komponen-api/helpers"
	"manajemen-komponen-api/middlewares/authMiddleware"
	"manajemen-komponen-api/models/usersModel"
	"manajemen-komponen-api/modules/users"
	"manajemen-komponen-api/services/auditTrails"
	"manajemen-komponen-api/services/logErrors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type userHandler struct {
	usersUsecase       users.UsersUsecase
	auditTrailsUsecase auditTrails.AuditTrailsUsecase
	logErrorsUC        logErrors.LogErrorsUsecase
}

func NewUserHTTPHandler(r *gin.Engine, userUC users.UsersUsecase, auditTrailsUC auditTrails.AuditTrailsUsecase, logErrorsUC logErrors.LogErrorsUsecase) {
	handler := userHandler{
		usersUsecase:       userUC,
		auditTrailsUsecase: auditTrailsUC,
		logErrorsUC:        logErrorsUC,
	}

	authorized := r.Group("/user")
	authorized.Use(auditTrailsUC.MiddlewareAuditTrail)
	{
		authorized.POST("/login", authMiddleware.BasicAuth, handler.LoginUser)
		authorized.POST("/save", authMiddleware.JwtAuthWithHeader, handler.SaveUser)
		authorized.GET("/list", authMiddleware.JwtAuthWithHeader, handler.GetAllUsers)
		authorized.GET("/list/:user_id", authMiddleware.JwtAuthWithHeader, handler.GetUserByID)
		authorized.GET("/role", authMiddleware.JwtAuthWithHeader, handler.GetRoles)
		authorized.GET("/role/:role_id", authMiddleware.JwtAuthWithHeader, handler.GetRoleByID)
		authorized.DELETE("/delete/:user_id", authMiddleware.JwtAuthWithHeader, handler.DeleteUserByID)
		authorized.PUT("/update", authMiddleware.JwtAuthWithHeader, handler.EditUser)
		authorized.PUT("/change-password", authMiddleware.JwtAuthWithHeader, handler.ChangePassword)
	}
}

// LoginUser godoc
// @Summary Login Users, please try using POSTMAN, Just ONLY THIS METHOD
// @Description Login to system
// @Tags AUTH
// @Accept  json
// @Produce  json
// @Param request_email_pwd body usersDTO.ReqLoginDTO false "json object email pwd"
// @Success 200 {object} usersDTO.ResLoginResult
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/login [post]
func (handler *userHandler) LoginUser(c *gin.Context) {
	var req usersModel.ReqLogin
	moduleName := c.HandlerName()
	userId := "0"
	errBind := c.BindJSON(&req)

	if errBind != nil {
		log.Error().Msg(errBind.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)
		return
	}

	res, err := handler.usersUsecase.LoginUser(&req)

	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("Invalid Format Email")
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"status":   http.StatusForbidden,
					"messages": "Invalid Format Email",
				},
			)
			return
		}

		if err.Error() == "2" {
			log.Error().Msg("Email is Not Exist")
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"status":   http.StatusForbidden,
					"messages": "Email is Not Exist",
				},
			)
			return
		}

		if err.Error() == "3" {
			log.Error().Msg("Incorect Password")
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"status":   http.StatusForbidden,
					"messages": "Incorect Password",
				},
			)
			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userId, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}
	log.Info().Msg("Login Success")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "Login Success",
			"result":   &res,
		},
	)
}

// SaveUser godoc
// @Summary Create an user
// @Description Create an user
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param request_insert body usersModel.ReqInsertUser false "json object insert user"
// @Success 200 {object} usersDTO.ResInsertUser
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/save [post]
func (handler *userHandler) SaveUser(c *gin.Context) {
	moduleName := c.HandlerName()
	userId := c.Request.Header.Get("userid")

	var req usersModel.ReqInsertUser
	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(errBind.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":    http.StatusBadRequest,
				"merssages": helpers.MsgBadReq,
			},
		)
		return
	}

	ID, err := handler.usersUsecase.CreateUser(&req)
	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("Invalid Format Email")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":    http.StatusBadRequest,
					"merssages": "Invalid Format Email",
				},
			)
			return
		}

		if err.Error() == "2" {
			log.Error().Msg("Password min 5 char and max 15 char")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":    http.StatusBadRequest,
					"merssages": "Password min 5 char and max 15 char",
				},
			)
			return
		}

		if err.Error() == "3" {
			log.Error().Msg("Duplicate, Email is Exist")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":    http.StatusBadRequest,
					"merssages": "Duplicate, Email is Exist",
				},
			)
			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userId, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	log.Info().Msg("User has been Created")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "User has been Created",
			"user_id":  ID,
		},
	)

	return
}

// GetAllUsers godoc
// @Summary get All Data Users
// @Description get All Data Users
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Success 200 {object} usersDTO.ResGetAllUsers
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/list [get]
func (handler *userHandler) GetAllUsers(c *gin.Context) {
	moduleName := c.HandlerName()
	userId := c.Request.Header.Get("userid")

	data, err := handler.usersUsecase.GetAllUsers()
	if err != nil {
		isError := handler.logErrorsUC.CheckIsError(err, userId, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	if len(*data) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":   http.StatusOK,
				"messages": "success",
				"result":   []string{},
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "success",
			"result":   &data,
		},
	)

	return
}

// GetUserByID godoc
// @Summary get User by id
// @Description get User by id
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param user_id path string true "user_id"
// @Success 200 {object} usersDTO.ResGetUserByID
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/list/{user_id} [get]
func (handler *userHandler) GetUserByID(c *gin.Context) {
	moduleName := c.HandlerName()
	userIdHeader := c.Request.Header.Get("userid")

	userIdParam := c.Param("user_id")

	if userIdParam == "" {
		log.Error().Msg("user id in request url cant empty")
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)

		return
	}

	myID, errConv := strconv.Atoi(userIdParam)
	if errConv != nil {
		log.Error().Msg(errConv.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": errConv.Error(),
			},
		)

		return
	}

	data, err := handler.usersUsecase.GetUserByID(myID)
	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("user id is not exist")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":   http.StatusBadRequest,
					"messages": "user id is not exist",
				},
			)

			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userIdHeader, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "success",
			"result":   &data,
		},
	)

	return

}

// EditUser godoc
// @Summary update user
// @Description update user
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param request_update body usersModel.ReqUpdateUser false "json object update user"
// @Success 200 {object} usersDTO.ResUpdateUser
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/update [put]
func (handler *userHandler) EditUser(c *gin.Context) {
	moduleName := c.HandlerName()
	userIdHeader := c.Request.Header.Get("userid")

	var req usersModel.ReqUpdateUser
	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(helpers.MsgBadReq)
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)
		return
	}

	err := handler.usersUsecase.EditUser(&req)
	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("User ID is not Exist")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":   http.StatusBadRequest,
					"messages": "User ID is not Exist",
				},
			)
			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userIdHeader, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	log.Info().Msg("User has been Updated")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "User has been Updated",
		},
	)

	return
}

// DeleteUserByID godoc
// @Summary delete user by ID
// @Description delete user by ID
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param user_id path string true "user_id"
// @Success 200 {object} usersDTO.ResDeleteUserByID
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/delete/{user_id} [delete]
func (handler *userHandler) DeleteUserByID(c *gin.Context) {
	moduleName := c.HandlerName()
	userIdHeader := c.Request.Header.Get("userid")

	userIdParam := c.Param("user_id")

	if userIdParam == "" {
		log.Error().Msg("user id in request url cant empty")
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)

		return
	}

	myID, errConv := strconv.Atoi(userIdParam)
	if errConv != nil {
		log.Error().Msg(errConv.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": errConv.Error(),
			},
		)

		return
	}

	err := handler.usersUsecase.DropUser(myID)
	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("user id is not exist")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":   http.StatusBadRequest,
					"messages": "user id is not exist",
				},
			)

			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userIdHeader, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	log.Info().Msg("User has been Deleted")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "User has been Deleted",
		},
	)

	return

}

// GetRoles godoc
// @Summary get roles
// @Description get roles
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Success 200 {object} usersDTO.ResGetRoles
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/role [get]
func (handler *userHandler) GetRoles(c *gin.Context) {
	moduleName := c.HandlerName()
	userIdHeader := c.Request.Header.Get("userid")

	data, err := handler.usersUsecase.GetRoles()

	if err != nil {
		isError := handler.logErrorsUC.CheckIsError(err, userIdHeader, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	log.Info().Msg("Get Roles has been success")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "success",
			"result":   &data,
		},
	)

	return
}

// GetRoleByID godoc
// @Summary get roles by id
// @Description get roles by id
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param role_id path string true "role_id"
// @Success 200 {object} usersDTO.ResGetRoleByID
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/role/{role_id} [get]
func (handler *userHandler) GetRoleByID(c *gin.Context) {
	moduleName := c.HandlerName()
	userIdHeader := c.Request.Header.Get("userid")

	roleID := c.Param("role_id")

	if roleID == "" {
		log.Error().Msg("role id in request url cant empty")
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)

		return
	}

	myID, errConv := strconv.Atoi(roleID)
	if errConv != nil {
		log.Error().Msg(errConv.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": errConv.Error(),
			},
		)

		return
	}

	data, err := handler.usersUsecase.GetRoleByID(myID)

	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("role id is not exist")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":   http.StatusBadRequest,
					"messages": "role id is not exist",
				},
			)

			return
		}

		isError := handler.logErrorsUC.CheckIsError(err, userIdHeader, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return
	}

	log.Info().Msg("Get Roles By ID has been success")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "success",
			"result":   &data,
		},
	)

	return
}

// ChangePassword godoc
// @Summary Change Password
// @Description Change Password
// @Tags USER
// @Accept  json
// @Produce  json
// @param Authorization header string true "TOKEN , please type Bearer before ApiKeyAuth"
// @Param userid header string true "userid"
// @Param request_change_password body usersModel.ReqChangePassword false "json object update password"
// @Success 200 {object} usersDTO.ResChangePassword
// @Failure 400 {object} globalDTO.BadRequest
// @Failure 500 {object} globalDTO.InternalServerErr
// @Failure 403 {object} globalDTO.ForbiddenRes
// @Failure 401 {object} globalDTO.UnAuthorized
// @Router /user/change-password [put]
func (handler *userHandler) ChangePassword(c *gin.Context) {
	var req usersModel.ReqChangePassword

	errBind := c.BindJSON(&req)
	if errBind != nil {
		log.Error().Msg(helpers.MsgBadReq)
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": helpers.MsgBadReq,
			},
		)
		return
	}

	userId := c.Request.Header.Get("userid")
	myID, errConv := strconv.Atoi(userId)
	if errConv != nil {
		log.Error().Msg(errConv.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": errConv.Error(),
			},
		)
		return
	}

	err := handler.usersUsecase.UpdateCredential(myID, &req)
	if err != nil {
		if err.Error() == "1" {
			log.Error().Msg("Old Password is Not Matches")
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":   http.StatusBadRequest,
					"messages": "Old Password is Not Matches",
				},
			)
			return
		}

		moduleName := c.HandlerName()
		isError := handler.logErrorsUC.CheckIsError(err, userId, moduleName)
		log.Error().Msg(err.Error())
		if isError {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":   http.StatusInternalServerError,
					"messages": helpers.MsgErr,
				},
			)

			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": helpers.MsgErr,
			},
		)

		return

	}

	log.Info().Msg("Change Password Success")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "Change Password Success",
		},
	)

	return
}
