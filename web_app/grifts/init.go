package grifts

import (
	"github.com/abdev/huiskluis/web_app/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
