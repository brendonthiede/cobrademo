/*
Copyright © 2021 Brendon Thiede <brendon.thiede@gmail.com>

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
	"fmt"

	"github.com/spf13/cobra"
)

var slice []string
var arr []string

// commaTestCmd represents the commaTest command
var commaTestCmd = &cobra.Command{
	Use:   "commaTest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commaTest called")
		fmt.Println("\nslice value:")
		for _, v := range slice {
			fmt.Println(v)
		}
		fmt.Println("\narray value:")
		for _, v := range arr {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(commaTestCmd)

	commaTestCmd.Flags().StringSliceVarP(&slice, "slice", "s", nil, "Testing a string slice (try commas)")
	commaTestCmd.Flags().StringArrayVarP(&arr, "array", "a", nil, "Testing a string array (try commas)")

}
