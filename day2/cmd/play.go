/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verscheures/AoC2023/day2/game"
)

var Source string

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		game.Play(Source)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.
	playCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	if err := playCmd.MarkFlagRequired("source"); err != nil {
		fmt.Println(err.Error())
		return
	}
}
