package webserver

import (
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/stevenkl/wkr/pkg/database"
	"github.com/stevenkl/wkr/pkg/models"
)

func jobsIndexHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	group := claims["group"].(string)

	if group != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

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
	runID, err := strconv.Atoi(c.Params("run_id"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	execution := new(models.ExecutionResultModel)
	err2 := database.Instance.GetExecutionResult(jobID, runID, execution)
	if err2 != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   execution,
	})
}

func createNewJobHandler(c *fiber.Ctx) error {
	job := new(models.JobModel)
	if err := c.BodyParser(job); err != nil {
		return err
	}
	job.ID = xid.New()
	dt := time.Now()
	// 2021-02-17T16:45:48.875Z
	job.CreatedAt = dt.Format("2006-01-02T15:04:05.000Z")

	if err := database.Instance.SaveJob(job); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// create json file
	return c.SendStatus(fiber.StatusCreated)
}

func executeJobHandler(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	job := new(models.JobModel)
	err := database.Instance.GetJob(jobID, job)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func deleteJobHandler(c *fiber.Ctx) error {
	job := new(models.JobModel)
	if err := c.BodyParser(job); err != nil {
		return err
	}
	err2 := database.Instance.DeleteJob(job)
	if err2 != nil {
		return err2
	}
	return c.SendStatus(fiber.StatusNoContent)
}
