package cmd

import (
	"github.com/spf13/cobra"
	_init "github.com/supabase/cli/internal/init"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		_init.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
