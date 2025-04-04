package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a colored string
	green := color.New(color.FgGreen).SprintFunc()
	message := green("Hello, World!")

	// Log the colored message
	logrus.Info(message)
}
