/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verscheures/AoC2023/day3/fix"
)

var Source string

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		engine := &fix.Engine{}
		fmt.Println("init engine")
		engine.Init(Source)
		fmt.Println("sum:", engine.Sum)
		fmt.Println("ratio:", engine.Ratio)
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)

	fixCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	if err := fixCmd.MarkFlagRequired("source"); err != nil {
		fmt.Println(err.Error())
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
