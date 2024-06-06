package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Update Project
// @Description Получает (или обновляет) все issues из проекта с ключом 'projectKey' и заносит в базу данных.
// @Param project query string true "Project Key"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/jira/updateProject [get]
func updateProject(c echo.Context) error {
	projectKey := c.QueryParam("project")
	if projectKey == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "project parameter is required"})
	}

	err := fetchAndStoreJiraProject(projectKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Project updated"})
}

// @Summary List Projects
// @Description Показать все доступные проекты с поддержкой пагинации и поиска.
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Param search query string false "Search"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/v1/jira/projects [get]
func listProjects(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	search := c.QueryParam("search")

	projects, pageInfo, err := getProjectsFromDB(limit, page, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"projects": projects, "pageInfo": pageInfo})
}
