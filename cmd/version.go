// Copyright (c) 2019 KIDTSUNAMI
// Author: alex@kidtsunami.com

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	COMPANY_NAME         = "KIDTSUNAMI UG"
	APP_NAMESPACE        = "cc.blockwatch"
	ORG_NAME             = "Blockwatch"
	APP_NAME             = "tzindex"
	API_VERSION          = "v005-2019-12-05"
	VERSION       string = "v5.2.0"
	GITCOMMIT     string = "dev"
	ENV_PREFIX           = "TZ"
)

var (
	// UserAgent is the string sent as user agent in http headers
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:60.0) Gecko/20100101 Firefox/60.0"
)

func Ua() string {
	return fmt.Sprintf("%s-%s/%s.%s",
		ORG_NAME,
		APP_NAME,
		VERSION,
		GITCOMMIT,
	)
}

func init() {
	UserAgent = Ua()
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + APP_NAME,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s tzindex %s -- %s\n", ORG_NAME, VERSION, GITCOMMIT)
		fmt.Printf("(c) Copyright 2018-2019 -- %s\n", COMPANY_NAME)
		fmt.Printf("Go version (client): %s\n", runtime.Version())
	},
}
