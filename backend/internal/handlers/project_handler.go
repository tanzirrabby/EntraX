package handlers

import (
	"net/http"
	"time"

	"github.com/example/entrax/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectHandler struct{}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func NewProjectHandler() ProjectHandler {
	return ProjectHandler{}
}

func (h ProjectHandler) ListProjects(c *gin.Context) {
	projects := []models.Project{
		{
			ID:          "3f6f48fa-aaa1-4324-9030-a5f11fd4d45f",
			Name:        "Identity Modernization",
			Description: "Enterprise SSO rollout and app migration",
			CreatedAt:   time.Now().Add(-48 * time.Hour),
		},
	}
	c.JSON(http.StatusOK, projects)
}

func (h ProjectHandler) CreateProject(c *gin.Context) {
	var request CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	project := models.Project{
		ID:          uuid.NewString(),
		Name:        request.Name,
		Description: request.Description,
		CreatedAt:   time.Now(),
	}

	c.JSON(http.StatusCreated, project)
}
