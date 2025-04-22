/*
Copyright © 2024 Jeff Kody <jeph@cscoding.io>

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
	"csserver/cmd/cstools/cmd/setup"
	"csserver/internal/config"
	"csserver/internal/gen"
	"fmt"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")

		//---TODO: do better
		mp := "/home/jeph/projects/cscoding21/csplanner/csserver/cmd/cstools/.cstools.yaml"

		err := gen.GenServices(mp)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Service generation complete")
		}

	},
}

// testdataCmd represents the testdataCmd command
var testdataCmd = &cobra.Command{
	Use:   "testdata",
	Short: "An operation that will setup test data within the application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testdata called")

		ctx := config.NewContext()

		err := setup.SetupTestData(ctx)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Test data generation complete")
		}

	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(testdataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
