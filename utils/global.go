package utils

import (
	"github.com/h4lim/go-sdk/app/types"
	"github.com/h4lim/go-sdk/database"
	opLogging "github.com/h4lim/go-sdk/logging"
)

var MMEN *types.MessageMap
var MMID *types.MessageMap
var DBModel *database.DBModel

var log = opLogging.MustGetLogger("go-sdk")
