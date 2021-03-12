/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"bk-bscp/internal/database"
	"bk-bscp/internal/types"
)

const (
	// ENVPREFIX is prefix for env variables.
	ENVPREFIX = "BSCP_BCSCONTROLLER"
)

// config for local module.
type config struct {
	viper *viper.Viper
}

// init initialize and check the module configs.
func (c *config) init(localConfigFile string) (*viper.Viper, error) {
	c.viper = viper.GetViper()
	c.viper.SetConfigFile(localConfigFile)

	if err := c.viper.ReadInConfig(); err != nil {
		return nil, err
	}
	c.viper.WatchConfig()

	if err := c.check(); err != nil {
		return nil, err
	}
	return c.viper, nil
}

func (c *config) envName(bindKey string) string {
	return ENVPREFIX + "_" + bindKey
}

// check bind the env vars and checks base config content.
func (c *config) check() error {
	c.viper.BindEnv("server.serviceName", c.envName("SERVICE_NAME"))
	c.viper.SetDefault("server.serviceName", "bk-bscp-bcs-controller")

	c.viper.BindEnv("server.metadata", c.envName("SERVICE_METADATA"))
	c.viper.SetDefault("server.metadata", "bk-bscp-bcs-controller")

	c.viper.BindEnv("server.endpoint.ip", c.envName("ENDPOINT_IP"))
	if !c.viper.IsSet("server.endpoint.ip") {
		return errors.New("config check, missing 'server.endpoint.ip'")
	}
	c.viper.BindEnv("server.endpoint.port", c.envName("ENDPOINT_PORT"))
	if !c.viper.IsSet("server.endpoint.port") {
		return errors.New("config check, missing 'server.endpoint.port'")
	}
	c.viper.BindEnv("server.discoveryTTL", c.envName("DISCOVERY_TTL"))
	c.viper.SetDefault("server.discoveryTTL", 10)

	c.viper.BindEnv("server.queryNewestLimit", c.envName("NEWEST_LIMIT"))
	c.viper.SetDefault("server.queryNewestLimit", database.BSCPQUERYNEWESTLIMIT)

	c.viper.BindEnv("metrics.endpoint", c.envName("METRICS_ENDPOINT"))
	c.viper.SetDefault("metrics.endpoint", ":9100")

	c.viper.BindEnv("metrics.path", c.envName("METRICS_PATH"))
	c.viper.SetDefault("metrics.path", "/metrics")

	c.viper.BindEnv("etcdCluster.endpoints", c.envName("ETCD_ENDPOINTS"))
	if !c.viper.IsSet("etcdCluster.endpoints") {
		return errors.New("config check, missing 'etcdCluster.endpoints'")
	}
	c.viper.BindEnv("etcdCluster.dialTimeout", c.envName("ETCD_DIAL_TIMEOUT"))
	c.viper.SetDefault("etcdCluster.dialTimeout", 10*time.Second)

	c.viper.BindEnv("etcdCluster.tls.certPassword", c.envName("ETCD_TLS_CERT_PASSWORD"))
	c.viper.BindEnv("etcdCluster.tls.caFile", c.envName("ETCD_TLS_CA_FILE"))
	c.viper.BindEnv("etcdCluster.tls.certFile", c.envName("ETCD_TLS_CERT_FILE"))
	c.viper.BindEnv("etcdCluster.tls.keyFile", c.envName("ETCD_TLS_KEY_FILE"))

	c.viper.BindEnv("datamanager.serviceName", c.envName("DM_SERVICE_NAME"))
	c.viper.SetDefault("datamanager.serviceName", "bk-bscp-datamanager")

	c.viper.BindEnv("datamanager.callTimeout", c.envName("DM_CALL_TIMEOUT"))
	c.viper.SetDefault("datamanager.callTimeout", 10*time.Second)

	c.viper.BindEnv("natsmqCluster.endpoints", c.envName("NATS_ENDPOINTS"))
	if !c.viper.IsSet("natsmqCluster.endpoints") {
		return errors.New("config check, missing 'natsmqCluster.endpoints'")
	}
	c.viper.BindEnv("natsmqCluster.tls.certPassword", c.envName("NATS_TLS_CERT_PASSWORD"))
	c.viper.BindEnv("natsmqCluster.tls.caFile", c.envName("NATS_TLS_CA_FILE"))
	c.viper.BindEnv("natsmqCluster.tls.certFile", c.envName("NATS_TLS_CERT_FILE"))
	c.viper.BindEnv("natsmqCluster.tls.keyFile", c.envName("NATS_TLS_KEY_FILE"))

	c.viper.BindEnv("natsmqCluster.timeout", c.envName("NATS_TIMEOUT"))
	c.viper.SetDefault("natsmqCluster.timeout", 10*time.Second)

	c.viper.BindEnv("natsmqCluster.reconWait", c.envName("NATS_RECONWAIT"))
	c.viper.SetDefault("natsmqCluster.reconWait", 3*time.Second)

	c.viper.BindEnv("natsmqCluster.maxRecons", c.envName("NATS_MAXRECONS"))
	c.viper.SetDefault("natsmqCluster.maxRecons", 10)

	c.viper.BindEnv("natsmqCluster.publishTopic", c.envName("NATS_PUB_TOPIC"))
	c.viper.SetDefault("natsmqCluster.publishTopic", types.BCSNatsMQPublishTopic)

	c.viper.BindEnv("logger.directory", c.envName("LOG_DIR"))
	c.viper.SetDefault("logger.directory", "./log")

	c.viper.BindEnv("logger.maxsize", c.envName("LOG_MAXSIZE"))
	c.viper.SetDefault("logger.maxsize", 200)

	c.viper.BindEnv("logger.maxnum", c.envName("LOG_MAXNUM"))
	c.viper.SetDefault("logger.maxnum", 5)

	c.viper.BindEnv("logger.stderr", c.envName("LOG_STDERR"))
	c.viper.SetDefault("logger.stderr", false)

	c.viper.BindEnv("logger.alsoStderr", c.envName("LOG_ALSOSTDERR"))
	c.viper.SetDefault("logger.alsoStderr", false)

	c.viper.BindEnv("logger.level", c.envName("LOG_LEVEL"))
	c.viper.SetDefault("logger.level", 3)

	c.viper.BindEnv("logger.stderrThreshold", c.envName("LOG_STDERR_THRESHOLD"))
	c.viper.SetDefault("logger.stderrThreshold", 2)

	return nil
}
