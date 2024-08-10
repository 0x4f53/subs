package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0x4f53/textsubs"
	"github.com/spf13/cobra"
)

var (
	domains     bool
	pair        bool
	keepDomains bool
	breakFused  bool
	resolve     bool
	output      []string
	input       string
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
	rootCmd.Flags().BoolVarP(&keepDomains, "keepDomains", "k", false, "Treat subdomains and only domains the same way")
	rootCmd.Flags().BoolVarP(&breakFused, "break", "b", false, "Attempt to break fused domains and subdomains (e.g.: 0x4f.inwwwapple.com becomes 0x4f.in www.apple.com)")
	rootCmd.Flags().BoolVarP(&pair, "pair", "p", false, "Get pairs as json output in the form of {subdomain:\"subdomain.example.com\", domain:\"example.com\"}")
	rootCmd.Flags().BoolVarP(&resolve, "resolve", "r", false, "Only get items that resolve (using local DNS settings)")
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
		output, _ = textsubs.DomainsOnly(string(file), breakFused)

		if resolve {
			output = textsubs.Resolve(output)
		}

	} else if pair {
		pairs, _ := textsubs.SubdomainAndDomainPair(string(file), keepDomains, breakFused)

		if resolve {

			var subdomainsSlice []string
			for _, item := range pairs {
				subdomainsSlice = append(subdomainsSlice, item.Subdomain)
			}

			subdomainsSlice = textsubs.Resolve(subdomainsSlice)

			for _, item := range pairs {
				for _, resolvedSubdomain := range subdomainsSlice {
					if item.Subdomain == resolvedSubdomain {
						jsonBytes, _ := json.Marshal(item)
						output = append(output, string(jsonBytes))
					}
				}
			}

		} else {

			for _, item := range pairs {
				jsonBytes, _ := json.Marshal(item)
				output = append(output, string(jsonBytes))
			}

		}

	} else {
		output, _ = textsubs.SubdomainsOnly(string(file), breakFused)

		if resolve {
			output = textsubs.Resolve(output)
		}
	}

	if len(output) > 0 {
		for _, item := range output {
			fmt.Fprintln(os.Stdout, item)
		}
		os.Exit(0)
	}
}
