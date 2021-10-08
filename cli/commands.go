package cli

import (
	"fmt"
	"github.com/mrNobody95/Gate/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gate",
	Short: "Golang algo trading engine",
	Long:  `Gate is an open source algo trading engine written by Golang.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gate",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gate - Golang Algo Trading Engine V0.0.1")
	},
}

func initStartEngineCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "start",
		Short: "start engine with given node config",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().NFlag() == 0 {
				cmd.Help()
			} else {
				brokerage, _ := cmd.Flags().GetString("brokerage")
				configPath, _ := cmd.Flags().GetString("config")
				core.StartNewNode(brokerage, configPath)
			}
		},
	}
	command.Flags().StringP("brokerage", "b", "nobitex", "Select brokerage node [nobitex] (default \"nobitex\")")
	command.Flags().StringP("config", "c", "", "Node config json file path - let it empty for using default config")
	return command
}
