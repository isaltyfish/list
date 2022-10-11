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
	Short: "列出指定文件下的内容",
	Long:  `列出指定文件下所有的文件名称，包括文件名和目录名。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 注意：需要通过长名称获取选项的值
		value, _ := cmd.Flags().GetBool("lst")
		targetDir, _ := cmd.Flags().GetString("directory")
		if targetDir == "" {
			wd, _ := os.Getwd()
			targetDir = wd
		}

		sep := " "
		if value {
			sep = "\n"
		}
		listDirContent(sep, targetDir)
	},
}

func listDirContent(sep string, targetDir string) {
	files, _ := ioutil.ReadDir(targetDir)
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("lst", "l", false, "列出文件名，并使用换行符分隔")
	rootCmd.Flags().StringP("directory", "d", "", "目标文件夹")
}
