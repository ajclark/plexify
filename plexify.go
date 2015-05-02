package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"strconv"
)

// Store runtime argument for plex scanner command
var plexCmd string

// Required JSON parameters
type requestJSON struct {
	Path       string `json:"path" binding:"required"`
	Section_ID int    `json:"section_id" binding:"required"`
}

// Make sure the path to Plex Media Scanner is specified at runtime
func init() {
	flag.StringVar(&plexCmd, "p", "/usr/local/bin/pms", "Path to Plex Media Scanner binary")
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if _, err := os.Stat(plexCmd); err != nil {
		panic(err)
	}
}

// Execute the plex media scanner with a path and a plex section id
func updatePlex(path string, section_id int, plexCmd string) error {
	s_id := strconv.Itoa(section_id)
	cmd := exec.Command(plexCmd, "-s", "-r", "-c", s_id, "-d", path)
	err := cmd.Start()
	return err
}

func main() {
	r := gin.Default()

	r.POST("/scan", func(c *gin.Context) {
		var json requestJSON
		if !c.Bind(&json) {
			c.AbortWithStatus(400)
			return
		}

		err := updatePlex(json.Path, json.Section_ID, plexCmd)

		if err != nil {
			c.JSON(500, gin.H{"status": err})
		} else {
			c.JSON(200, gin.H{"status": "updating"})
		}
	})
	r.Run(":8080")
}
