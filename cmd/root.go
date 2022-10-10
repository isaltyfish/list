/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "list",
	Short: "列出当前文件下的内容",
	Long: `列出当前文件下所有的文件名称，包括文件名和目录名。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 注意：需要通过长名称获取选项的值
		value, _ := cmd.Flags().GetBool("lst")
		sep := " "
		if value {
			sep = "\n"
		}
		listDirContent(sep)
	},
}

func listDirContent(sep string) {
	wd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(wd)
	fileNames := make([]string, 0, len(files))
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	fmt.Println(strings.Join(fileNames, sep))
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.list.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("lst", "l", false, "列出文件名，并使用换行符分隔")
}
