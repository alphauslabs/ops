package params

const (
	HostDev  = "opsd-dev-cnugyv5cta-an.a.run.app"
	HostNext = "opsd-next-u554nqhjka-an.a.run.app"
	HostProd = "opsd-prod-u554nqhjka-an.a.run.app"
)

var (
	CredentialsFile string // service acct file for GCP access
	AccessToken     string // auto-set, use as Bearer in subcommands
	RunEnv          string // dev, next, prod (default)
	ServiceHost     string // auto-set
	Bare            bool   // minimal log output, good for jq
)
