package webserver

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stevenkl/wkr/pkg/database"
	"github.com/stevenkl/wkr/pkg/models"
)

func jobsIndexHandler(c *fiber.Ctx) error {
	// jobs := make([]models.JobModel, 0)

	jobs := database.Instance.GetAll()
	if jobs == nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Println(jobs)

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   jobs,
	})
}

func jobsDetailsHandler(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	job := new(models.JobModel)
	err := database.Instance.Get(jobID, job)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"status": "succss",
		"data":   job,
	})
}
