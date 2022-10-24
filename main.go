/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/jim-at-jibba/studybuddy/cmd"
	"github.com/jim-at-jibba/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
