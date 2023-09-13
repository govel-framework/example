package actions

import (
	"fmt"
	"net/http"

	"github.com/govel-framework/govel"
)

// A basic hello world action.
func Index(c *govel.Context) {
	c.Text(http.StatusOK, "Hello World!")
}

// A basic POST action.
func Post(c *govel.Context) {
	form, err := c.NewForm()

	if err != nil {
		// could not parse form
		c.Text(http.StatusBadRequest, "Bad Request")
		return
	}

	name := form.Get("name")

	fmt.Println("Name: ", name)
}

// A basic form validation action.
func Validate(c *govel.Context) {
	form, err := c.NewForm()

	if err != nil {
		// could not parse form
		c.Text(http.StatusBadRequest, "Bad Request")
		return
	}

	// validate the form
	rules := govel.Map{
		"name":     []string{"required", "max:255", "alpha_num"},
		"username": []string{"required", "max:255"},
		"password": []string{"required", "min:8", "max:255", "confirm"},
	}

	onError := govel.OnError{
		"name": govel.SMap{
			"required":  "Name is required",
			"max":       "Name is too long",
			"alpha_num": "Name can only contain letters and numbers",
		},
		"username": govel.SMap{
			"required": "Username is required",
			"max":      "Username is too long",
		},
		"email": govel.SMap{
			"required": "Email is required",
			"email":    "Email is invalid",
			"max":      "Email is too long",
			"unique":   "Email already exists, please choose another one",
		},
		"password": govel.SMap{
			"required": "Password is required",
			"min":      "Password is too short",
			"max":      "Password is too long",
			"confirm":  "Passwords do not match",
		},
	}

	data, errors := form.Validate(rules, onError)

	if len(errors) > 0 {
		// validation failed
		c.Text(http.StatusBadRequest, "Bad Request")
		return
	}

	// validation passed
	c.Text(http.StatusOK, "Success")
	fmt.Println(data)
}
