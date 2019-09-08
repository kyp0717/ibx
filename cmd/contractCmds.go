/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ibx/client"
)

// contractCmdsCmd represents the contractCmds command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "A brief description of your command",
	Long:  "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("contractCmds called")
		ctx, err := client.Default.Contract.PostSecdefSearch(nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", ctx)

	},
}

func init() {
	rootCmd.AddCommand(contractCmd)
	// contractCmdsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
