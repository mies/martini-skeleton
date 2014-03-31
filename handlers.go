package main

import (
	"github.com/martini-contrib/render"

)

func Index(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})

}
