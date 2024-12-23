/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reqUrl := resolveUrl(args[0])
		client := &http.Client{}

		req, _ := http.NewRequest(http.MethodPut, reqUrl.String(), nil)
		req.Header.Set("X-aws-ec2-metadata-token-ttl-seconds", "21600")
		res, err := client.Do(req)
		if err != nil {
			os.Exit(1)
		}
		defer res.Body.Close()

		t, _ := io.ReadAll(res.Body)
		storeToken(string(t))
	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
