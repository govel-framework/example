/*
* This file is an example of how to use govel.
 */
package main

import (
	"github.com/govel-framework/example/actions"
	"github.com/govel-framework/govel"
	"github.com/govel-framework/lamb"
)

func main() {
	// here you can define your routes, govel currently supports GET, POST, PUT, DELETE
	// soon we will add support for other methods
	govel.Get("/", actions.Index).Name("index")

	govel.Post("/post-example", actions.Post)

	// an example of how lamb works
	govel.Get("/render/{name}", actions.Render)

	// You can extend govel (similar to how lamb, our templating engine, and lion, our ORM, do) by adding your own custom module to the InitModules function.
	// This function takes a list of functions as a parameter, and each of these functions accepts a parameter of type map[interface{}]interface{} and returns an error.
	// The map passed to these functions contains the configuration data from the configuration YAML file.
	// This allows you to validate or load any necessary configuration for your library.
	govel.InitModules(lamb.Init)

	// govel starts with a configuration file, it must be a YAML file where you can define the port and other options
	// LoadConfigFile also starts the server
	govel.LoadConfigFIle("./config.yaml")
}
