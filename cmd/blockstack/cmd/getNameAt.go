// Copyright © 2017 Jack Zampolin <jack.zampolin@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getNameAtCmd represents the getNameAt command
var getNameAtCmd = &cobra.Command{
	Use:   "get_name_at",
	Short: "Calls the get_name_at rpc method on the configured host",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		res, err := client.GetNameAt(args[0], validateIntArg(args[1], "blockHeight"))
		if err != nil {
			fmt.Println("Error, TODO: Make some real JSON errors for the blockstack library", err)
		}
		fmt.Println(res.JSON())
	},
}

func init() {
	RootCmd.AddCommand(getNameAtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNameAtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNameAtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
