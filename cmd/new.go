// Copyright (C) 2021- Soroush Safari <soroush_safarii@yahoo.com>
//
// This file is part of Odin.
//
// Odin is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Odin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
// along with Odin.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"log"
	"os"

	"github.com/coci/odin/odin"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new post",
	Long:  `This will create new post template inside the content directory.`,
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "" {
			log.Println("you have to provide title for your post.")
			os.Exit(0)
		}

		odin.New(args[0])

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
