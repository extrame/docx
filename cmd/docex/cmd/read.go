/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	docx "github.com/extrame/docx"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a docx file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		doc, err := docx.Open(args[0])
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		for _, para := range doc.Paragraphs() {

			style := para.GetStyle()

			if style != nil {
				fmt.Printf("Paragraph with the style ->%v(%d)\n", style.Name, style.HeadingLevel())
			} else {
				fmt.Printf("Paragraph without style\n")
			}

			fmt.Printf("\t text ->%s\n", para.Text())
		}
		fmt.Println("End of main")
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
