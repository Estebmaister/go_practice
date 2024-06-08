package cmd

import (
	"battle-of-monsters/app/db"
	"log"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command.
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations of database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "up":
				log.Println("migration upping...")
				db.Up()
			case "down":
				log.Println("migration downing...")
				db.Down()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
