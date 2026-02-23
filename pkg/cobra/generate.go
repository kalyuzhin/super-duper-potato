/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewGenerateCmd – ...
func NewGenerateCmd(app App) *cobra.Command {
	var (
		length uint8
	)

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate new secure password",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			password, err := app.GenerateNewSecurePassword(cmd.Context(), length)
			if err != nil {
				return err
			}

			fmt.Printf("Generated password: %s\n\r", password)

			return nil
		},
	}

	cmd.Flags().Uint8VarP(&length, "length", "L", 0, "Password length")

	cmd.MarkFlagRequired("length")

	return cmd
}
