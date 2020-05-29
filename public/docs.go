// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package public

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Sai Kris",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/answer-video/{candidateId}/{sessionId}/{questionId}": {
            "put": {
                "description": "SaveAnswerVideo.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SaveAnswerVideo.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "this is a test file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Candidate id",
                        "name": "candidateId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Session id",
                        "name": "sessionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Session"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/interview-completion/{candidateId}/{sessionId}": {
            "post": {
                "description": "called when the interview is completed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "InterviewCompleted.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate id",
                        "name": "candidateId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Session id",
                        "name": "sessionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/question": {
            "put": {
                "description": "SaveQuestionMetadata saves question metadata.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SaveQuestionMetadata saves question metadata.",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.QuestionMetadata"
                        }
                    }
                ],
                "responses": {
                    "202": {},
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/question/{questionId}": {
            "get": {
                "description": "Get Question video by question id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "video/mp4"
                ],
                "summary": "Get Question video by question id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/questions": {
            "get": {
                "description": "Get Questions by session id.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Questions by session id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session id",
                        "name": "sessionId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Questions"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/session": {
            "put": {
                "description": "creates an interview session.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "creates an interview session.",
                "parameters": [
                    {
                        "description": "creates an interview session",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Session"
                        }
                    }
                ],
                "responses": {
                    "202": {},
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/session/{sessionId}": {
            "get": {
                "description": "Get Session by session id.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Session by session id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session id",
                        "name": "sessionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Session"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {},
                    "404": {},
                    "500": {}
                }
            }
        }
    },
    "definitions": {
        "models.Question": {
            "type": "object",
            "properties": {
                "answerTimeInSeconds": {
                    "type": "integer"
                },
                "questionId": {
                    "type": "string"
                },
                "questionName": {
                    "type": "string"
                },
                "questionText": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "models.QuestionMetadata": {
            "type": "object",
            "properties": {
                "answerText": {
                    "type": "string"
                },
                "answerTimeInSeconds": {
                    "type": "integer"
                },
                "importantKeywords": {
                    "type": "string"
                },
                "questionId": {
                    "type": "string"
                },
                "questionName": {
                    "type": "string"
                },
                "questionText": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "models.Questions": {
            "type": "object",
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Question"
                    }
                }
            }
        },
        "models.Session": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "sessionId": {
                    "type": "string"
                },
                "totalTimeInSeconds": {
                    "type": "integer"
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
	Host:        "localhost:8083",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Interviews API written in Golang",
	Description: "Interviews API  Golang to simplify the interview process.",
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
