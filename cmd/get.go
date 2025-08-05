package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mf/internal/storage"
	"mf/internal/totp"
)

var getCmd = &cobra.Command{
	Use:   "get [ACCOUNT_NAME]",
	Short: "Gera um token TOTP para a conta especificada",
	Long:  `Gera um token TOTP (Time-based One-Time Password) para a conta especificada.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		accountName := args[0]

		store, err := storage.NewSecure()
		if err != nil {
			return fmt.Errorf("erro ao inicializar storage: %w", err)
		}

		account, err := store.LoadAccount(accountName)
		if err != nil {
			return fmt.Errorf("erro ao carregar conta: %w", err)
		}

		token, err := totp.GenerateToken(account.Secret)
		if err != nil {
			return fmt.Errorf("erro ao gerar token: %w", err)
		}

		fmt.Println(token)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
