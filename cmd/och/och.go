package main

import (
	"fmt"
	"github.com/tosmi/openshift-cluster-health/pkg/version"
)

func main() {
	fmt.Println("Hello World! This is version", version.VERSION, version.GITCOMMIT)
}
