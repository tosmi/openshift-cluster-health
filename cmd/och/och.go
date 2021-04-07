package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tosmi/openshift-cluster-health/pkg/version"
)

var (
	rootCmd = &cobra.Command{
		Use:   "och",
		Short: "OpenShift Cluster Healthcheck",
		Long:  `och runs various check againts a running OpenShift Cluster. `,
	}
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	fmt.Println("Hello World! This is version", version.VERSION, version.GITCOMMIT)
	log.Debug("test")
	rootCmd.Execute()
}
