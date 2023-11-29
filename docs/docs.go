// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "nandarusfikri@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/forgot-password": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for Request Forgot Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot Password",
                "operationId": "User-ForgotPassword",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ForgotPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "operationId": "User-Login",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/auth/reset-password": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for confirm reset password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset Password",
                "operationId": "User-ResetPassword",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ResetPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Product List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product List",
                "operationId": "Item-ProductList",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/user": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API untuk mengedit data user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Update",
                "operationId": "User-UserUpdate",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API untuk menambahkan user baru",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Insert",
                "operationId": "User-UserInsert",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserInsert"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/user/change-password": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for change password user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Change Password",
                "operationId": "User-ChangePassword",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API untuk mengambil data list user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User List",
                "operationId": "User-UserList",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UsersResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ChangePassword": {
            "type": "object",
            "required": [
                "new_password",
                "old_password",
                "user_id"
            ],
            "properties": {
                "new_password": {
                    "type": "string",
                    "example": "Password1!"
                },
                "old_password": {
                    "type": "string",
                    "example": "Password1!"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.ForgotPassword": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "nanda@gmail.com"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "super.admin@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "dto.ProductsRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "order_field": {
                    "type": "string",
                    "example": "id|desc"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "search_text": {
                    "type": "string",
                    "example": "Search name ku"
                }
            }
        },
        "dto.ResetPassword": {
            "type": "object",
            "required": [
                "email",
                "new_password",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.UserInsert": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "nandarusfikri@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Nanda Rusfikri"
                },
                "password": {
                    "type": "string",
                    "example": "Password1!"
                }
            }
        },
        "dto.UserUpdate": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "nandarusfikri@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_active": {
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "type": "string",
                    "example": "Nanda Rusfikri"
                },
                "phone": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "dto.UsersRequest": {
            "type": "object",
            "properties": {
                "is_active": {
                    "description": "Email      string ` + "`" + `json:\"email\"` + "`" + `\nPhone      string ` + "`" + `json:\"phone\"` + "`" + `",
                    "type": "boolean"
                },
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "order_field": {
                    "type": "string",
                    "example": "id|desc"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "search_text": {
                    "type": "string"
                }
            }
        },
        "dto.UsersResponse": {
            "type": "object",
            "properties": {
                "avatar_path": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header",
            "x-extension-openapi": "{\"example\": \"value on a json format\"}"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-template-project",
	Description:      "This Project Template Go which I usually use.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
