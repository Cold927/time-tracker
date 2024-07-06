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
        "/tasks/countdown/end/{tid}": {
            "patch": {
                "description": "Закончить отсчет времени по задаче для пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Закончить отсчет времени по задаче для пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID задачи",
                        "name": "tid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/tasks/countdown/start/{uid}": {
            "post": {
                "description": "Начать отсчет времени по задаче для пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Начать отсчет времени по задаче для пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Описание задачи",
                        "name": "task",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.TaskCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/tasks/info/{uid}": {
            "get": {
                "description": "Получение трудозатрат по пользователю за период задача-сумма часов и минут",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Получение трудозатрат по пользователю",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-07-01T00:00:00",
                        "description": "Начальная дата",
                        "name": "startDate",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2024-07-01T23:59:59",
                        "description": "Конечная дата",
                        "name": "endDate",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TaskResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
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
                            "$ref": "#/definitions/model.UserCreate"
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
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
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
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.userResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/info": {
            "get": {
                "description": "Получение данных о пользователе по серии и номеру паспорта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Получение данных о пользователе по паспорту",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Поиск по серии паспорта",
                        "name": "passportSeries",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Поиск по номеру паспорта",
                        "name": "passportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.userResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "get": {
                "description": "Получение данных о всех пользователях",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Получение данных о всех пользователях",
                "parameters": [
                    {
                        "type": "string",
                        "default": "1",
                        "description": "Укажите с какой страницы смотреть",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "10",
                        "description": "Укажите какое количество выводить",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "asc, desc",
                        "description": "Сортировать данные",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Id, Surname, Name, Patronymic, Address, PassportSeries, PassportNumber",
                        "description": "Поле для сортировки",
                        "name": "field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Поиск по полям",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Pagination"
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
                            "$ref": "#/definitions/model.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPSuccess"
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
        "model.TaskCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Описание задачи"
                },
                "title": {
                    "type": "string",
                    "example": "Новая задача"
                }
            }
        },
        "model.TaskResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "period_time": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
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
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "passportNumber": {
                    "type": "integer"
                },
                "passportSeries": {
                    "type": "integer"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.UserCreate": {
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
                "passportNumber": {
                    "type": "string",
                    "example": "1234 567890"
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
        },
        "utils.HTTPSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "Status: OK"
                }
            }
        },
        "utils.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "rows": {},
                "sort": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
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
