package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Collection struct {
	gorm.Model
	Name              string
	ProjectID         uint
	HttpRequests      []HttpRequest
	GrpcRequests      []GrpcRequest
	WebsocketRequests []WebsocketRequest
}
