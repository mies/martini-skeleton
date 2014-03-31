package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", Index)
	m.Run()
}
