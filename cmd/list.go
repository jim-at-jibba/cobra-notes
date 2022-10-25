/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jim-at-jibba/studybuddy/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See a list of all notes you've added",
	Long:  "See a list of all notes you've added",
	Run: func(cmd *cobra.Command, args []string) {
		data.DispayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
