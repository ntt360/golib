package web

import (
	"net/http"
	"reflect"
	"strings"
)

const (
	DEFAULT_CONTROLLER_NAME = "Index"
	DEFAULT_ACTION_NAME     = "Index"
	NOTFOUND_ACTION_NAME    = "NotFound"

	PATH_TRIM_STRING  = "/ "
	PATH_SPLIT_STRING = "/"

	EMPTY_STRING = ""

	AUTO_CALL_METHOD_INIT         = "Init"
	AUTO_CALL_METHOD_BEFORE       = "Before"
	AUTO_CALL_METHOD_AFTER        = "After"
	AUTO_CALL_METHOD_UNINITIALIZE = "UnInitialize"
)

var (
	_controller_manager map[string]I_Controller
)

func init() {
	_controller_manager = make(map[string]I_Controller)
}

//------------------------------
// type T_Router
//------------------------------
type T_Router struct {
	ControllerName string
	ActionName     string
}

func NewRouter() *T_Router {
	return &T_Router{}
}

// init controller name and action name from url path
func (this *T_Router) initControllerNameAndActionNameFromPath(path string) {
	this.ControllerName = DEFAULT_CONTROLLER_NAME
	this.ActionName = DEFAULT_ACTION_NAME

	path = strings.Trim(path, PATH_TRIM_STRING)

	if EMPTY_STRING != path {

		pathList := strings.Split(path, PATH_SPLIT_STRING)
		pathLength := len(pathList)

		if pathLength == 2 {
			this.ControllerName = strings.Title(pathList[0])
			this.ActionName = strings.Title(pathList[1])
		} else if pathLength == 1 {
			this.ControllerName = strings.Title(pathList[0])
		}

	}
}

// register controller into _controller_manager
func (this *T_Router) RegisterController(name string, c I_Controller) {
	name = strings.ToLower(name)
	_controller_manager[name] = c
}

// get controller by name from _controller_manager
func (this *T_Router) GetController(name string) (c I_Controller, ok bool) {
	name = strings.ToLower(name)
	c, ok = _controller_manager[name]
	return

}

// get current controller, return &T_NotFoundController{} if it not exist
func (this *T_Router) GetCurrentController() (c I_Controller) {
	c, ok := this.GetController(this.ControllerName)
	if !ok {
		c = &T_NotFoundController{}
	}
	return c
}

// dispatch request to matched action of controller
// call method: Init -> Before -> ActionName -> After -> UnInitialize
func (this *T_Router) Dispatch(w http.ResponseWriter, r *http.Request) {

	context := &T_Context{
		Request:        r,
		ResponseWriter: w,

		DisableView: false,
		ViewName:    strings.ToLower(this.ControllerName + "_" + this.ActionName),
		ViewData:    make(map[string]interface{}),
	}

	currentController := this.GetCurrentController()
	controllerReflect := reflect.ValueOf(currentController)

	param := []reflect.Value{reflect.ValueOf(context)}

	method := controllerReflect.MethodByName(this.ActionName)
	if (reflect.Value{} == method) {
		method = controllerReflect.MethodByName(NOTFOUND_ACTION_NAME)
	}

	// call action
	controllerReflect.MethodByName(AUTO_CALL_METHOD_INIT).Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_BEFORE).Call(param)
	method.Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_AFTER).Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_UNINITIALIZE).Call(param)

	// show view
	if !context.DisableView {
		ShowView(context)
	}
}

// ServeHTTP method
func (this *T_Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.initControllerNameAndActionNameFromPath(r.URL.Path)
	this.Dispatch(w, r)
}
