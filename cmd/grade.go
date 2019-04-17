package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"majora/models"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

var gradeCmd = &cobra.Command{
	Use:   "grade",
	Short: "Grade an assignment",
	Run: func(cmd *cobra.Command, args []string) {
		// read in json config
		f, _ := cmd.Flags().GetString("file")
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		var config models.Config
		err = json.Unmarshal(data, &config)
		if err != nil {
			log.Fatal(err)
		}

		// start server
		server(f)

		// start chromedp

		// create context
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// run task list
		err = chromedp.Run(ctx, grade("http://localhost:3000", &config))
		if err != nil {
			log.Fatal(err)
		}
		// loop through config.Requirments and select
		fmt.Println(args)
	},
}

func grade(host string, config *models.Config) chromedp.Tasks {
	return chromedp.Tasks{}
}

func server(file string) {
	http.Handle("/", http.FileServer(http.Dir(file)))
	http.ListenAndServe(":3000", nil)
}

func init() {
	RootCmd.AddCommand(gradeCmd)
	gradeCmd.Flags().StringP("file", "f", "", "JSON file assignment")
	gradeCmd.Flags().StringP("assignment", "a", "", "Assignment to grade")
}
