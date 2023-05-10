/*
 * Copyright (c) 2022-2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/michaelcoll/sqlc-addon/internal/addon"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "write",
	Short: "Writes the new code",
	Long: `
Writes the new code`,
	Run: func(cmd *cobra.Command, args []string) {
		err := addon.WriteTemplate(".", "connect.go.gotmpl")
		if err != nil {
			fmt.Printf("%s Can't write connect.go file : %v\n", color.RedString("✗"), err)
		}

		err = addon.WriteTemplate(".", "migration.go.gotmpl")
		if err != nil {
			fmt.Printf("%s Can't write mgration.go file : %v\n", color.RedString("✗"), err)
		}

		fmt.Printf("%s file connect.go and mgration.go files created 🎉\n", color.GreenString("✓"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
