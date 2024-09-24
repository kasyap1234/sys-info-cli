package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var diskCmd = &cobra.Command{
    Use:   "disk",
    Short: "Get disk statistics",
    Run: func(cmd *cobra.Command, args []string) {
       diskStats();
    },
}

func diskStats(){
	fmt.Println("Disk Statistics")
}
func init() {
    rootCmd.AddCommand(diskCmd)
}


