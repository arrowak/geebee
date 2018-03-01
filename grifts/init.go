package grifts

import (
	"github.com/arunko350/geebee/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
