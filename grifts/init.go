package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mitch000001/s3-gui/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
