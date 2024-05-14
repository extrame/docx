/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	docx "github.com/extrame/docx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a docx file, show as text",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var prefix = cmd.Flag("prefix").Value.String()

		if len(args) == 0 {
			cmd.PrintErr("Please provide a docx file")
			return
		}

		var fileName = filepath.Base(args[0])

		var prefixes = []string{fileName}

		doc, err := docx.Open(args[0])
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		var splitLevel int
		var currSplitLevel int
		for _, para := range doc.Paragraphs() {
			style := para.GetStyle()

			if prefix == "h" {

				if style != nil {
					if len(prefixes) > splitLevel {
						prefixes = prefixes[:splitLevel]
					} else if len(prefixes) < splitLevel {
						for i := len(prefixes); i < splitLevel; i++ {
							prefixes = append(prefixes, "")
						}
					}
				}
				prefixes = append(prefixes, para.Text())
			} else if prefix == "o" {
				currSplitLevel = para.GetOutlineLevel()
				if currSplitLevel > -1 {
					if splitLevel < currSplitLevel {
						for i := len(prefixes); i < splitLevel; i++ {
							prefixes = append(prefixes, "")
						}
						prefixes = append(prefixes, para.Text())
					} else if splitLevel == currSplitLevel && len(prefixes) > 0 {
						for i := len(prefixes); i < splitLevel+2; i++ {
							prefixes = append(prefixes, "")
						}
						prefixes[splitLevel+1] = para.Text()
					} else {
						prefixes[currSplitLevel+1] = para.Text()
						prefixes = prefixes[:currSplitLevel+2]
					}
					logrus.Info("Outline level: ", prefixes)
					splitLevel = currSplitLevel
				}
			}

			if prefix == "h" {
				if style == nil {
					fmt.Printf("%s: %s\n", prefixes, para.Text())
				}
			} else if prefix == "o" {
				if currSplitLevel < 0 {
					fmt.Printf("%s: %s\n", strings.Join(prefixes, ">"), para.Text())
				}
			} else {
				fmt.Printf("%s\n", para.Text())
			}
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
	readCmd.Flags().StringP("prefix", "p", "", "add a prefix to each paragraph(headlevels[h]/outlinelevels[o])")
}
