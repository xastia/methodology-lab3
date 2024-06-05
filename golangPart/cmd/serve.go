package cmd

import (
	"fizzbuzz/lib"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run an http server to answer fizzbuzz queries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server...")

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		type QueryContext struct {
			QueryRaw string
			Query    *int
			Result   *string
			Error    error
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Received request:", r.Method)
			if r.Method != http.MethodPost {
				err = tmpl.Execute(w, nil)
				if err != nil {
					fmt.Println("Error executing template:", err)
				}
				return
			}

			context := QueryContext{
				QueryRaw: r.FormValue("query"),
				Query:    nil,
				Result:   nil,
				Error:    nil,
			}
			fmt.Println("Received query:", context.QueryRaw)
			query, err := strconv.Atoi(context.QueryRaw)
			if err == nil {
				context.Query = &query
				result := lib.FizzBuzz(*context.Query)
				context.Result = &result
				fmt.Println("Computed result:", *context.Result)
			} else {
				context.Error = err
				fmt.Println("Error converting query to int:", err)
			}

			err = tmpl.Execute(w, context)
			if err != nil {
				fmt.Println("Error executing template with context:", err)
			}
			return
		})

		port, _ := cmd.Flags().GetInt("port")
		fmt.Printf("Listening on http://0.0.0.0:%d\n", port)
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().Int("port", 8080, "Port to listen to clients connections")
}
