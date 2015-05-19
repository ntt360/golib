package web

import (
	"net/http"
)

//------------------------------
// type T_Context
//------------------------------
type T_Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter

	DisableView bool
	ViewName    string
	ViewData    map[string]interface{}
}

func (this *T_Context) AssignViewData(key string, value interface{}) {
	this.DisableView = false
	this.ViewData[key] = value
}

func (this *T_Context) SetViewName(name string) {
	this.DisableView = false
	this.ViewName = name
}
