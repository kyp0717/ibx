/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ibx/client"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out of ib web portal",
	Long:  "log out of ib web portal",
	Run: func(cmd *cobra.Command, args []string) {
		logoutOK, err := client.Default.Session.PostLogout(nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", logoutOK)

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
	// Here you will define your flags and configuration settings.
}
