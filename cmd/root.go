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
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/mpostument/grafana-sync/grafana"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var customHeaders map[string]string

var rootCmd = &cobra.Command{
	Use:     "grafana-sync",
	Short:   "Root command for grafana interaction",
	Long:    `Root command for grafana interaction.`,
	Version: "1.5.0",
}

var pullDashboardsCmd = &cobra.Command{
	Use:   "pull-dashboards",
	Short: "Pull grafana dashboards in to the directory",
	Long: `Save to the directory grafana dashboards.
Directory name specified by flag --directory. If flag --tag is used,
only dashboards with given tag are pulled`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			folderId int
			err      error
		)
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		tag, _ := cmd.Flags().GetString("tag")
		folderName, _ := cmd.Flags().GetString("folderName")

		if folderName != "" {
			folderId, err = grafana.FindFolderId(url, apiKey, folderName)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			folderId, _ = cmd.Flags().GetInt("folderId")
		}

		if err := grafana.PullDashboard(url, apiKey, directory, tag, folderId); err != nil {
			log.Fatalln("Pull dashboards command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pushDashboardsCmd = &cobra.Command{
	Use:   "push-dashboards",
	Short: "Push grafana dashboards from directory",
	Long:  `Read json with dashboards description and publish to grafana.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			folderId int
			err      error
		)
		url, _ := cmd.Flags().GetString("url")
		apiKey, _ := cmd.Flags().GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		folderName, _ := cmd.Flags().GetString("folderName")

		if folderName != "" {
			folderId, err = grafana.FindFolderId(url, apiKey, folderName)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			folderId, _ = cmd.Flags().GetInt("folderId")
		}

		if err := grafana.PushDashboard(url, apiKey, directory, folderId); err != nil {
			log.Fatalln("Push dashboards command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pullFoldersCmd = &cobra.Command{
	Use:   "pull-folders",
	Short: "Pull grafana folders json in to the directory",
	Long: `Save to the directory grafana folders json.
Directory name specified by flag --directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PullFolders(url, apiKey, directory); err != nil {
			log.Fatalln("Pull folders command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pushFoldersCmd = &cobra.Command{
	Use:   "push-folders",
	Short: "Read json and create grafana folders",
	Long:  `Read json with folders description and publish to grafana.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PushFolder(url, apiKey, directory); err != nil {
			log.Fatalln("Push folders command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pullNotificationsCmd = &cobra.Command{
	Use:   "pull-notifications",
	Short: "Pull grafana notifications json in to the directory",
	Long: `Save to the directory grafana folders json.
Directory name specified by flag --directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PullNotifications(url, apiKey, directory); err != nil {
			log.Fatalln("Pull notifications command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pushNotificationsCmd = &cobra.Command{
	Use:   "push-notifications",
	Short: "Read json and create grafana notifications",
	Long:  `Read json with notifications description and publish to grafana.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PushNotification(url, apiKey, directory); err != nil {
			log.Fatalln("Push notifications command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pullDataSourcesCmd = &cobra.Command{
	Use:   "pull-datasources",
	Short: "Pull grafana datasources json in to the directory",
	Long: `Save to the directory grafana datasources json.
Directory name specified by flag --directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PullDatasources(url, apiKey, directory); err != nil {
			log.Fatalln("Pull datasources command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
	},
}

var pushDataSourcesCmd = &cobra.Command{
	Use:   "push-datasources",
	Short: "Read json and create grafana datasources",
	Long:  `Read json with datasources description and publish to grafana.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		apiKey := viper.GetString("apikey")
		directory, _ := cmd.Flags().GetString("directory")
		if err := grafana.PushDatasources(url, apiKey, directory); err != nil {
			log.Fatalln("Push datasources command failed", err)
		}
		if grafana.ExecutionErrorHappened {
			os.Exit(1)
		}
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
	rootCmd.PersistentFlags().StringToStringVar(&customHeaders, "customHeaders", map[string]string{}, "Key-value pairs of custom http headers (key1=value1,key2=value2)")
	pullDataSourcesCmd.PersistentFlags().StringP("tag", "t", "", "Dashboard tag to read")
	pushDashboardsCmd.PersistentFlags().IntP("folderId", "f", 0, "Directory Id to which push dashboards")
	pushDashboardsCmd.PersistentFlags().StringP("folderName", "n", "", "Directory name to which push dashboards")
	pullDashboardsCmd.PersistentFlags().IntP("folderId", "f", -1, "Directory Id from which pull dashboards")
	pullDashboardsCmd.PersistentFlags().StringP("folderName", "n", "", "Directory name from which pull dashboards")
	pullDashboardsCmd.PersistentFlags().StringP("tag", "t", "", "Dashboard tag to p")

	if err := viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey")); err != nil {
		log.Println(err)
	}

	if err := viper.BindPFlag("customHeaders", rootCmd.PersistentFlags().Lookup("customHeaders")); err != nil {
		log.Println(err)
	}

	rootCmd.AddCommand(pullDashboardsCmd)
	rootCmd.AddCommand(pushDashboardsCmd)
	rootCmd.AddCommand(pullFoldersCmd)
	rootCmd.AddCommand(pushFoldersCmd)
	rootCmd.AddCommand(pullNotificationsCmd)
	rootCmd.AddCommand(pushNotificationsCmd)
	rootCmd.AddCommand(pullDataSourcesCmd)
	rootCmd.AddCommand(pushDataSourcesCmd)
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

	grafana.InitHttpClient(viper.GetStringMapString("customHeaders"))
}
