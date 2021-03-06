// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Simple Go Api for AjoCard Project",
    "title": "AjoCard-http-go-server",
    "contact": {
      "email": "shekarau.buba@outlook.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    },
    "/hello/{user}": {
      "get": {
        "description": "Return Greeting to user",
        "parameters": [
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid character in 'user' were provided"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Simple Go Api for AjoCard Project",
    "title": "AjoCard-http-go-server",
    "contact": {
      "email": "shekarau.buba@outlook.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    },
    "/hello/{user}": {
      "get": {
        "description": "Return Greeting to user",
        "parameters": [
          {
            "type": "string",
            "name": "user",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the greeting",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid character in 'user' were provided"
          }
        }
      }
    }
  }
}`))
}
