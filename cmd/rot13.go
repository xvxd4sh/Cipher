/*
Copyright © 2020 xvxd4sH

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

var read bool = false
var write bool = false

// rot13Cmd represents the rot13 command
var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "Rotational based on 13 index ",
	Long: `a simple letter substitution cipher that replaces a letter with the 13th letter after it, in the alphabet. 
	ROT13 is a special case of the Caesar cipher which was developed in ancient Rome. 
	
	Difficulty - very weak
	
	Example: 
		cipher rot13 -r []string
		cipher rot13 -w []string `,
	Run: func(cmd *cobra.Command, args []string) {
		read, err := cmd.Flags().GetBool("read")
		write, err2 := cmd.Flags().GetBool("write")
		if (read == true) && (write == true) {
			fmt.Println("Choose One please")
		} else if read == true {
			//fmt.Println("reading secret file")
			srv.Rot13Read(args)
			if err != nil {
				fmt.Println(err)
			}
		} else if write == true {
			//fmt.Println("writing a secret file")
			srv.Rot13Write(args)
			if err != nil {
				fmt.Println(err2)
			}
		} else {
			fmt.Println("error")
		}
	},
}

func init() {
	rootCmd.AddCommand(rot13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rot13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rot13Cmd.Flags().BoolP("read", "r", false, "decipher the sceret message")
	rot13Cmd.Flags().BoolP("write", "w", false, "cipher the sceret message")

}
