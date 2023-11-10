package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type Task struct {
	Cron     string `json:"cron"`
	TaskName string `json:"task_name"`
	Exec     string `json:"exec"`
}

var TaskStorage []map[cron.EntryID]Task

func GetTasks(c *gin.Context) {
	var results []map[string]interface{}
	for _, e := range Cron.Entries() {
		results = append(results, map[string]interface{}{
			"Task ID": e.ID,
			// "task_name":      TaskStorage[e.ID],
			"Execution Time": e.Next,
		})
	}
	c.JSON(http.StatusOK, results)
}

func AddTask(c *gin.Context) {
	task := Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// add cron job
	eid, err := Cron.AddFunc(task.Cron, func() {
		ExecuteTask(task.Exec)
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(eid))
	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"message": "Success add task with ID : " + id,
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	eid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// remove cron job
	Cron.Remove(cron.EntryID(eid))

	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"message": `Delete Task with ID : ` + id,
	})
}
