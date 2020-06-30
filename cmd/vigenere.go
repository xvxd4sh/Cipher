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

	"github.com/spf13/cobra"
)

// vigenereCmd represents the vigenere command
var vigenereCmd = &cobra.Command{
	Use:   "vigenere",
	Short: "cipher and decypher using a polyalphabetic cipher",
	Long: `the use of multiple caesar cipher to encode and decode messages 
	
	difficulty -- easy`,
	Run: func(cmd *cobra.Command, args []string) {
		read, _ := cmd.Flags().GetBool("read")
		write, _ := cmd.Flags().GetBool("write")
		key, _ := cmd.Flags().GetString("key")

		//fmt.Println(key)
		if read {
			//fmt.Println(args)
			srv.VigRead(args, key)
		} else if write {
			//fmt.Println(args)
			srv.VigWrite(args, key)
		}
		//fmt.Println("vigenere called")
	},
}

func init() {
	rootCmd.AddCommand(vigenereCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vigenereCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vigenereCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	vigenereCmd.Flags().BoolP("read", "r", false, "decrypting the secret message")
	vigenereCmd.Flags().BoolP("write", "w", false, "encrypting the secret message")
	vigenereCmd.Flags().StringP("key", "k", "", "base key")
}
