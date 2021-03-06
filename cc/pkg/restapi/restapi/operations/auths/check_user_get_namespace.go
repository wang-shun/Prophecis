// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CheckUserGetNamespaceHandlerFunc turns a function with the right signature into a check user get namespace handler
type CheckUserGetNamespaceHandlerFunc func(CheckUserGetNamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckUserGetNamespaceHandlerFunc) Handle(params CheckUserGetNamespaceParams) middleware.Responder {
	return fn(params)
}

// CheckUserGetNamespaceHandler interface for that can handle valid check user get namespace params
type CheckUserGetNamespaceHandler interface {
	Handle(CheckUserGetNamespaceParams) middleware.Responder
}

// NewCheckUserGetNamespace creates a new http.Handler for the check user get namespace operation
func NewCheckUserGetNamespace(ctx *middleware.Context, handler CheckUserGetNamespaceHandler) *CheckUserGetNamespace {
	return &CheckUserGetNamespace{Context: ctx, Handler: handler}
}

/*CheckUserGetNamespace swagger:route GET /cc/v1/auth/access/admin/users/{username} Auths checkUserGetNamespace

auth by username .

Optional extended description in Markdown.

*/
type CheckUserGetNamespace struct {
	Context *middleware.Context
	Handler CheckUserGetNamespaceHandler
}

func (o *CheckUserGetNamespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckUserGetNamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
