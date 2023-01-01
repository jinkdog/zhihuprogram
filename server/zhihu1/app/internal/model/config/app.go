package config

import "net/http"

type App struct {
	Cookie *Cookie `mapstructure:"cookie" yaml:"cookie"`
}

type Cookie struct { //包含了原始的http包中的cookie以及secret构成的新cookie
	Secret      string `mapstructure:"secret" yaml:"secret"`
	http.Cookie `mapstructure:",squash"`
}
