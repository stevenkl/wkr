package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stevenkl/wkr/pkg/database"
	"github.com/stevenkl/wkr/pkg/models"
)

func jobsIndexHandler(c *fiber.Ctx) error {
	// jobs := make([]models.JobModel, 0)

	jobs, err := database.Instance.GetAllJobs()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   jobs,
	})
}

func jobsDetailsHandler(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	job := new(models.JobModel)
	err := database.Instance.GetJob(jobID, job)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"status": "succss",
		"data":   job,
	})
}

func jobExecutionsHandler(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	runID := c.Params("run_id")

	execution := new(models.ExecutionResultModel)
	err := database.GetExecutionResult(jobID, runID, execution)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data": execution,
	})
}