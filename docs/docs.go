// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nguyen Tri",
            "email": "tringuyen2762001@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/cusotmers": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "route get customer for admin",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "get customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "post": {
                "description": "send email to user for reset password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "option when user forgot password",
                "parameters": [
                    {
                        "description": "user email",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ForgotPassword.userEmailReq"
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
        "/auth/login": {
            "post": {
                "description": "Use for login response the access_Token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserToken"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update Password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Updated Password",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdatePassword.userPassword"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/category/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "get category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "default: limit = 10",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "default: offset = 1",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BasePaginationResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create new category require admin permission",
                "parameters": [
                    {
                        "description": "New category",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/category/{id}/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "delete category by id require admin permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id category update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/category/{id}/update": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "update category by id require admin permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id category update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "category",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/media/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Create new image return image",
                "operationId": "image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "upfile",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/product/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create new product require admin permission",
                "parameters": [
                    {
                        "description": "New product",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "delete product by id require admin permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}/update": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "update product require admin permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id product update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "product",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "route get user Id from token then get user information",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user information by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "update user information",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "Updated UserInformation",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ForgotPassword.userEmailReq": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 150
                }
            }
        },
        "controller.UpdatePassword.userPassword": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "models.BasePaginationResponse": {
            "type": "object",
            "properties": {
                "items": {},
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "models.CreateCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "models.CreateProduct": {
            "type": "object",
            "required": [
                "category_id",
                "detail",
                "difficulty",
                "images",
                "light",
                "name",
                "size",
                "type",
                "warter"
            ],
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "light": {
                    "type": "string",
                    "maxLength": 50
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "string",
                    "maxLength": 10
                },
                "type": {
                    "type": "string",
                    "maxLength": 20
                },
                "warter": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "models.CreateUser": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "identifier",
                "lastName",
                "password",
                "phoneNumber"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 150
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 50
                },
                "identifier": {
                    "type": "string",
                    "maxLength": 100
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.UpdateCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "models.UpdateProduct": {
            "type": "object",
            "required": [
                "detail",
                "difficulty",
                "images",
                "is_publish",
                "light",
                "size",
                "type",
                "warter"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_publish": {
                    "type": "boolean"
                },
                "light": {
                    "type": "string",
                    "maxLength": 50
                },
                "size": {
                    "type": "string",
                    "maxLength": 10
                },
                "type": {
                    "type": "string",
                    "maxLength": 20
                },
                "warter": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "models.UpdateUser": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "phoneNumber"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 150
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 50
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "identifier": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.UserLogin": {
            "type": "object",
            "required": [
                "identifier",
                "password"
            ],
            "properties": {
                "identifier": {
                    "type": "string",
                    "maxLength": 100
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "models.UserToken": {
            "type": "object",
            "properties": {
                "access_Token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Fiber Go API",
	Description:      "greendeco",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
