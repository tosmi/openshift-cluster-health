package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

func main() {
	//fmt.Println("Hello World! This is version", version.VERSION, version.GITCOMMIT)
	//log.Debug("test")
	rootCmd.Execute()
}
