package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
} 

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "tables migrations",
	Long: "tables migrations",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Migrate()
	},
}
