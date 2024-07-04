// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks/countdown/end": {
            "patch": {
                "description": "Закончить отсчет времени по задаче для пользователя",
                "tags": [
                    "tasks"
                ],
                "summary": "Закончить отсчет времени по задаче для пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID задачи",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/countdown/start": {
            "post": {
                "description": "Начать отсчет времени по задаче для пользователя",
                "tags": [
                    "tasks"
                ],
                "summary": "Начать отсчет времени по задаче для пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/create": {
            "post": {
                "description": "Создает нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Создает нового пользователя",
                "parameters": [
                    {
                        "description": "Новый пользователь",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.userResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/delete/{id}": {
            "delete": {
                "description": "Изменение данных пользователя",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Удаление пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/find/{id}": {
            "get": {
                "description": "Получение данных о пользователе по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Получение данных о пользователе по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/info": {
            "get": {
                "description": "Получение данных о всех пользователях",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Получение данных о всех пользователях",
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/update/{id}": {
            "patch": {
                "description": "Изменение данных пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Изменение данных пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Изменение данных пользователя",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.userResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "г. Москва, ул. Ленина, д. 5, кв. 1"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Status: Bad Request"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
