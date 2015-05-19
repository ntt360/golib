package web

import (
	"net/http"
)

//------------------------------
//type I_Controller
//------------------------------
type I_Controller interface {
	Init(context *T_Context)
	Before(context *T_Context)
	NotFound(context *T_Context)
	After(context *T_Context)
	UnInitialize(context *T_Context)
}

//------------------------------
// type T_Controller
//------------------------------
type T_Controller struct {
}

// TODO init controller
func (this *T_Controller) Init(context *T_Context) {

}

// Before runs before request function execution.
func (this *T_Controller) Before(context *T_Context) {

}

// not found action
func (this *T_Controller) NotFound(context *T_Context) {
	context.DisableView = true
	http.Error(context.ResponseWriter, "404 not found", http.StatusNotFound)
}

// After runs after request function execution.
func (this *T_Controller) After(context *T_Context) {

}

// TODO uninitialize controller
func (this *T_Controller) UnInitialize(context *T_Context) {

}

//------------------------------
// type T_NotFoundController
//------------------------------
type T_NotFoundController struct {
	T_Controller
}
