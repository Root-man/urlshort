/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/root-man/urlshort/server"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "urlshort",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		yamlFilePath, err := cmd.Flags().GetString("yaml-file")
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("yaml-file flag error: %s", err.Error()))
		}

		jsonFilePath, err := cmd.Flags().GetString("json-file")
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("json-file flag error: %s", err.Error()))
		}

		if yamlFilePath != "" && jsonFilePath != "" {
			cobra.CompErrorln("both JSON and YAML file are specified, not sure which to use")
		} else if yamlFilePath != "" {
			server.RunYAML(yamlFilePath)
		} else if jsonFilePath != "" {
			server.RunJSON(jsonFilePath)
		} else {
			cobra.CompErrorln("at least one config file must be specified")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.urlshort.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("yaml-file", "y", "", "Path to YAML file")
	rootCmd.Flags().StringP("json-file", "j", "", "Path to JSON file")
}
