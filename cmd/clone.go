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
	"github.com/bitxel/gitbk/config"
	"github.com/bitxel/gitbk/gitbk"

	"github.com/spf13/cobra"
)

var (
	url string
	branch string
	workdir string
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Project{
			URL: url,
			Branch: branch,
		    WorkDir: workdir,
		}
		gitbk.Clone(cfg)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringVarP(&url, "url", "u", "", "remote repo URL")
	cloneCmd.Flags().StringVarP(&branch, "branch", "b", "master", "default branch to push to")
	cloneCmd.Flags().StringVarP(&branch, "workdir", "w", "", "default branch to push to")
}
