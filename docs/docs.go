// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Muhammad nur basari",
            "email": "m.nurbasari@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/change-password": {
            "put": {
                "description": "Change Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "json object update password",
                        "name": "request_change_password",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usersModel.ReqChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResChangePassword"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/delete/{user_id}": {
            "delete": {
                "description": "delete user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "delete user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResDeleteUserByID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "get All Data Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "get All Data Users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResGetAllUsers"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/list/{user_id}": {
            "get": {
                "description": "get User by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "get User by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResGetUserByID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login to system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AUTH"
                ],
                "summary": "Login Users, please try using POSTMAN, Just ONLY THIS METHOD",
                "parameters": [
                    {
                        "description": "json object email pwd",
                        "name": "request_email_pwd",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ReqLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResLoginResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/role": {
            "get": {
                "description": "get roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "get roles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResGetRoles"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/role/{role_id}": {
            "get": {
                "description": "get roles by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "get roles by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResGetRoleByID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/save": {
            "post": {
                "description": "Create an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "Create an user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "json object insert user",
                        "name": "request_insert",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usersModel.ReqInsertUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResInsertUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "description": "update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TOKEN , please type Bearer before ApiKeyAuth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "userid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "json object update user",
                        "name": "request_update",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usersModel.ReqUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usersDTO.ResUpdateUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.BadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.UnAuthorized"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.ForbiddenRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/globalDTO.InternalServerErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "globalDTO.BadRequest": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "MISSING PARAMETERS"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "globalDTO.ForbiddenRes": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Invalid Token"
                },
                "status": {
                    "type": "integer",
                    "example": 403
                }
            }
        },
        "globalDTO.InternalServerErr": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Please Contact admin, Server is Having Error"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "globalDTO.UnAuthorized": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "unauthorized user"
                },
                "status": {
                    "type": "integer",
                    "example": 401
                }
            }
        },
        "usersDTO.ReqLoginDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "xxxxxx@xxxxxx.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456123"
                }
            }
        },
        "usersDTO.ResChangePassword": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Change Password Success"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResDeleteUserByID": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "User has been Deleted"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResGetAllUsers": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Success"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usersModel.ResGetUsers"
                    }
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResGetRoleByID": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Success"
                },
                "result": {
                    "$ref": "#/definitions/usersModel.Roles"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResGetRoles": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Success"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usersModel.Roles"
                    }
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResGetUserByID": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Success"
                },
                "result": {
                    "$ref": "#/definitions/usersModel.ResGetUsers"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResInsertUser": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "User has been Created"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "usersDTO.ResLoginResult": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "Login Success"
                },
                "result": {
                    "$ref": "#/definitions/usersModel.ResLogin"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersDTO.ResUpdateUser": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "string",
                    "example": "User has been Updated"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usersModel.ReqChangePassword": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string",
                    "example": "987654321"
                },
                "old_password": {
                    "type": "string",
                    "example": "123456789"
                }
            }
        },
        "usersModel.ReqInsertUser": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role_id",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "bla@bla.com"
                },
                "password": {
                    "type": "string",
                    "example": "your password"
                },
                "role_id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "your username"
                }
            }
        },
        "usersModel.ReqUpdateUser": {
            "type": "object",
            "required": [
                "email",
                "role_id",
                "user_id",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "bla@bla.com"
                },
                "role_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "your username"
                }
            }
        },
        "usersModel.ResGetUsers": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2021-03-13T07:15:11+07:00"
                },
                "email": {
                    "type": "string",
                    "example": "bla@bla.com"
                },
                "role_id": {
                    "type": "integer",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "2021-03-13T07:15:11+07:00"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "your username"
                }
            }
        },
        "usersModel.ResLogin": {
            "type": "object",
            "properties": {
                "role_id": {
                    "type": "integer",
                    "example": 1
                },
                "token": {
                    "type": "string",
                    "example": "asvvasvdavvdhavbhdhabvhdas.sabfhbhasb.ajsfbhbashb"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "usersModel.Roles": {
            "type": "object",
            "properties": {
                "role_id": {
                    "type": "integer",
                    "example": 1
                },
                "role_name": {
                    "type": "string",
                    "example": "Admin"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "localhost:12345",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Manajemen Komponen API",
	Description: "This is a Manajemen Komponen server API Documentation.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
