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
		p, _ := cmd.Flags().GetInt("port")
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
		fmt.Println(config)

		// start server
		go server(f, p)

		// start chromedp

		// create context
		ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
		defer cancel()

		// run task list
		err = chromedp.Run(ctx, grade(fmt.Sprintf("http://localhost:%d", p), &config))
		if err != nil {
			log.Fatal(err)
		}
		// loop through config.Requirments and select
		fmt.Println(args)
		fmt.Println(config)
	},
}

func grade(host string, config *models.Config) chromedp.Tasks {
	actions := make([]chromedp.Action, len(config.Requirements))
	actions = append(actions, chromedp.Navigate(host))

	//for r, i := range config.Requirements {

	//}
	return actions
}

func server(file string, port int) {
	http.Handle("/", http.FileServer(http.Dir(file)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func init() {
	RootCmd.AddCommand(gradeCmd)
	gradeCmd.Flags().IntP("port", "p", 8000, "Port to run web server used for grading")
	gradeCmd.Flags().StringP("file", "f", "", "JSON file assignment")
	gradeCmd.Flags().StringP("assignment", "a", "", "Assignment to grade")
}
