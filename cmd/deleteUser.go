// Copyright © 2018 Dennis273 <dennisc695@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "Delete current User",
	Long: `Delete the user that is currently logged in.
	This operation would automatically remove you from every meetings that you participated in, and cancel meetings you holded.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteUser called")
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
