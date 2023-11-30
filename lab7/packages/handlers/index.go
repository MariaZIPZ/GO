package handlers

import (
	"fmt"
	"lab7/packages/templates"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.PageStart, templates.IndexTemplate)
	fmt.Fprint(w, templates.PageEnd)
}
