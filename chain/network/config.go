package network

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"io/ioutil"
	"os"
	"path/filepath"
	"regalcoin/chain/interfaces"
	"time"
	"github.com/adrg/xdg"
	"github.com/urfave/cli/v2"
)

const Prefix = "regalp2p"
const PrefixChainConfig = "regalchain"
const ContextKey = "config"

var configFile = filepath.Join(Prefix, "config.json")
var genesisFile = filepath.Join(PrefixChainConfig, "genesis.json")

var DefaultConfig = Config{
	BootstrapPeers: []string{},
	DialTimeout: time.Minute,
	Protocols: ServiceDefinitions,

}

type Config struct {
	NetVersion string `json:"-"`
	Path string `json:"-"`
	Exist bool `json:"-"`
	BootstrapPeers []string
	DialTimeout time.Duration
	Protocols []string
	Blockchain interfaces.IBlockchain
	Nodes interfaces.INode
	Genesis string `json:"-"`
	SyncWorkerCount int
	LSyncWorkerCount int
	MaxPeers int
}

func init() {
	for _, maddr := range dht.DefaultBootstrapPeers{
		DefaultConfig.BootstrapPeers = append(DefaultConfig.BootstrapPeers, maddr.String())
	}
}

func (c *Config) Save() error {
	log.Infoln("Saving regalcoin configuration file to", c.Path)
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if c.Path == "" {
		c.Path, err = xdg.ConfigFile(configFile)
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(c.Path, data, 0o744)

}

func (c *Config) Apply(ctx *cli.Context) {
	if ctx.IsSet("workers") {
		if ctx.Command.Name == "start" {
			c.SyncWorkerCount = ctx.Int("workers")
		} else if ctx.Command.Name == "start-local" {
			c.LSyncWorkerCount = 0
		}

		if ctx.IsSet("protocols") {
			c.Protocols = ctx.StringSlice("protocols")
		}

		if ctx.IsSet("maxPeers") {
			c.MaxPeers = ctx.Int("max_peers")
		}


	}
}

func (c *Config) BootstrapAddrsInfos() ([]peer.AddrInfo, error) {
	var pis []peer.AddrInfo
	for _, madrStr := range c.BootstrapPeers {
		maddr, err := ma.NewMultiaddr(madrStr)
		if err != nil {
			return nil, err
		}
		pi, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			return nil, err
		}
		pis = append(pis, *pi)
	}
	return pis, nil
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		// If no configuration file was given use xdg file.
		var err error
		path, err = xdg.ConfigFile(configFile)
		if err != nil {
			return nil, err
		}
	}

	log.Debugln("Loading configuration from:", path)
	config := DefaultConfig
	config.Path = path
	data, err := ioutil.ReadFile(path)
	genesis, _ := LoadGenesis(genesisFile)

	if err == nil {
		err = json.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}
		config.Exist = true
		config.Genesis = genesis
		return &config, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	} else {
		return &config, config.Save()
	}
}

func LoadGenesis(path string) (string, error) {
	if path == "" {
		// If no configuration file was given use xdg file.
		var err error
		path, err = xdg.ConfigFile(genesisFile)
		if err != nil {
			return "", err
		}
	}
	log.Debugln("Loading genesis from:", path)
	data, err := ioutil.ReadFile(path)
	if err == nil {
		genesisString := string(data)

		return genesisString, nil
	} else if !os.IsNotExist(err) {
		return "", err
	} else {
		return "", nil
	}
}

func FillContext(c *cli.Context) (context.Context, *Config, error) {
	conf, err := LoadConfig(c.String("config"))
	if err != nil {
		return c.Context, nil, err
	}

	// Apply command line argument configurations.
	conf.Apply(c)

	// Print full configuration.
	log.Debugln("Configuration (CLI params overwrite file config):\n", conf)

	// Populate the context with the configuration.
	return context.WithValue(c.Context, ContextKey, conf), conf, nil
}

func FromContext(ctx context.Context) (*Config, error) {
	obj := ctx.Value(ContextKey)
	if obj == nil {
		return nil, errors.New("config not found in context")
	}

	config, ok := obj.(*Config)
	if !ok {
		return nil, errors.New("config in context has wrong type")
	}

	return config, nil
}