package cassandra

import (
	"errors"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/sashajdn/sasha/libraries/environment"
	"go.uber.org/zap"
)

const (
	defaultConfigDir      = "config"
	defaultConfigFileName = "cassandra.cql"
)

var (
	ErrInvalidConfig     = errors.New("invalid config")
	ErrNoConfigFileFound = errors.New("no config file found")
)

type Opt func(c *gocql.ClusterConfig)

func WithConsistencyLevel(consistencyLevel gocql.Consistency) Opt {
	return func(c *gocql.ClusterConfig) {
		c.Consistency = consistencyLevel
	}
}

func WithProtocolVersion(protocolVersion int) Opt {
	return func(c *gocql.ClusterConfig) {
		c.ProtoVersion = protocolVersion
	}
}

func NewConfigFromEnv(env environment.Cassandra, logger *zap.SugaredLogger) Config {
	return Config{
		Hosts:    env.SeedNodeIPs,
		Keyspace: env.Keyspace,
		Logger:   logger,
	}
}

type Config struct {
	Keyspace   string
	Hosts      []string
	Logger     *zap.SugaredLogger
	Opts       []Opt
	configPath string
}

func (c Config) Validate() error {
	if len(c.Hosts) == 0 {
		return fmt.Errorf("no hosts set: %w", ErrInvalidConfig)
	}

	if c.Keyspace == "" {
		return fmt.Errorf("empty keyspace: %w", ErrInvalidConfig)
	}

	return nil
}

func (c Config) ConfigPath() string {
	if c.configPath != "" {
		return c.configPath
	}

	return defaultConfigPath
}
