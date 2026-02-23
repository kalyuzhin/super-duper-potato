/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRmCmd – ...
func NewRmCmd(app App) *cobra.Command {
	var (
		userID  int64
		service string
	)

	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Delete vault data",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			masterPassword, err := readSecret("Master password: ")
			if err != nil {
				return err
			}

			if service == "" {
				serviceTmp, err := read("Service: ")
				if err != nil {
					return err
				}

				service = serviceTmp
			}

			err = app.DeleteVaultData(cmd.Context(), userID, masterPassword, service)
			if err != nil {
				return err
			}

			fmt.Println("vault data has been deleted")

			return nil
		},
	}

	cmd.Flags().Int64VarP(&userID, "user", "U", 0, "User ID")
	cmd.Flags().StringVarP(&service, "service", "S", "", "Service name")

	cmd.MarkFlagRequired("user")

	return cmd
}
