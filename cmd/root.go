/*
Copyright © 2020 Maksym Postument 777rip777@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mpostument/grafana-sync/pkg"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grafana-sync",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey, _ := cmd.Flags().GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		tag, _ := cmd.Flags().GetString("tag")
		pkg.PullDashboard(url, apiKey, directory, tag)
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey, _ := cmd.Flags().GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		pkg.PushDashboard(url, apiKey, directory)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grafana-sync.yaml)")
	rootCmd.PersistentFlags().StringP("url", "u", "http://localhost:3000", "Grafana Url with port")
	rootCmd.PersistentFlags().StringP("directory", "d", ".", "Directory where to save dashboards")
	rootCmd.PersistentFlags().StringP("apikey", "a", "", "Grafana api key")
	rootCmd.PersistentFlags().StringP("tag", "t", "", "Dashboard tag to read")

	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(pushCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".grafana-sync" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".grafana-sync")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
