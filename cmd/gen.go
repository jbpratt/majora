package cmd

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"majora/models"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate README assignments based on a JSON config",
	Run: func(cmd *cobra.Command, args []string) {
		// read in JSON
		// Execute into template
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
		t := template.Must(template.ParseFiles("templates/README.md.tmpl"))
		name := strings.TrimSuffix(f, filepath.Ext(f))
		err = os.Mkdir(name, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		dst, err := os.Create(name + "/README.md")
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()
		err = t.Execute(dst, config)
		if err != nil {
			log.Fatal(err)
		}
		err = dst.Sync()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("file", "f", "", "JSON file assignment")
}
