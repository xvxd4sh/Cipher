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

// DoubleColumnarCmd represents the DoubleColumnar command
var doublecolumnarCmd = &cobra.Command{
	Use:   "doublecolumnar",
	Short: "using 2 different key word to cipher the secret message",
	Long: `creating a cipher text based on a dual key base transpositional system
	
	difficulty -- normal 
	
	Example: 
		cipher doublecolumnar -k string -k string -r []string 
		cipher doublecolumnar -k string -k string -w []string`,
	Run: func(cmd *cobra.Command, args []string) {
		read, _ := cmd.Flags().GetBool("read")
		write, _ := cmd.Flags().GetBool("write")
		key, _ := cmd.Flags().GetStringSlice("key")

		key1 := key[0]
		key2 := key[1]

		if read {
			srv.DcolRead(args, key1, key2)
		} else if write {
			srv.DcolWrite(args, key1, key2)

		} else {
			fmt.Println("error")
		}
	},
}

func init() {
	rootCmd.AddCommand(doublecolumnarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DoubleColumnarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DoubleColumnarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doublecolumnarCmd.Flags().BoolP("read", "r", false, "reading the secret message")
	doublecolumnarCmd.Flags().BoolP("write", "w", false, " writing the secret message")
	doublecolumnarCmd.Flags().StringSliceP("key", "k", []string{}, "2 string keys for base ")
	doublecolumnarCmd.MarkFlagRequired("key")
}
