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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthLoginResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AuthRegisterResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/juzs": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get juzs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "juzs"
                ],
                "summary": "Get juzs",
                "parameters": [
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "modified_at",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "modified_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "value",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JuzGetResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create juzs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "juzs"
                ],
                "summary": "Create juzs",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.JuzCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JuzCreateResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/juzs/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete juzs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "juzs"
                ],
                "summary": "Delete juzs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JuzDeleteResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update juzs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "juzs"
                ],
                "summary": "Update juzs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.JuzUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JuzUpdateResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        },
        "/juzs/{id}/{child}/{child_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get juzs by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "juzs"
                ],
                "summary": "Get juzs by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id path",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JuzGetByIDResponseDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "abstraction.PaginationInfo": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "more_records": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "sort_by": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
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
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.AuthLoginResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthLoginResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.AuthRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponse": {
            "type": "object",
            "required": [
                "email",
                "is_active",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
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
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.AuthRegisterResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.AuthRegisterResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.JuzCreateRequest": {
            "type": "object",
            "required": [
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                }
            }
        },
        "dto.JuzCreateResponse": {
            "type": "object",
            "required": [
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                }
            }
        },
        "dto.JuzCreateResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.JuzCreateResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.JuzDeleteResponse": {
            "type": "object",
            "required": [
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "surahs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SurahEntityModel"
                    }
                }
            }
        },
        "dto.JuzDeleteResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.JuzDeleteResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.JuzGetByIDResponse": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "surahs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SurahEntityModel"
                    }
                }
            }
        },
        "dto.JuzGetByIDResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.JuzGetByIDResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.JuzGetResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.JuzEntityModel"
                            }
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "dto.JuzUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                }
            }
        },
        "dto.JuzUpdateResponse": {
            "type": "object",
            "required": [
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                },
                "surahs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SurahEntityModel"
                    }
                }
            }
        },
        "dto.JuzUpdateResponseDoc": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "properties": {
                        "data": {
                            "$ref": "#/definitions/dto.JuzUpdateResponse"
                        },
                        "meta": {
                            "$ref": "#/definitions/response.Meta"
                        }
                    }
                }
            }
        },
        "model.SurahEntityModel": {
            "type": "object",
            "required": [
                "nama_surah",
                "no_surah",
                "juz_id",
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_surah": {
                    "type": "string"
                },
                "no_surah": {
                    "type": "string"
                },
                "juz_id": {
                    "description": "relations",
                    "type": "integer"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                }
            }
        },
        "model.JuzEntityModel": {
            "type": "object",
            "required": [
                "nama_juz",
                "no_juz"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "modified_at": {
                    "type": "string"
                },
                "modified_by": {
                    "type": "string"
                },
                "nama_juz": {
                    "type": "string"
                },
                "no_juz": {
                    "type": "string"
                },
                "surahs": {
                    "description": "relations",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SurahEntityModel"
                    }
                }
            }
        },
        "response.Meta": {
            "type": "object",
            "properties": {
                "info": {
                    "$ref": "#/definitions/abstraction.PaginationInfo"
                },
                "message": {
                    "type": "string",
                    "default": "true"
                },
                "success": {
                    "type": "boolean",
                    "default": true
                }
            }
        },
        "response.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/response.Meta"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Host:        "localhost:3030",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "quran",
	Description: "This is a doc for quran.",
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
