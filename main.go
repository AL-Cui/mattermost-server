// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"os"

	"mattermost-server/cmd"
	_ "mattermost-server/cmd/commands"

	// Plugins
	_ "mattermost-server/model/gitlab"

	// Enterprise Imports
	_ "mattermost-server/imports"

	// Enterprise Deps
	_ "github.com/dgryski/dgoogauth"
	_ "github.com/go-ldap/ldap"
	_ "github.com/hashicorp/memberlist"
	_ "github.com/mattermost/rsc/qr"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/tylerb/graceful"
	_ "gopkg.in/olivere/elastic.v5"

	// Temp imports for new dependencies
	_ "github.com/gorilla/schema"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
