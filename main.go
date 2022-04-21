package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You should specify listening path.")
	}

	r := gin.Default()

	exec := CreateCommandHandler(os.Args[2:])
	r.POST("/", func(c *gin.Context) {
		var req CommandRequest
		if err := c.BindJSON(&req); err != nil {
			err = fmt.Errorf("bad request: %w", err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err := exec(req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.Status(http.StatusOK)
	})

	r.DELETE("/", func(ctx *gin.Context) {
		os.Exit(0)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(os.Args[1])
}

type CommandRequest struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func CreateCommandHandler(whitelist []string) func(req CommandRequest) error {
	return func(req CommandRequest) error {
		for _, v := range whitelist {
			if v == req.Command {
				command := exec.Command(req.Command, req.Args...)
				return command.Start()
			}
		}

		return fmt.Errorf("not allowed command")
	}
}
