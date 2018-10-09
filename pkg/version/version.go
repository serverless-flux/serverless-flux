package version

// Version is the version number of the initializer
var Version string

// BuildDate is the date the application was built
var BuildDate string

// GitCommit is the Git commit hash that was used when building the release
var GitCommit string

func init() {
	Version = "dev"
	GitCommit = "none"
	BuildDate = "uknown"
}
