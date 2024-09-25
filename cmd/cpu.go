package cmd

import (
	"fmt"
	

	"github.com/fatih/color"

	"github.com/prometheus/procfs"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)
func procsFunc() {
    procs, err := procfs.Self()
    if err != nil {
        color.Red("Error getting process information:", err)

        return
    }
    stat, err := procs.Stat()
    if err != nil {
       color.Red("Error getting process information:", err)
        return
    }
    color.Cyan("cpu time: %.2f\n", stat.CPUTime())
}
var procsCpuCmd = &cobra.Command{
    Use:   "procs",
    Short: "Get CPU statistics",
    Run: func(cmd *cobra.Command, args []string) {
        procsFunc()
    },
}
var cpuCmd = &cobra.Command{
    Use:   "cpu",
    Short: "Get CPU statistics",
    Run: func(cmd *cobra.Command, args []string) {
        CpuStats()
    },
}

func CpuStats() {
    // Get CPU usage percentage
    
    cpuUsage, err := cpu.Percent(10, false)
    if err != nil {
        color.Red("Error getting CPU usage:", err)
        return
    }

    // Print CPU usage
    color.Cyan("CPU Usage: %.2f%%\n", cpuUsage[0])

    // Get CPU information
    cpuInformation, err := cpu.Info()
    if err != nil {
        color.Red("Error getting CPU information:", err)
        return
    }

    // Print CPU information
    color.HiWhiteString("CPU Information:")
    for _, info := range cpuInformation {
        color.Cyan("  CPU %d:", info.CPU)
        color.Cyan("    Vendor ID: %s\n", info.VendorID)
        color.Cyan("  Vendor ID: %s\n", info.VendorID)
        color.Cyan("  CPU Family: %s\n", info.Family)
        color.Cyan("  CPU Model: %s\n", info.Model)
        color.Cyan("  CPU Stepping: %d\n", info.Stepping)
        color.Cyan("  Physical ID: %s\n", info.PhysicalID)
        color.Cyan("  Core ID: %s\n", info.CoreID)
        color.Cyan("  Cores: %d\n", info.Cores)
        color.Cyan("  Max MHz: %.2f MHz\n", info.Mhz)
        color.Cyan("  Cache Size: %d KB\n", info.CacheSize)
        fmt.Println("  Flags:", info.Flags)
        fmt.Println()
    }
   
    fmt.Println("----------------------------------");

}


func init() {
    rootCmd.AddCommand(cpuCmd)
    rootCmd.AddCommand(procsCpuCmd)

}
