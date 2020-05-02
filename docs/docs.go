// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-26 16:55:57.693304 -0300 -03 m=+10.321064253

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
        "contact": {
            "name": "Manuel Porto",
            "email": "manu.porto94@gmail.com"
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
        "/nodes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Obtiene todos los nodos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_models.NodeDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Crea un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api_models.NodeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Obtiene la información completa de un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_models.NodeDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "nodes"
                ],
                "summary": "Borra un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}/configuration": {
            "get": {
                "description": "Devuelve un mapa de estados (no un array como se ve a continuación) en donde la clave de cada uno es el nombre del mismo.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Obtiene la configuración de un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_models.StateDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Devuelve un mapa de estados (no un array como se ve a continuación) en donde la clave de cada uno es el nombre del mismo.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Crea o actualiza la configuración para el nodo dado.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_models.StateDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}/manual-reading": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Obtiene el estado de medición manual de un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_models.ManualReadingDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Actualiza el estado de medición manual de un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_models.ManualReadingDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}/readings": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readings"
                ],
                "summary": "Obtiene las mediciones de un nodo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_models.GetReadingDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readings"
                ],
                "summary": "Crea una medición",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos de la medición",
                        "name": "reading",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_models.ReadingDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_models.GetReadingDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}/readings/{reading_id}/photos": {
            "get": {
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "readings"
                ],
                "summary": "Obtiene la foto de una medición",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de la medición",
                        "name": "reading_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "image/jpeg"
                ],
                "tags": [
                    "readings"
                ],
                "summary": "Agrega una foto a la medición",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de la medición",
                        "name": "reading_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Foto de la medición",
                        "name": "picture",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/png"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/nodes/{node_id}/readings/{reading_id}s": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readings"
                ],
                "summary": "Obtiene los datos de una medición",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del nodo",
                        "name": "node_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID de la medición",
                        "name": "reading_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_models.GetReadingDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_models.GetReadingDTO": {
            "type": "object",
            "properties": {
                "manualReading": {
                    "type": "boolean",
                    "example": false
                },
                "readingId": {
                    "type": "string",
                    "example": "00336270-8191-11ea-a43d-0242ac120003"
                },
                "readingTime": {
                    "type": "string",
                    "example": "2020-04-26T19:47:53.391Z"
                },
                "waterLevel": {
                    "type": "number",
                    "example": 60
                }
            }
        },
        "api_models.ManualReadingDTO": {
            "type": "object",
            "properties": {
                "manualReading": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "api_models.NodeDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Nodo instalado debajo de un puente"
                },
                "id": {
                    "type": "string",
                    "example": "lujan-1"
                },
                "manual_reading": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "api_models.ReadingDTO": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string",
                    "example": "2020-04-26T19:47:53.391Z"
                },
                "waterLevel": {
                    "type": "number",
                    "example": 60
                }
            }
        },
        "api_models.StateDTO": {
            "type": "object",
            "properties": {
                "interval": {
                    "type": "integer",
                    "example": 25
                },
                "lowerLimit": {
                    "type": "number",
                    "example": 0.5
                },
                "next": {
                    "type": "string",
                    "example": "Alto"
                },
                "picturesNum": {
                    "type": "integer",
                    "example": 1
                },
                "prev": {
                    "type": "string",
                    "example": "Bajo"
                },
                "upperLimit": {
                    "type": "number",
                    "example": 1
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "internal": {
                    "description": "Stores the error returned by an external dependency",
                    "type": "error"
                },
                "message": {
                    "type": "object"
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
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Hydro Monitor Web API",
	Description: "This is the Hydro Monitor Web API",
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
