package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/irsyadfata99/go-sec-app/pkg/config"
	"github.com/irsyadfata99/go-sec-app/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) error {
    var tc map[string]*template.Template
    
    // Get the template cache from the app config
    if app.UseCache {
        tc = app.TemplateCache
    } else {
        // Rebuild the cache for development
        tc, _ = CreateTemplateCache()
    }
    
    // Get the requested template from cache
    t, ok := tc[tmpl]
    if !ok {
        log.Printf("Template %s not found in cache", tmpl)
        // List all available templates for debugging
        for k := range tc {
            log.Printf("Available template: %s", k)
        }
        return fmt.Errorf("could not get template %s from cache", tmpl)
    }
    
    // Execute the template
    err := t.Execute(w, td)
    if err != nil {
        return err
    }
    
    return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
    myCache := map[string]*template.Template{}
    
    // Get all .page.tmpl files from templates directory
    pages, err := filepath.Glob("./templates/*.page.tmpl")
    if err != nil {
        return myCache, err
    }
    
    // Loop through all page templates and add them to the cache
    for _, page := range pages {
        name := filepath.Base(page)
        fmt.Printf("Page being loaded: %s\n", name) // Add this debug line
        
        // Parse the template
        ts, err := template.New(name).ParseFiles(page)
        if err != nil {
            return myCache, err
        }
        
        // Add layouts
        matches, err := filepath.Glob("./templates/*.layout.tmpl")
        if err != nil {
            return myCache, err
        }
        
        if len(matches) > 0 {
            ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
            if err != nil {
                return myCache, err
            }
        }
        
        myCache[name] = ts
    }
    
    return myCache, nil
}
