package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mf/internal/storage"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as contas disponíveis",
	Long:  `Lista todas as contas disponíveis para geração de tokens TOTP.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := storage.NewSecure()
		if err != nil {
			return fmt.Errorf("erro ao inicializar storage: %w", err)
		}

		accounts, err := store.ListAccounts()
		if err != nil {
			return fmt.Errorf("erro ao listar contas: %w", err)
		}

		if len(accounts) == 0 {
			fmt.Println("Nenhuma conta encontrada.")
			return nil
		}

		fmt.Println("Contas disponíveis:")
		for _, account := range accounts {
			fmt.Printf("  %s\n", account)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
