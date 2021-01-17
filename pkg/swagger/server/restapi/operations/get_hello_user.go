// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHelloUserHandlerFunc turns a function with the right signature into a get hello user handler
type GetHelloUserHandlerFunc func(GetHelloUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHelloUserHandlerFunc) Handle(params GetHelloUserParams) middleware.Responder {
	return fn(params)
}

// GetHelloUserHandler interface for that can handle valid get hello user params
type GetHelloUserHandler interface {
	Handle(GetHelloUserParams) middleware.Responder
}

// NewGetHelloUser creates a new http.Handler for the get hello user operation
func NewGetHelloUser(ctx *middleware.Context, handler GetHelloUserHandler) *GetHelloUser {
	return &GetHelloUser{Context: ctx, Handler: handler}
}

/*GetHelloUser swagger:route GET /hello/{user} getHelloUser

Return Greeting to user

*/
type GetHelloUser struct {
	Context *middleware.Context
	Handler GetHelloUserHandler
}

func (o *GetHelloUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHelloUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}