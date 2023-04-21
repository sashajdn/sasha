package dao

import (
	"context"
	"sync"
	"time"

	"github.com/sashajdn/sasha/libraries/cassandra"
	"github.com/sashajdn/sasha/libraries/environment"
	"go.uber.org/zap"
)

var (
	cluster *cassandra.Cluster
	once    sync.Once
)

func Init(serviceName string, env environment.Cassandra, logger *zap.SugaredLogger) error {
	var err error
	once.Do(func() {
		clusterConfig := cassandra.NewConfigFromEnv(env, logger)

		cluster, err = cassandra.New(serviceName, clusterConfig)
		if err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = cluster.Setup(ctx); err != nil {
			return
		}
	})

	if err != nil {
		return err
	}

	return nil
}
