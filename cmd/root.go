package cmd

import (
	"broadcast-server/internal/server"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	START    bool
	CONNECT  bool
	PORT     string
	USERNAME string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "broadcast-server",
	Short: "Mini CLI chat application",
	Long:  `A CLI chat application, easily broadcast messages messages to all the clients that are connected to the server. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if PORT == "" {
			fmt.Println("port must be set")
			return
		}

		if START {
			server.StartServer(PORT)
			return
		}

		if CONNECT {
			if USERNAME == "" {
				fmt.Println("username must be set for connecting to the server")
				return
			}
			server.StartClient(PORT, USERNAME)
			return
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.broadcast-server.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&START, "start", "s", false, "Start the websocket server")
	rootCmd.PersistentFlags().BoolVarP(&CONNECT, "connect", "c", false, "Connect to the websocket server")
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "p", "", "Port of the websocket server")
	rootCmd.PersistentFlags().StringVarP(&USERNAME, "username", "u", "", "Username of the websocket client")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
