package cmd

import (
	"errors"

	killgrave "github.com/friendsofgo/killgrave/internal"
	server "github.com/friendsofgo/killgrave/internal/server/http"
	"github.com/radovskyb/watcher"
	"github.com/spf13/cobra"
)

var _version = "unknown_version"

const (
	_defaultImpostersPath = "imposters"
	_defaultConfigFile    = ""
	_defaultHost          = "localhost"
	_defaultPort          = 3000
	_defaultProxyMode     = killgrave.ProxyNone
	_defaultStrictSlash   = true

	_impostersFlag = "imposters"
	_configFlag    = "config"
	_hostFlag      = "host"
	_portFlag      = "port"
	_watcherFlag   = "watcher"
	_secureFlag    = "secure"
	_proxyModeFlag = "proxy-mode"
	_proxyURLFlag  = "proxy-url"
)

var (
	errGetDataFromImpostersFlag = errors.New("error trying to get data from imposters flag")
	errGetDataFromHostFlag      = errors.New("error trying to get data from host flag")
	errGetDataFromPortFlag      = errors.New("error trying to get data from port flag")
	errGetDataFromSecureFlag    = errors.New("error trying to get data from secure flag")
	errMandatoryURL             = errors.New("the field proxy-url is mandatory if you selected a proxy mode")
)

// NewKillgraveCmd returns cobra.Command to run killgrave command
func NewKillgraveCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runHTTP(cmd *cobra.Command, cfg killgrave.Config) error { _ = "STUB: not implemented"; return nil }

// TODO: refactor the method NewServer of the pkg server/http should be contain how to initialize the http server
func runServer(cfg killgrave.Config) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

func runWatcher(cfg killgrave.Config, currentSrv *server.Server) (*watcher.Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareConfig(cmd *cobra.Command) (killgrave.Config, error) {
	_ = "STUB: not implemented"
	return *new(killgrave.Config), nil
}

func configureProxyMode(cmd *cobra.Command, cfg *killgrave.Config) error {
	_ = "STUB: not implemented"
	return nil
}
