package cmd

import (
	"autoOps/cmd/migrate"
	"autoOps/cmd/server"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "me",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `me`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 me，可以使用 -h 查看命令`
		fmt.Printf("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(server.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
