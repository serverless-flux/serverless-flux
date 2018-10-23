package release

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/pkg/errors"

	sls "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"
	"github.com/serverless-operator/serverless-operator/pkg/config"
	"github.com/serverless-operator/serverless-operator/pkg/env"
	// nolint:lll
)

type Release struct {
	config *config.Config
}

func New(config *config.Config) *Release {
	return &Release{
		config: config,
	}
}

func IsReleaseComplete(status string, defaultRet bool) bool {
	switch status {
	case "DEPLOYED":
		return true
	case "DELETE_COMPLETE":
		return false
	}
	return false
}

// GetReleaseName either retrieves the release name from the Custom Resource or constructs a new one
//  in the form : $Namespace-$CustomResourceName
func GetReleaseName(slr sls.ServerlessRelease) string {
	namespace := slr.Namespace
	if namespace == "" {
		namespace = "default"
	}
	releaseName := slr.Spec.ReleaseName
	if releaseName == "" {
		releaseName = fmt.Sprintf("%s-%s", namespace, slr.Name)
	}

	return releaseName
}

func (r *Release) Install(releaseName string, slr sls.ServerlessRelease) error {
	logger := r.config.Logger
	logger.Infof("install releaseName= %s", releaseName)

	namespace := slr.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}

	args := []string{r.config.Depenendencies.ServerlessPath, "deploy"}
	args = append(args, "--stage", slr.Spec.Stage)
	args = append(args, "--region", slr.Spec.Region)
	args = append(args, "--package", slr.Spec.PackagePath)
	if slr.Spec.Verbose {
		args = append(args, "--verbose")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, r.config.Depenendencies.NodePath, args...)

	cmd.Env = []string{}
	for _, envVar := range slr.Spec.Env {
		if envVar.ValueFrom == nil {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", envVar.Name, envVar.Value))
			continue
		}

		value, err := env.GetEnvVarValue(r.config.KubeClientset, namespace, envVar.ValueFrom)
		if err != nil {
			return errors.Wrapf(err, "error setting environmant variable %s for release %s", envVar.Name, releaseName)
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", envVar.Name, value))
	}

	output, err := cmd.CombinedOutput()
	logger.Infof("serverless output: %s\n", string(output))
	if err != nil {
		logger.Infof("serverless err: %v", err)
	}
	return err

}
