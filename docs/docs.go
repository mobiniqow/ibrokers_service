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
        "/broker/api/v1/": {
            "get": {
                "description": "Get all brokers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "broker"
                ],
                "summary": "List of brokers",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/broker.Response"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new broker with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "broker"
                ],
                "summary": "Create broker",
                "parameters": [
                    {
                        "description": "Broker information",
                        "name": "broker",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/broker.CreateBrokerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/broker.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    }
                }
            }
        },
        "/broker/api/v1/{id}": {
            "get": {
                "description": "Retrieve details of a broker by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "broker"
                ],
                "summary": "Get broker details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Broker ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/broker.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid UUID format",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "404": {
                        "description": "Broker not found",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update broker details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "broker"
                ],
                "summary": "Update broker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Broker ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Broker information",
                        "name": "broker",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/broker.CreateBrokerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/broker.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "404": {
                        "description": "Broker not found",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a broker by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "broker"
                ],
                "summary": "Delete broker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Broker ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid UUID format",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "404": {
                        "description": "Broker not found",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/basics.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "basics.APIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "broker.CreateBrokerRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "nationalId": {
                    "type": "string"
                },
                "persianName": {
                    "type": "string"
                },
                "spotId": {
                    "type": "integer"
                }
            }
        },
        "broker.Response": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nationalId": {
                    "type": "string"
                },
                "persianName": {
                    "type": "string"
                },
                "spotId": {
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