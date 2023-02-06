/*
Copyright © 2023 Red Hat

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mdbooth/openstack-go/httpRequester"
	"github.com/mdbooth/openstack-go/openstack/identity/v3/tokens"
)

// appCredsCmd represents the appCreds command
var appCredsCmd = &cobra.Command{
	Use:   "appCreds",
	Short: "Authenticate using application credentials",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		runAppCreds()
	},
}

var (
	authEndpoint     string
	credentialID     string
	credentialSecret string
)

func init() {
	appCredsCmd.Flags().StringVarP(&authEndpoint, "auth-endpoint", "a", "", "OpenStack authentication endpoint")
	appCredsCmd.MarkFlagRequired("auth-endpoint")

	appCredsCmd.Flags().StringVarP(&credentialID, "credential-id", "i", "", "OpenStack application credential ID")
	appCredsCmd.Flags().StringVarP(&credentialSecret, "credential-secret", "s", "", "OpenStack application credential secret")
	appCredsCmd.MarkFlagsRequiredTogether("credential-id", "credential-secret")
	appCredsCmd.MarkFlagRequired("credential-id")
	appCredsCmd.MarkFlagRequired("credential-secret")

	rootCmd.AddCommand(appCredsCmd)
}

func runAppCreds() {
	credentials := tokens.RequestAuthIdentityApplicationCredentialByID{
		ID:     credentialID,
		Secret: credentialSecret,
	}

	tokenRequest := tokens.Request{
		Auth: tokens.RequestAuth{
			Identity: tokens.RequestAuthIdentity{
				Methods: []tokens.RequestAuthIdentityMethods{
					tokens.RequestAuthIdentityMethodsApplicationCredential,
				},
				ApplicationCredential: credentials.AsRequestAuthIdentityApplicationCredential(),
			},
		},
	}

	ctx := context.TODO()
	httpRequester := httpRequester.NewRequester(authEndpoint)
	response, err := tokenRequest.AuthenticateUser(ctx, httpRequester)
	if err != nil {
		panic(err)
	}
	fmt.Println("response: ", response)
}
