package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/config"
)

// Exit codes
const (
	ExitOK      = 0
	ExitGeneral = 1
	ExitInput   = 2
	ExitAuth    = 3
	ExitNetwork = 4
)

// global flags populated before any subcommand runs
var (
	flagPretty  bool
	flagQuiet   bool
	flagTimeout time.Duration
)

// rootCmd is the top-level command.
var rootCmd = &cobra.Command{
	Use:   "raku-cli",
	Short: "CLI for Rakuten Web Service APIs",
	Long: `raku-cli wraps the Rakuten Web Service API.
Output is always JSON on stdout; status messages go to stderr.`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(ExitGeneral)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&flagPretty, "pretty", false, "Pretty-print JSON output")
	rootCmd.PersistentFlags().BoolVar(&flagQuiet, "quiet", false, "Suppress stderr log output")
	rootCmd.PersistentFlags().DurationVar(&flagTimeout, "timeout", 30*time.Second, "HTTP request timeout")
}

// AddCommand registers a subcommand on the root.
func AddCommand(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}

// LoadRakutenClient loads config and returns a Client for standard APIs (Travel, Books, etc.)
// Uses app_id (numeric). Exits on missing credentials.
func LoadRakutenClient() *client.Client {
	cfg, err := config.Load()
	if err != nil {
		logErr("load config: %v", err)
		os.Exit(ExitGeneral)
	}
	if cfg.Rakuten.AppID == "" {
		logErr("Rakuten App ID not set — use RAKUTEN_APP_ID env var or ~/.config/raku-cli/config.yaml")
		os.Exit(ExitAuth)
	}
	return client.New(cfg.Rakuten.AppID, cfg.Rakuten.AffiliateID, cfg.Rakuten.AccessKey, cfg.Rakuten.Origin, flagTimeout)
}

// LoadIchibaClient loads config and returns a Client for the Ichiba API.
// Prefers ichiba_app_id (UUID) if set, falls back to app_id.
func LoadIchibaClient() *client.Client {
	cfg, err := config.Load()
	if err != nil {
		logErr("load config: %v", err)
		os.Exit(ExitGeneral)
	}
	appID := cfg.Rakuten.IchibaAppID
	if appID == "" {
		appID = cfg.Rakuten.AppID
	}
	if appID == "" {
		logErr("Rakuten App ID not set — use RAKUTEN_ICHIBA_APP_ID or RAKUTEN_APP_ID env var or ~/.config/raku-cli/config.yaml")
		os.Exit(ExitAuth)
	}
	return client.New(appID, cfg.Rakuten.AffiliateID, cfg.Rakuten.AccessKey, cfg.Rakuten.Origin, flagTimeout)
}

// HandleError maps errors to exit codes and terminates the process.
func HandleError(err error) {
	if err == nil {
		return
	}
	var authErr *client.AuthError
	var apiErr *client.APIError
	switch {
	case errors.As(err, &authErr):
		logErr("%v", err)
		os.Exit(ExitAuth)
	case errors.As(err, &apiErr):
		logErr("%v", err)
		os.Exit(ExitNetwork)
	default:
		logErr("%v", err)
		os.Exit(ExitGeneral)
	}
}

// Logf writes a message to stderr (unless --quiet).
func Logf(format string, args ...any) {
	logErr(format, args...)
}

func logErr(format string, args ...any) {
	if !flagQuiet {
		fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	}
}

// Pretty returns whether --pretty was set.
func Pretty() bool { return flagPretty }

// Quiet returns whether --quiet was set.
func Quiet() bool { return flagQuiet }
