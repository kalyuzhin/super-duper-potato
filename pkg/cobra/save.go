/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewSaveCmd – ...
func NewSaveCmd(app App) *cobra.Command {
	var (
		userID  int64
		service string
	)

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save data to vault",
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

			login, err := read("Login: ")
			if err != nil {
				return err
			}

			password, err := readSecret("Password: ")
			if err != nil {
				return err
			}

			err = app.SaveNewPassword(cmd.Context(), userID, masterPassword, service, login, password)
			if err != nil {
				return err
			}

			fmt.Println("credentials has been saved")

			return nil
		},
	}

	cmd.Flags().Int64VarP(&userID, "user", "U", 0, "User ID")
	cmd.Flags().StringVarP(&service, "service", "S", "", "Service name")

	cmd.MarkFlagRequired("user")

	return cmd
}
