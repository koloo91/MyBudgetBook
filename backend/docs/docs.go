// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-26 22:19:41.04258 +0100 CET m=+0.029888595

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/accounts": {
            "get": {
                "description": "Get all accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get all accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AccountsVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "Create account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AccountVo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.AccountVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            }
        },
        "/api/alive": {
            "get": {
                "description": "Checks if the user is logged in",
                "tags": [
                    "Ping"
                ],
                "summary": "Checks if the user is logged in",
                "responses": {
                    "204": {}
                }
            }
        },
        "/api/bookings": {
            "get": {
                "description": "Get bookings in the given time range",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Get bookings in the given time range",
                "parameters": [
                    {
                        "type": "string",
                        "description": "start date of the range: 2006-01-02T15:04:05Z07:00",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end date of the range: 2006-01-02T15:04:05Z07:00",
                        "name": "endDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BookingsVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new booking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Create a new booking",
                "parameters": [
                    {
                        "description": "Create booking",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.BookingVo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.BookingVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            }
        },
        "/api/bookings/{id}": {
            "put": {
                "description": "Updates a booking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Updates a booking",
                "parameters": [
                    {
                        "description": "Create booking",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.BookingVo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Booking id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "ONE",
                            "ALL"
                        ],
                        "type": "string",
                        "default": "ONE",
                        "description": "update only this entry or all of the standing order",
                        "name": "updateStrategy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BookingVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a booking",
                "tags": [
                    "Bookings"
                ],
                "summary": "Deletes a booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            }
        },
        "/api/categories": {
            "get": {
                "description": "Get current balances",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balances"
                ],
                "summary": "Get current balances",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AccountBalancesVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Create a new category",
                "parameters": [
                    {
                        "description": "Create category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CategoryVo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.CategoryVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            }
        },
        "/api/categories/{id}": {
            "put": {
                "description": "Updates a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Updates a category",
                "parameters": [
                    {
                        "description": "Update category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CategoryVo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CategoryVo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorVo"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "get": {
                "description": "Checks if the service is running",
                "tags": [
                    "Alive"
                ],
                "summary": "Checks if the service is running",
                "responses": {
                    "204": {}
                }
            }
        }
    },
    "definitions": {
        "model.AccountBalanceVo": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.AccountBalancesVo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AccountBalanceVo"
                    }
                }
            }
        },
        "model.AccountVo": {
            "type": "object",
            "required": [
                "name",
                "startingBalance"
            ],
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "startingBalance": {
                    "type": "number"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.AccountsVo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AccountVo"
                    }
                }
            }
        },
        "model.BookingVo": {
            "type": "object",
            "required": [
                "accountId",
                "amount",
                "categoryId",
                "date",
                "title"
            ],
            "properties": {
                "accountId": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "categoryId": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "standingOrderId": {
                    "type": "string"
                },
                "standingOrderLastDay": {
                    "type": "string"
                },
                "standingOrderPeriod": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.BookingsVo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BookingVo"
                    }
                }
            }
        },
        "model.CategoriesVo": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CategoryVo"
                    }
                }
            }
        },
        "model.CategoryVo": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.ErrorVo": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
