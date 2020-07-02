/*
Copyright © 2020 xvxd4sh

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
	"cipher/cmd/srv"
	"fmt"

	"github.com/spf13/cobra"
)

// columnarCmd represents the columnar command
var columnarCmd = &cobra.Command{
	Use:   "columnar",
	Short: "cipher based on a dual-layer transpositional cipher",
	Long: `Using a matrix of a base word to cipher a secret message. 
	
	dificulty -- normal
	
	Example: 
		cipher columnar -k string -r []string 
		cipher columanr -k string -w []string 
		`,
	Run: func(cmd *cobra.Command, args []string) {
		read, _ := cmd.Flags().GetBool("read")
		write, _ := cmd.Flags().GetBool("write")
		key, _ := cmd.Flags().GetString("key")
		if read {
			reads := srv.ColumnaRead(args, key)
			fmt.Println(reads)
		} else if write {
			writes := srv.ColumnaWrite(args, key)
			fmt.Println(writes)
		} else {
			fmt.Println("error")
		}
	},
}

func init() {
	rootCmd.AddCommand(columnarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// columnarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// columnarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	columnarCmd.Flags().BoolP("read", "r", false, "reading the secret message")
	columnarCmd.Flags().BoolP("write", "w", false, "writing the secret message")
	columnarCmd.Flags().StringP("key", "k", "", "base word for key")
	columnarCmd.MarkFlagRequired("key")
}
