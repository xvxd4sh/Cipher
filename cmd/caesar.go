/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// caesarCmd represents the caesar command
var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "encryption and decryption based on caesar cipher",
	Long: `The use of keys to shift the points and encrypting and decrypting based on that singular key 
	
	difficulty -- weak
	
	Example: 
		cipher caesar -k int -r []string 
		cipher caesar -k int -w []string `,
	Run: func(cmd *cobra.Command, args []string) {
		read, _ := cmd.Flags().GetBool("read")
		write, _ := cmd.Flags().GetBool("write")
		key, _ := cmd.Flags().GetInt("key")
		if read {
			//fmt.Println("reading")
			srv.CaesarRead(args, key)
		} else if write {
			//fmt.Println("writing")
			srv.CaesarWrite(args, key)
		} else {
			fmt.Println("error")
		}
		//fmt.Println("caesar called")
	},
}

func init() {
	rootCmd.AddCommand(caesarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// caesarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// caesarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	caesarCmd.Flags().BoolP("read", "r", false, "decrypting the secret message")
	caesarCmd.Flags().BoolP("write", "w", false, "encrypting the secret message")
	caesarCmd.Flags().IntP("key", "k", 0, "base key")
	//caesarCmd.SetUsageTemplate("Cipher caesar -k key -r/w 'secret messages'")
	//caesarCmd.SetHelpTemplate("Cipher caesar -k key -r/w 'secret messages'")
	caesarCmd.MarkFlagRequired("key")
}
