package controller

import (
	"github.com/goadesign/goa"
	"github.com/yu-ichiko/goa-security-example/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Get runs the get action.
func (c *OperandsController) Get(ctx *app.GetOperandsContext) error {
	defer ctx.Request.Body.Close()
	return ctx.OK([]byte("ok"))
}
