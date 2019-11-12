// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-12 19:56:05.655386 +0800 CST m=+0.047382527

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
            "name": "API Support",
            "url": "https://github.com/mritd/ginmvc",
            "email": "mritd@linux.com"
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
        "/api/v1/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"success\",\"timestamp\":1570889247}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    },
                    "400": {
                        "description": "{\"message\":\"user not registered or password is wrong\",\"timestamp\":1570889272}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    },
                    "500": {
                        "description": "{\"message\":\"invalid connection\",\"timestamp\":1570887977}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/user/logout": {
            "post": {
                "description": "user logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "user logout",
                "responses": {
                    "200": {
                        "description": "{\"message\":\"success\",\"timestamp\":1570889247}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    },
                    "500": {
                        "description": "{\"message\":\"error message\",\"timestamp\":1570887977}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "description": "user register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "user register",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\":\"success\",\"timestamp\":1570887849}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    },
                    "400": {
                        "description": "{\"message\":\"user email [mritd@linux.com] already register\",\"timestamp\":1570887884}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    },
                    "500": {
                        "description": "{\"message\":\"invalid connection\",\"timestamp\":1570887977}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "health check",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Monitor API"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "{\"message\":\"success\",\"timestamp\":1570889247}",
                        "schema": {
                            "$ref": "#/definitions/models.CommonResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CommonResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        },
        "models.NullBool": {
            "type": "object"
        },
        "models.NullInt32": {
            "type": "object"
        },
        "models.NullInt64": {
            "type": "object"
        },
        "models.NullString": {
            "type": "object"
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createTime": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullInt64"
                },
                "email": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullString"
                },
                "id": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullInt32"
                },
                "lock": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullBool"
                },
                "loginTime": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullInt64"
                },
                "mobile": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullString"
                },
                "name": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullString"
                },
                "password": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullString"
                },
                "salt": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullString"
                },
                "updateTime": {
                    "type": "object",
                    "$ref": "#/definitions/models.NullInt64"
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
	Host:        ":8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Gin MVC Template",
	Description: "Gin MVC Template.",
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
