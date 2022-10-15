/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "查找指定目录下指定前缀开头的文件",
	Long: `查找指定目录下指定前缀开头的文件。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 这里实现命令逻辑
		prefix, _ := cmd.Flags().GetString("prefix")
		dir, _ :=cmd.Flags().GetString("directory")
		searchByPrefix(dir, prefix)
	},
}

func searchByPrefix(dir string, prefix string) {
	files, _ := ioutil.ReadDir(dir)
	fileNames := make([]string, 0, len(files))
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			fileNames = append(fileNames, file.Name())
		}
	}
	fmt.Println(strings.Join(fileNames, "\n"))
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("prefix", "p", "", "待查找的文件名前缀")
	searchCmd.Flags().StringP("directory", "d", "", "目标查找目录")
}
