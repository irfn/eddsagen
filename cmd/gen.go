package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	eddsa "golang.org/x/crypto/ed25519"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Key pair for EdDSA - ed25519",
	Long:  `hex encoded key pair using EdDSA ed25519 alg.`,
	Run: func(cmd *cobra.Command, args []string) {
		a, b, _ := eddsa.GenerateKey(nil)
		fmt.Println("Public: ", hex.EncodeToString(a))
		fmt.Println("Private: ", hex.EncodeToString(b))
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
