package cassandra

import "fmt"

func buildConfigPath(serviceName string) string {
	return fmt.Sprintf("/%s/%s/%s", serviceName, defaultConfigDir, defaultConfigFileName)
}
