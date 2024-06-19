package models

import (
	"github.com/chenmingyong0423/go-mongox"
)

type Name struct {
	mongox.Model `bson:",inline"`
	FirstName    string `bson:"first_name"`
	LastName     string `bson:"last_name"`
}
