package actions

import (
	"github.com/govel-framework/govel"
	"github.com/govel-framework/lamb"
)

func Render(c *govel.Context) {
	name := c.Param("name")

	// all lamb templates must end with .lamb.html
	lamb.Render(c, "welcome", govel.Map{
		"name": name,
	})
}
