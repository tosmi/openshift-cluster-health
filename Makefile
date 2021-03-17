PROJECT := github.com/tosmi/openshift-cluster-health
ifdef GITCOMMIT
        GITCOMMIT := $(GITCOMMIT)
else
        GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
endif

COMMON_LDFLAGS := -X ${PROJECT}/pkg/version.GITCOMMIT=${GITCOMMIT}
BUILD_FLAGS := -mod=vendor -ldflags="${COMMON_LDFLAGS}"

default: bin

.PHONY:  bin
bin:
	go build ${BUILD_FLAGS} cmd/och/och.go
