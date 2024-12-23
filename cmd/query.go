/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type myJsonStruct struct {
	Path string `json:"path"`
}

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reqUrl := resolveUrl(args[0])

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
		if err != nil {
			return
		}

		str := retrieveToken()
		if len(str) != 0 {
			req.Header.Set("X-aws-ec2-metadata-token", str)
		}

		response, err := client.Do(req)
		if err != nil {
			return
		}
		defer response.Body.Close()

		buf, err := processResponseString(*response)
		if err != nil {
			fmt.Print(buf)
			return
		}

		fmt.Println(processOutput(reqUrl.RequestURI(), buf))
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
