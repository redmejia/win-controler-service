package api

import "log"

type ApiConfig struct {
	Port              string
	Infolog, Errorlog *log.Logger
}
