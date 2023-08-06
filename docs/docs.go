// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/channels": {
            "get": {
                "description": "Retrieves all channels based on provided IDs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Channels",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Comma-separated list of channel IDs",
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Channel"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Updates channels by fetching from Youtube using provided Channel IDs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Channels from Youtube",
                "parameters": [
                    {
                        "description": "Array of Channel IDs",
                        "name": "channelIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Channels updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates channels by fetching from Youtube using provided Channel IDs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Channels from Youtube",
                "parameters": [
                    {
                        "description": "Array of Channel IDs",
                        "name": "channelIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Channels created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Retrieve all songs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all songs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Song"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update songs based on provided cronType",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update songs from Youtube",
                "parameters": [
                    {
                        "description": "Type of the cron",
                        "name": "cronType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Songs updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Updates songs by fetching from Youtube using provided Video IDs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Song from Youtube",
                "parameters": [
                    {
                        "description": "Array of Video IDs",
                        "name": "videoIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Songs updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Channel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "snippet": {
                    "$ref": "#/definitions/entities.ChannelSnippet"
                },
                "statistics": {
                    "$ref": "#/definitions/entities.ChannelStatistics"
                }
            }
        },
        "entities.ChannelSnippet": {
            "type": "object",
            "properties": {
                "customUrl": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "thumbnails": {
                    "$ref": "#/definitions/entities.Thumbnails"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entities.ChannelStatistics": {
            "type": "object",
            "properties": {
                "hiddenSubscriberCount": {
                    "type": "boolean"
                },
                "subscriberCount": {
                    "type": "string"
                },
                "videoCount": {
                    "type": "string"
                },
                "viewCount": {
                    "type": "string"
                }
            }
        },
        "entities.Song": {
            "type": "object",
            "properties": {
                "channelId": {
                    "type": "string"
                },
                "channelTitle": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "thumbnails": {
                    "$ref": "#/definitions/entities.Thumbnails"
                },
                "title": {
                    "type": "string"
                },
                "viewCount": {
                    "$ref": "#/definitions/entities.Views"
                }
            }
        },
        "entities.Thumbnail": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "entities.Thumbnails": {
            "type": "object",
            "properties": {
                "default": {
                    "$ref": "#/definitions/entities.Thumbnail"
                },
                "high": {
                    "$ref": "#/definitions/entities.Thumbnail"
                },
                "maxres": {
                    "$ref": "#/definitions/entities.Thumbnail"
                },
                "medium": {
                    "$ref": "#/definitions/entities.Thumbnail"
                },
                "standard": {
                    "$ref": "#/definitions/entities.Thumbnail"
                }
            }
        },
        "entities.Views": {
            "type": "object",
            "properties": {
                "daily": {
                    "type": "string"
                },
                "monthly": {
                    "type": "string"
                },
                "total": {
                    "type": "string"
                },
                "weekly": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "VSPO Common API",
	Description:      "This is the API documentation for VSPO Common services.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
