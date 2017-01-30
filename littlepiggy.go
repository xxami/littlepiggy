package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			context := Context{"Hello, world!"}
			tmpl.Execute(w, context)
		}
	})

	http.ListenAndServe(":8000", nil)
}

type Context struct {
	Message string
}

const doc = `
<!DOCTYPE html>
<html>
<head><title></title></head>
<body>
    <h1>{{.Message}}</h1>
</body>
</html>
`
