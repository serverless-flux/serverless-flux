package operator

import (
	"github.com/serverless-operator/serverless-operator/pkg/config"
)

// Operator implements the serverless releaser using
// the operator pattern
type Operator struct {
	Config *config.Config
}
