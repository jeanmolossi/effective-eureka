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
        "termsOfService": "github.com/jeanmolossi/effective-eureka/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/course": {
            "post": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Create a course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Course creation",
                "parameters": [
                    {
                        "description": "Course object which will be created",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateCourse"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCreateCourseBadRequestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            }
        },
        "/course/{courseID}": {
            "get": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Get a course by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Course retrieval",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "courseID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseOk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseByIDBadRequestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseNotFoundErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Edit a course basic information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Course edition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "courseID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Course object which will be edited",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.EditCourseInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseOk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpEditCourseInfoBadRequestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCourseNotFoundErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            }
        },
        "/course/{courseID}/module": {
            "post": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Create a module",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Module creation",
                "parameters": [
                    {
                        "description": "Module object which will be created",
                        "name": "module",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateModule"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "courseID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpModuleCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpBadRequestErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpMissingAuthenticationErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            }
        },
        "/module/{moduleID}": {
            "get": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Get a module",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Module retrieval",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Module ID",
                        "name": "moduleID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpModuleOk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpBadRequestErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpMissingAuthenticationErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpNotFoundErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Edit a module",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modules"
                ],
                "summary": "Module retrieval",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Module ID",
                        "name": "moduleID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Module object which will be updated",
                        "name": "module",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.EditModuleInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpModuleOk"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpBadRequestErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpMissingAuthenticationErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpNotFoundErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HttpInternalServerErr"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "A simple health check.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Ping the server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingOk"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/httputil.PingInternalServerErr"
                        }
                    }
                }
            }
        },
        "/students/me": {
            "get": {
                "security": [
                    {
                        "access_token": []
                    }
                ],
                "description": "Get auth student.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Get auth student.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpStudentRegistered"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCreateStudentBadRequestErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpStudentForbiddenErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpStudentInternalServerErr"
                        }
                    }
                }
            }
        },
        "/students/register": {
            "post": {
                "description": "Register a new student.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Register a new student.",
                "parameters": [
                    {
                        "description": "Student information",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.StudentInfo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpStudentRegistered"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpCreateStudentBadRequestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HttpStudentInternalServerErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginCredentials": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 6,
                    "example": "123456789"
                },
                "username": {
                    "type": "string",
                    "example": "jean@email.com"
                }
            }
        },
        "handler.HttpBadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.FieldError"
                    }
                }
            }
        },
        "handler.HttpCourseByIDBadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Missing course_id param"
                }
            }
        },
        "handler.HttpCourseCreated": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "05d4d9d3-01a3-4fd3-8d3e-e3178522f514"
                },
                "course_published": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "handler.HttpCourseNotFoundErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Course Not Found"
                }
            }
        },
        "handler.HttpCourseOk": {
            "type": "object",
            "properties": {
                "course_description": {
                    "type": "string",
                    "example": "Effective Eureka is a course about Go."
                },
                "course_id": {
                    "type": "string",
                    "example": "05d4d9d3-01a3-4fd3-8d3e-e3178522f514"
                },
                "course_published": {
                    "type": "boolean",
                    "example": false
                },
                "course_thumbnail": {
                    "type": "string",
                    "example": "https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"
                },
                "course_title": {
                    "type": "string",
                    "example": "Effective Eureka"
                }
            }
        },
        "handler.HttpCreateCourseBadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.FieldError"
                    }
                }
            }
        },
        "handler.HttpCreateStudentBadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.FieldError"
                    }
                }
            }
        },
        "handler.HttpEditCourseInfoBadRequestErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.FieldError"
                    }
                }
            }
        },
        "handler.HttpModuleCreated": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "05d4d9d3-01a3-4fd3-8d3e-e3178522f515"
                },
                "module_description": {
                    "type": "string",
                    "example": "Effective Eureka is a course about Go."
                },
                "module_id": {
                    "type": "string",
                    "example": "4aa77560-9c90-4128-b308-ad5c0515b5d7"
                },
                "module_published": {
                    "type": "boolean",
                    "example": false
                },
                "module_thumbnail": {
                    "type": "string",
                    "example": "https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"
                },
                "module_title": {
                    "type": "string",
                    "example": "Effective Eureka"
                }
            }
        },
        "handler.HttpModuleOk": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "05d4d9d3-01a3-4fd3-8d3e-e3178522f515"
                },
                "module_description": {
                    "type": "string",
                    "example": "Effective Eureka is a course about Go."
                },
                "module_id": {
                    "type": "string",
                    "example": "4aa77560-9c90-4128-b308-ad5c0515b5d7"
                },
                "module_published": {
                    "type": "boolean",
                    "example": false
                },
                "module_thumbnail": {
                    "type": "string",
                    "example": "https://effective-eureka.s3.amazonaws.com/courses/effective-eureka/thumbnail.png"
                },
                "module_title": {
                    "type": "string",
                    "example": "Effective Eureka"
                }
            }
        },
        "handler.HttpStudentForbiddenErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Missing authentication"
                }
            }
        },
        "handler.HttpStudentInternalServerErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "handler.HttpStudentRegistered": {
            "type": "object",
            "properties": {
                "student_email": {
                    "type": "string",
                    "example": "john@doe.com"
                },
                "student_id": {
                    "type": "string",
                    "example": "05d4d9d3-01a3-4fd3-8d3e-e3178522f514"
                }
            }
        },
        "httputil.HttpInternalServerErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "httputil.HttpMissingAuthenticationErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "missing authentication"
                }
            }
        },
        "httputil.HttpNotFoundErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "not found"
                }
            }
        },
        "httputil.PingInternalServerErr": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "unexpected error"
                }
            }
        },
        "httputil.PingOk": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "input.CreateCourse": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "This is a catalog video manager API."
                },
                "published": {
                    "type": "boolean",
                    "example": true
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://effective-eureka.s3.amazonaws.com/courses/thumbnail/1.jpg"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka"
                }
            }
        },
        "input.CreateModule": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "This is a catalog video manager API."
                },
                "published": {
                    "type": "boolean",
                    "example": true
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://effective-eureka.s3.amazonaws.com/courses/thumbnail/1.jpg"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka"
                }
            }
        },
        "input.EditCourseInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka is a course about effective eureka."
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://example.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka"
                }
            }
        },
        "input.EditModuleInfo": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "f0f8e8c4-8b8f-4d8e-b8e7-8f9e939ca9e8"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka is a course about effective eureka."
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://example.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Effective Eureka"
                }
            }
        },
        "input.StudentInfo": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8,
                    "example": "123456789"
                },
                "username": {
                    "type": "string",
                    "example": "john@doe.com"
                }
            }
        },
        "shared.FieldError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "field_name"
                },
                "message": {
                    "type": "string",
                    "example": "field_name is required"
                }
            }
        }
    },
    "securityDefinitions": {
        "access_token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
