package render

import (
	"bread-bake/pkg/config"
	"bread-bake/pkg/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// variable for global config
var conf *config.AppConfig

func SetConfig(config *config.AppConfig) {
	conf = config
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template
	if conf.UseCache {
		cache = conf.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}

	template, ok := cache[tmpl]
	if !ok {
		log.Fatal("Error reading cache")
	}
	td = AddDefaultData(td)
	err := template.Execute(w, td)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
