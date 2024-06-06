package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resource struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Links []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func getResources(c echo.Context) error {
	resources := []Resource{
		{
			ID:    "1",
			Name:  "Resource 1",
			Value: "Value 1",
			Links: []Link{
				{"self", "/api/v1/resources/1"},
			},
		},
		{
			ID:    "2",
			Name:  "Resource 2",
			Value: "Value 2",
			Links: []Link{
				{"self", "/api/v1/resources/2"},
			},
		},
	}
	return c.JSON(http.StatusOK, resources)
}

func createResource(c echo.Context) error {
	resource := new(Resource)
	if err := c.Bind(resource); err != nil {
		return err
	}
	resource.ID = "new-id" // This should be generated
	resource.Links = []Link{
		{"self", "/api/v1/resources/" + resource.ID},
	}
	return c.JSON(http.StatusCreated, resource)
}

func updateResource(c echo.Context) error {
	id := c.Param("id")
	resource := new(Resource)
	if err := c.Bind(resource); err != nil {
		return err
	}
	resource.ID = id
	resource.Links = []Link{
		{"self", "/api/v1/resources/" + resource.ID},
	}
	return c.JSON(http.StatusOK, resource)
}

func deleteResource(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Resource deleted",
		"id":      id,
	})
}
