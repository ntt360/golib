package web

import (
	"html/template"
	"net/http"
)

var (
	_view_path  string
	_view_parse *template.Template
)

func init() {

}

func SetViewPath(path string) {
	_view_path = path
	_view_parse = template.Must(template.ParseGlob(_view_path + "*/*"))
}

//------------------------------
//type I_Controller
//------------------------------
type I_Controller interface {
}

//------------------------------
// type T_Controller
//------------------------------
type T_Controller struct {
}

// render template
// TODO deal template file not exist
func (this *T_Controller) Render(w http.ResponseWriter, tpl_name string, context map[string]interface{}) {
	tpl_file := _view_path + tpl_name
	tpl, err := template.ParseFiles(tpl_file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, context)
	return
}

// show view
// TODO deal template file not exist
func (this *T_Controller) ShowView(w http.ResponseWriter, view_name string, context map[string]interface{}) {
	err := _view_parse.ExecuteTemplate(w, view_name, context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *T_Controller) NotFound(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "404 not found", http.StatusNotFound)
}

//------------------------------
// type T_NotFoundController
//------------------------------
type T_NotFoundController struct {
	T_Controller
}
