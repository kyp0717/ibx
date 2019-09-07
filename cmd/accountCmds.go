/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>
limitations under the License.
*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"ibx/client"
	"net/http"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "get list of accounts",
	Long:  "get list of accounts",
	Run: func(cmd *cobra.Command, args []string) {
		acct := client.Default.Account
		list, err := acct.GetIserverAccounts(nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", list)

	},
}

func init() {
	setPolicy()
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
