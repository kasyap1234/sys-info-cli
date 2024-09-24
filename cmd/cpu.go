package cmd

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/spf13/cobra"
)

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
        fmt.Println("Error getting CPU usage:", err)
        return
    }

    // Print CPU usage
    fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage[0])

    // Get CPU information
    cpuInformation, err := cpu.Info()
    if err != nil {
        fmt.Println("Error getting CPU information:", err)
        return
    }

    // Print CPU information
    fmt.Println("CPU Information:")
    for _, info := range cpuInformation {
        fmt.Printf("  Model: %s\n", info.ModelName)
        fmt.Printf("  Vendor ID: %s\n", info.VendorID)
        fmt.Printf("  CPU Family: %s\n", info.Family)
        fmt.Printf("  CPU Model: %s\n", info.Model)
        fmt.Printf("  CPU Stepping: %d\n", info.Stepping)
        fmt.Printf("  Physical ID: %s\n", info.PhysicalID)
        fmt.Printf("  Core ID: %s\n", info.CoreID)
        fmt.Printf("  Cores: %d\n", info.Cores)
        fmt.Printf("  Max MHz: %.2f MHz\n", info.Mhz)
        fmt.Printf("  Cache Size: %d KB\n", info.CacheSize)
        fmt.Println("  Flags:", info.Flags)
        fmt.Println()
    }
}

func init() {
    rootCmd.AddCommand(cpuCmd)
}
