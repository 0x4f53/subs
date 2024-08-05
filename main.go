package main

import (
	"fmt"
	"os"

	"github.com/0x4f53/textsubs"
	"github.com/spf13/cobra"
)

var (
	domains bool
	unique  bool
	output  []string
	input   string
)

var rootCmd = &cobra.Command{
	Use:   "subs [input_file]",
	Short: "subs",
	Long:  "Grab valid subdomains from files!\n(Visit https://github.com/0x4f53/subs for more details)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input = args[0]
	},
}

func main() {
	rootCmd.Flags().BoolVarP(&domains, "domains", "d", false, "Get domains only")
	rootCmd.Flags().BoolVarP(&unique, "unique", "u", false, "Only get unique entries")
	rootCmd.Flags().BoolP("help", "h", false, "Help")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not read "+input)
		os.Exit(-1)
	}

	if domains && unique {
		output = textsubs.DomainsOnly(string(file), true)
	} else if domains && !unique {
		output = textsubs.DomainsOnly(string(file), false)
	} else if !unique {
		output = textsubs.SubdomainsOnly(string(file), false)
	} else {
		output = textsubs.SubdomainsOnly(string(file), true)
	}

	if len(output) > 0 {
		for _, item := range output {
			fmt.Fprintln(os.Stdout, item)
		}
		os.Exit(0)
	}
}
