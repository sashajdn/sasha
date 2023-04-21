package cassandra

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const cassandraLoggerTag = "cassandra"

func New(serviceName string, cfg Config) (*Cluster, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	var logger = cfg.Logger
	if logger == nil {
		l, err := zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("failed to create logger as not passed: %w", err)
		}

		logger = l.Sugar()
	}

	logger = logger.With(
		zap.String("keyspace", cfg.Keyspace),
		zap.String("service_name", serviceName),
		zap.String("entity", cassandraLoggerTag),
	)

	clusterConfig := gocql.NewCluster(cfg.Hosts...)

	logger.Info("Generated cassandra cluster config; applying options")

	clusterConfig.Keyspace = cfg.Keyspace
	for _, opt := range cfg.Opts {
		opt(clusterConfig)
	}

	logger.Info("Creating cassandra session")

	session, err := clusterConfig.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("create gocql session: %w", err)
	}

	logger.Info("Cassandra session created")

	return &Cluster{
		serviceName:   serviceName,
		session:       session,
		clusterConfig: clusterConfig,
		logger:        logger,
	}, nil
}

type Cluster struct {
	serviceName   string
	config        Config
	clusterConfig *gocql.ClusterConfig
	session       *gocql.Session
	logger        *zap.SugaredLogger
	setupOnce     sync.Once
}

func (c *Cluster) Exec(ctx context.Context, query string) error {
	return c.session.Query(query).WithContext(ctx).Exec()
}

func (c *Cluster) Query(ctx context.Context, query string, consistency gocql.Consistency, dest ...any) error {
	return c.session.Query(query).WithContext(ctx).Scan(dest...)
}

func (c *Cluster) QueryIter(ctx context.Context, query string, consistency gocql.Consistency, dest ...any) gocql.Scanner {
	return c.session.Query(query).WithContext(ctx).Iter().Scanner()
}

func (c *Cluster) Setup(ctx context.Context) error {
	var err error
	c.setupOnce.Do(func() {
		fp := buildConfigPath(c.serviceName)

		if _, err = os.Stat(fp); err != nil {
			err = ErrNoConfigFileFound
			return
		}

		var b []byte
		b, err = ioutil.ReadFile(fp)
		if err != nil {
			err = fmt.Errorf("read file: %w")
			return
		}

		if err = c.Exec(ctx, string(b)); err != nil {
			err = fmt.Errorf("apply config: %w", err)
			return
		}

        // TODO: apply migrations
	})

	if err != nil {
		return fmt.Errorf("setup cluster: %w", err)
	}

	return nil
}

func (c *Cluster) Close() {
	c.logger.Info("Closing cassandra session")
	c.session.Close()
	c.logger.Info("Cassandra session closed")
}
