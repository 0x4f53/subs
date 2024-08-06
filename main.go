package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0x4f53/textsubs"
	"github.com/spf13/cobra"
)

var (
	domains bool
	pair    bool
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
	rootCmd.Flags().BoolVarP(&pair, "pair", "p", false, "Get pairs as json output in the form of {subdomain:\"subdomain.example.com\", domain:\"example.com\"}")
	rootCmd.Flags().BoolP("help", "h", false, "Help")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-2)
	}

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not read \""+input+"\"")
		os.Exit(-1)
	}

	if domains {
		output, _ = textsubs.DomainsOnly(string(file), unique)
	} else if pair {
		pairs, _ := textsubs.SubdomainAndDomainPair(string(file), unique)
		for _, item := range pairs {
			jsonBytes, _ := json.Marshal(item)
			output = append(output, string(jsonBytes))
		}
	} else {
		output, _ = textsubs.SubdomainsOnly(string(file), unique)
	}

	if len(output) > 0 {
		for _, item := range output {
			fmt.Fprintln(os.Stdout, item)
		}
		os.Exit(0)
	}
}
