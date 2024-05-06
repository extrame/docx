/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	docx "github.com/extrame/docx"
	"github.com/extrame/docx/transformer/pdf"
	"github.com/spf13/cobra"
)

var output struct {
	Format         string
	Path           string
	FontSearchPath string
}

// transCmd represents the trans command
var transCmd = &cobra.Command{
	Use:   "trans",
	Short: "transformate docx file to another format (pdf until now)",
	Long:  `Please set font directory to use this command if you need transformate UTF-8 document to pdf format.`,
	Run: func(cmd *cobra.Command, args []string) {
		doc, err := docx.Open(args[0])
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		if output.Path == "" {
			output.Path = args[0] + "." + output.Format
		}
		pdf.Trans(args[0], doc, output.Path, output.FontSearchPath)
	},
}

func init() {
	rootCmd.AddCommand(transCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	transCmd.PersistentFlags().StringVarP(&output.Format, "output", "o", "pdf", "output format (pdf or html)")
	transCmd.PersistentFlags().StringVarP(&output.Path, "path", "p", "", "output path")
	transCmd.PersistentFlags().StringVarP(&output.FontSearchPath, "font", "f", "", "font search path(directory)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
