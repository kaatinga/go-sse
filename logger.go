package sse

import (
	"github.com/kaatinga/dummylogger"
)

var log = dummylogger.Get

func Init(logger dummylogger.I) {
	dummylogger.Set(logger)
}
