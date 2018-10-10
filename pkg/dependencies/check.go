package dependencies

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/blang/semver"
	"github.com/pkg/errors"

	"github.com/serverless-operator/serverless-operator/pkg/config"
)

const (
	// MinSlsVersion defines the minimum version of the Serverless
	// framework that we need
	MinSlsVersion = "1.28.0"

	// MinNodeVersion defines the minimum version of the Serverless
	// framework that we need
	MinNodeVersion = "6.5.0"
)

// Check will check if the required dependencies are installed and
// have the required versions
func Check(config *config.Config) error {

	if err := checkNode(config); err != nil {
		return errors.Wrap(err, "checking nodejs dependency")
	}

	if err := checkServerless(config); err != nil {
		return errors.Wrap(err, "checking serverless framework version")
	}

	return nil
}

func checkNode(config *config.Config) error {
	if config.Depenendencies.NodePath == "" {
		nodepath, err := exec.LookPath("node")
		if err != nil {
			return fmt.Errorf("nodejs not found within the PATH. Add it to the path or use the -n flag")
		}
		config.Depenendencies.NodePath = nodepath
	}

	if _, err := os.Stat(config.Depenendencies.NodePath); os.IsNotExist(err) {
		return fmt.Errorf("nodejs doesn't exist at supplied path: %s", config.Depenendencies.NodePath)
	}

	cmd := exec.Command(config.Depenendencies.NodePath, "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "error getting node version")
	}

	config.Logger.Debugf("nodejs version %s is required, checking actual version\n", MinNodeVersion)

	trimmed := strings.TrimSuffix(string(out), "\n")
	actualVersion, err := semver.Parse(strings.TrimLeft(trimmed, "v"))
	if err != nil {
		return errors.Wrapf(err, "parsing actual nodejs version string %q", actualVersion)
	}

	requiredVersion, err := semver.Parse(MinNodeVersion)
	if err != nil {
		return errors.Wrapf(err, "parsing required nodejs version string %q", requiredVersion)
	}

	if actualVersion.LT(requiredVersion) {
		return fmt.Errorf("nodejs version %s was found at %s, minimum required version is %s",
			actualVersion, config.Depenendencies.NodePath, requiredVersion)
	}
	config.Logger.Debugf("nodejs version %s was found at %s\n", actualVersion, config.Depenendencies.NodePath)

	return nil
}

func checkServerless(config *config.Config) error {
	if config.Depenendencies.ServerlessPath == "" {
		slsPath, err := exec.LookPath("serverless")
		if err != nil {
			return fmt.Errorf("serverless framework not found within the PATH. Add it to the path or use the -s flag")
		}
		config.Depenendencies.ServerlessPath = slsPath
	}

	if _, err := os.Stat(config.Depenendencies.ServerlessPath); os.IsNotExist(err) {
		return fmt.Errorf("serverless framework executable doesn't exist at supplied path: %s",
			config.Depenendencies.ServerlessPath)
	}

	cmd := exec.Command(config.Depenendencies.NodePath, config.Depenendencies.ServerlessPath, "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "error getting serverless framework version")
	}

	config.Logger.Debugf("serverless framework version %s is required, checking actual version\n", MinSlsVersion)

	trimmed := strings.TrimSuffix(string(out), "\n")
	actualVersion, err := semver.Parse(trimmed)
	if err != nil {
		return errors.Wrapf(err, "parsing actual serverless framework version string %q", actualVersion)
	}

	requiredVersion, err := semver.Parse(MinSlsVersion)
	if err != nil {
		return errors.Wrapf(err, "parsing required serverless framework version string %q", requiredVersion)
	}

	if actualVersion.LT(requiredVersion) {
		return fmt.Errorf("serverless framework version %s was found at %s, minimum required version is %s",
			actualVersion, config.Depenendencies.ServerlessPath, requiredVersion)
	}
	config.Logger.Debugf("serverless framework version %s was found at %s\n",
		actualVersion, config.Depenendencies.ServerlessPath)

	return nil
}
