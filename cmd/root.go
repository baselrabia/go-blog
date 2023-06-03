package cmd
import (
	"os"
	"fmt"
  
	"github.com/spf13/cobra"
  )

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Help command ",
	Long: `Display help command from the cmd`, 	
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }