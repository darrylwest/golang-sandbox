package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    var root = &cobra.Command{
        Use:   "cobra-cli [OPTIONS]",
        Short: "My Cobra CLI test...",
        Long: "this is a verbose explaination of what this thing does...",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("root command with args: %v\n", args[1:])
        },
    }

    root.Flags().Bool("version", false, "Shot the version number")

    root.SetArgs(os.Args)
    root.Execute()
}
