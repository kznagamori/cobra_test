/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subcommand01Cmd represents the subcommand01 command
var subcommand01Cmd = &cobra.Command{
	Use:   "subcommand01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		toggle, _ := cmd.Flags().GetBool("toggle")
		if toggle {
			fmt.Println("subcommand01 toggle called")
		}
		optionA, _ := cmd.Flags().GetString("optionA")
		if optionA != "" {
			fmt.Println("subcommand01 optionA called")
		}
		optionB, _ := cmd.Flags().GetString("optionB")
		if optionB != "" {
			fmt.Println("subcommand01 optionB called")
		}
		optionC, _ := cmd.Flags().GetString("optionC")
		if optionC != "" {
			fmt.Println("subcommand01 optionC called")
		}
		fmt.Println("subcommand01 called")
	},
}

func init() {
	rootCmd.AddCommand(subcommand01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subcommand01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	subcommand01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	subcommand01Cmd.Flags().StringP("optionA", "a", "", "Option A")
	subcommand01Cmd.Flags().StringP("optionB", "b", "", "Option B")
	subcommand01Cmd.Flags().StringP("optionC", "c", "", "Option C")
	subcommand01Cmd.MarkFlagsMutuallyExclusive("optionA", "optionC")

}
