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
		go server(fmt.Sprintf(":%d", p))
		// start chromedp

		// create context
		ctx, cancel := chromedp.NewContext(
			context.Background(), chromedp.WithDebugf(log.Printf))
		defer cancel()

		// run task list
		err = chromedp.Run(ctx, grade(fmt.Sprintf("http://localhost:%d/", p), &config))
		//if err = chromedp.Run(ctx,
		//	chromedp.Navigate(fmt.Sprintf("http://localhost:%d/", p))); err != nil {
		//	log.Fatal(err)
		//}
		//if err = chromedp.Run(ctx, chromedp.WaitVisible("#Title", chromedp.ByID)); err != nil {
		//	log.Fatal(err)
		//}
		//if err = chromedp.Run(ctx, chromedp.Title(&title)); err != nil {
		//	log.Fatal(err)
		}
		// loop through config.Requirments and select
		//fmt.Println(args)
		//fmt.Println(config)
	},
}

func grade(host string, config *models.Config) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible("#Title", chromedp.ByID),
		chromedp.Title(),
	}
}

func server(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(res, indexHTML)
	})
	return http.ListenAndServe(addr, mux)
}

func init() {
	RootCmd.AddCommand(gradeCmd)
	gradeCmd.Flags().IntP("port", "p", 8000, "Port to run web server used for grading")
	gradeCmd.Flags().StringP("file", "f", "", "JSON file assignment")
	gradeCmd.Flags().StringP("assignment", "a", "", "Assignment to grade")
}

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>MIS497PA1</title>
</head>
<body>
    <center>
    <div>
        <h1 id="Title">The Prestige <small id="ReleaseYear">(2006)</small></h1>
        <img src="https://m.media-amazon.com/images/M/MV5BMjA4NDI0MTIxNF5BMl5BanBnXkFtZTYwNTM0MzY2._V1_UX182_CR0,0,182,268_AL_.jpg" alt="Cover" id="CoverArt" />
        <h5 id="Rating"> 8.5/10 </h5>
        <a id="IMDb" href="https://www.imdb.com/title/tt0482571/">IMDb</a>
        <br>
        <iframe width="624" height="480" id="Trailer" src="https://www.youtube.com/embed/JucYLWfFiAs"  allowfullscreen></iframe>
        <p id="Summary">After a tragic accident, two stage magicians engage in a battle to create the ultimate illusion while sacrificing everything they have to outwit each other. </p>
        <ul id="Actors">
            <li>Hugh Jackman</li>
            <li>Christian Bale</li>
            <li>Michael Caine</li>
            <li>Scarlett Johansson</li>
        </ul>
    </div>
    </center>
</body>
</html>
`
