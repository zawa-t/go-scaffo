/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zawa-t/go-scaffo/src/generate"
)

var arch string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate [name]", // メインコマンド名
	Aliases: []string{"g"},     // エイリアスの定義
	Short:   "Generate a new application scaffold",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := generate.Arg{
			AppName:  args[0],
			ArchName: arch,
		}
		if err := generate.Scaffold(arg); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Application scaffold '%s' created successfully!\n", arg.AppName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVar(&arch, "arch", "", "Specify architecture name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
