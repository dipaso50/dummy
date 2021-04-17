/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"dummy/fileGenerator"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	numFiles   int
	destFilder string
	fileExt    string
	megas      int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "Dummy file generator",
	Long:  ` Dummy file generator, creates files with random names`,
	Run: func(cmd *cobra.Command, args []string) {
		fileGenerator.Generate(numFiles, destFilder, megas, fileExt)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().IntVarP(&numFiles, "numFiles", "n", 1, "Number of files to generate")

	rootCmd.Flags().StringVarP(&destFilder, "destinationFolder", "d", "", "Destination folder")
	rootCmd.MarkFlagRequired("destinationFolder")

	rootCmd.Flags().IntVarP(&megas, "megas", "m", 10, "Megas for each file")

	rootCmd.Flags().StringVarP(&fileExt, "fileExtension", "x", "", "Files will be generated with this file ext")
}
