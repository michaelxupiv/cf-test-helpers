package internal

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	"github.com/onsi/gomega/gexec"
)

type starter interface {
	Start(runner.Reporter, string, ...string) (*gexec.Session, error)
}
