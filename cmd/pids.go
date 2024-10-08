package cmd

import (
	"fmt"
	"strconv"

	
	"github.com/fatih/color"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)
var pidsCmd = &cobra.Command{
    Use:   "List Pids",
    Short: "List all  process IDs running currently on the system",
    Run: func(cmd *cobra.Command, args []string) {
         running()
    },


}
var existsPidCmd = &cobra.Command{
    Use:   "ExistsPid",
    Short: "Check if a process with the given PID exists",
	Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        existsPid(args)
    },
}
func existsPid(args []string){
	pidArg :=args[0]; 
	pid,err :=strconv.Atoi(pidArg); 
	if err !=nil {
color.Red("Invalid PID:", err); 
		return 
	}
	exists,err :=process.PidExists(int32(pid)); 
	if err !=nil {
		color.Red("Error checking PID:", err);
		return
	}
	if exists {
		color.Cyan("Process with PID %d exists.\n", pid)
	} else {
		fmt.Printf("Process with PID %d does not exist.\n", pid)
	}

}


func running (){

	runningPids ,err :=process.Pids()
	if err != nil {
		color.Red("Error getting running processes:", err)
		return
	}
	color.Yellow("Running Processes : ",runningPids); 
	for _, pid := range runningPids {
		fmt.Println(pid)
	}

}

func init (){
	rootCmd.AddCommand(pidsCmd);
	rootCmd.AddCommand(existsPidCmd);
}
