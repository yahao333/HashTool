package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yahao333/hashtool/internal/hash"
)

var (
	algorithm string
	text      string
)

var cliCmd = &cobra.Command{
	Use:   "hash [text]",
	Short: "Generate hash for given text",
	Long:  `Generate hash for given text using specified algorithm.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text = args[0]
		result, err := hash.GenerateHash(text, strings.ToLower(algorithm))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "sha1", "Hash algorithm (md5, sha1, sha256, sha512)")
}