package web

import (
	"html/template"
	"net/http"
)

var (
	_view_path  string
	_view_cache *template.Template
)

func SetViewPath(path string) {
	_view_path = path
	_view_cache = template.Must(template.ParseGlob(_view_path + "*/*"))
}

// render template file
func RenderFile(file string, context *T_Context) {
	file = _view_path + file
	tpl, err := template.ParseFiles(file)
	if err != nil {
		// TODO deal template file not exist
		http.Error(context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(context.ResponseWriter, context.ViewData)
	return
}

// show view
func ShowView(context *T_Context) {
	err := _view_cache.ExecuteTemplate(context.ResponseWriter, context.ViewName, context.ViewData)
	if err != nil {
		// TODO deal template file not exist
		http.Error(context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
