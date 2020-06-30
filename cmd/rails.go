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

// railsCmd represents the rails command
var railsCmd = &cobra.Command{
	Use:   "rails",
	Short: "cipher based on a transpositional cipher",
	Long: `using location manipulation of visual allignment of imaginary path drawn
	 on a grid to convert string to cipher
	 
	 difficulty -- easy`,
	Run: func(cmd *cobra.Command, args []string) {
		read, _ := cmd.Flags().GetBool("read")
		write, _ := cmd.Flags().GetBool("write")
		key, _ := cmd.Flags().GetInt("key")
		if read {
			srv.RailsRead(args, key)
		} else if write {
			srv.RailsWrite(args, key)
		} else {
			fmt.Println("error")
		}
		//fmt.Println("rails called")
	},
}

func init() {
	rootCmd.AddCommand(railsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// railsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// railsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	railsCmd.Flags().BoolP("read", "r", false, "decrypt the secret message")
	railsCmd.Flags().BoolP("write", "w", false, "encrypt the secret message")
	railsCmd.Flags().IntP("key", "k", 0, "base key")

}
