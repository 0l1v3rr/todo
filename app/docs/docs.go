// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://0l1v3rr.github.io",
            "email": "oliver.mrakovics@gmail.com"
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
        "/login": {
            "post": {
                "description": "Logs in a user and saves the JWT in cookies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User endpoints"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User to log in",
                        "name": "user",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If the login was successful.",
                        "schema": {
                            "$ref": "#/definitions/util.Success"
                        }
                    },
                    "400": {
                        "description": "If the provided user is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "403": {
                        "description": "If the password is incorrect or the user is not activated.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "404": {
                        "description": "If the user with the specified email does not exist.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Logs out the currently logged-in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User endpoints"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "If the logout was success.",
                        "schema": {
                            "$ref": "#/definitions/util.Success"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user into the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User endpoints"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "User to register",
                        "name": "user",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "If the user has been created successfully.",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "If the provided user is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "409": {
                        "description": "If the specified email already exists.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "If there was a server error while creating the user.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        },
        "/tasks/list/{id}": {
            "get": {
                "description": "Returns all the tasks in the specified list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Get list tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "list ID",
                        "name": "id",
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
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "If the id is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "If there was a db error..",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "Returns the task with the specified id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Get tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "If the id is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "404": {
                        "description": "If the task doesn't exist.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Edits the task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Edit task",
                "parameters": [
                    {
                        "description": "Task to create",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "If the task or the id is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "401": {
                        "description": "If the user is not logged in.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "403": {
                        "description": "If the user has no permission to do this.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "404": {
                        "description": "If the task does not exist.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "If there was a db error.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new task in the db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Create task",
                "parameters": [
                    {
                        "description": "Task to create",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "If the task is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "401": {
                        "description": "If the user is not logged in.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "If there was a db error.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the task",
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "If the id is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "401": {
                        "description": "If the user is not logged in.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "403": {
                        "description": "If the user has no permission to do this.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "404": {
                        "description": "If the task does not exist.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "If there was a db error.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            },
            "patch": {
                "description": "Changes the task status to its opposite value",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task endpoints"
                ],
                "summary": "Change task status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "If the id is not valid.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "401": {
                        "description": "If the user is not logged in.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "403": {
                        "description": "If the user has no permission to do this.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "404": {
                        "description": "If the task doesn't exist.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Returns the currently logged-in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User endpoints"
                ],
                "summary": "Logged In User",
                "responses": {
                    "200": {
                        "description": "If the user is logged in.",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "401": {
                        "description": "If the user is not logged in.",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LoginUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "SuperSecret69"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_done": {
                    "type": "boolean"
                },
                "list_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_enabled": {
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "util.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "An unknown error occurred."
                }
            }
        },
        "util.Success": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success!"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Advanced ToDo application",
	Description:      "This is the API of the ToDo application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}