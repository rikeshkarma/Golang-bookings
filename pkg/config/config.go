package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//AppConfig holds application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache bool
	InfoLog *log.Logger
	InProd bool
	Session *scs.SessionManager
}