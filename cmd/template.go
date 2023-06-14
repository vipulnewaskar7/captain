package cmd

import (
	"captain/cmd/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "template",
	Long:  `Render docker compose file on stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		values, _ := cmd.Flags().GetString("values")
		template, _ := cmd.Flags().GetString("template")
		var r util.Renderable
		r.ComposeFileTemplate = template
		currentMap := map[string]interface{}{}
		content, _ := os.ReadFile(values)
		if err := yaml.Unmarshal(content, &currentMap); err != nil {
			fmt.Println(err)
		}

		r.Values = currentMap
		r.Render()

	},
}

func init() {
	templateCmd.PersistentFlags().String("values", "", "Values File")
	templateCmd.PersistentFlags().String("template", "", "Templates File")

	rootCmd.AddCommand(templateCmd)
}
