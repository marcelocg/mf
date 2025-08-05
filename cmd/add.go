package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mf/internal/storage"
	"mf/internal/totp"
	"mf/internal/types"
)

var addCmd = &cobra.Command{
	Use:   "add [ACCOUNT_NAME] [SECRET]",
	Short: "Adiciona uma nova conta para geração de tokens TOTP",
	Long:  `Adiciona uma nova conta com o nome especificado e o secret fornecido para geração de tokens TOTP.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		accountName := args[0]
		secret := args[1]

		if err := totp.ValidateSecret(secret); err != nil {
			return fmt.Errorf("secret inválido: %w", err)
		}

		store, err := storage.NewSecure()
		if err != nil {
			return fmt.Errorf("erro ao inicializar storage: %w", err)
		}

		account := types.Account{
			Name:   accountName,
			Secret: secret,
		}

		if err := store.SaveAccount(account); err != nil {
			return fmt.Errorf("erro ao salvar conta: %w", err)
		}

		fmt.Printf("Conta '%s' adicionada com sucesso.\n", accountName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
