/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package autoscaler

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	cmproto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	cmoptions "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/options"
)

const (
	templateName = "AutoScalerValues"
)

var valuesTemplate = `# Generated by bcs-cluster-manager

replicaCount: {{.ReplicaCount}}

namespace: {{.Namespace}}

command:
- ./bcs-cluster-autoscaler
- --v=4
- --stderrthreshold=info
- --namespace={{.Namespace}}
- --cloud-provider=bcs
- --estimator=clusterresource
{{range .Nodes}}- --nodes={{.}}
{{end}}
- --expander={{.Expander}}
- --skip-nodes-with-local-storage={{.SkipNodesWithLocalStorage}}
- --initial-node-group-backoff-duration=10s
- --max-node-group-backoff-duration=3m
- --node-group-backoff-reset-timeout=5m
- --scale-down-enabled={{.IsScaleDownEnable}}
- --max-empty-bulk-delete={{.MaxEmptyBulkDelete}}
- --scale-down-unneeded-time={{.ScaleDownUnneededTime}}
- --scale-down-utilization-threshold={{.ScaleDownUtilizationThreahold}}
- --scale-down-gpu-utilization-threshold={{.ScaleDownGpuUtilizationThreshold}}
- --ok-total-unready-count={{.OkTotalUnreadyCount}}
- --max-total-unready-percentage={{.MaxTotalUnreadyPercentage}}
- --scale-down-unready-time={{.ScaleDownUnreadyTime}}
- --buffer-resource-ratio={{.BufferResourceRatio}}
- --max-graceful-termination-sec={{.MaxGracefulTerminationSec}}
- --scan-interval={{.ScanInterval}}
- --max-node-provision-time={{.MaxNodeProvisionTime}}
- --scale-up-from-zero={{.ScaleUpFromZero}}
- --scale-down-delay-after-add={{.ScaleDownDelayAfterAdd}}
- --scale-down-delay-after-delete={{.ScaleDownDelayAfterDelete}}
- --scale-down-delay-after-failure={{.ScaleDownDelayAfterFailure}}

env:
  apiAddress: "{{.APIAddress}}"
  token: "{{.Token}}"
  operator: "bcs"
  encryption: "no"

nodeSelector: null

tolerations:
- effect: "NoSchedule"
  key: "node-role.kubernetes.io/master"
  operator: "Exists"

affinity:
  nodeAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 10
      preference:
        matchExpressions:
        - key: node-role.kubernetes.io/master
          operator: Exists
`

// AutoScalerValues is the values for the autoscaler application
type AutoScalerValues struct {
	Namespace                        string
	APIAddress                       string
	Token                            string
	ReplicaCount                     int
	Nodes                            []string
	Expander                         string
	SkipNodesWithLocalStorage        bool
	IsScaleDownEnable                bool
	MaxEmptyBulkDelete               uint32
	ScaleDownUnneededTime            time.Duration
	ScaleDownUtilizationThreahold    float64
	ScaleDownGpuUtilizationThreshold float64
	OkTotalUnreadyCount              uint32
	MaxTotalUnreadyPercentage        uint32
	ScaleDownUnreadyTime             time.Duration
	BufferResourceRatio              float64
	MaxGracefulTerminationSec        uint32
	ScanInterval                     time.Duration
	MaxNodeProvisionTime             time.Duration
	ScaleUpFromZero                  bool
	ScaleDownDelayAfterAdd           time.Duration
	ScaleDownDelayAfterDelete        time.Duration
	ScaleDownDelayAfterFailure       time.Duration
}

// AutoScaler component paras
type AutoScaler struct {
	// NodeGroups cluster nodeGroup list
	NodeGroups []cmproto.NodeGroup
	// AutoScalingOption autoScaling deploy paras
	AutoScalingOption *cmproto.ClusterAutoScalingOption
	// Replicas
	Replicas int
}

// GetValues get autoScaler values
func (as *AutoScaler) GetValues() (string, error) {
	if len(as.NodeGroups) == 0 {
		as.Replicas = 0
	}
	tmpl, err := template.New(templateName).Parse(valuesTemplate)
	if err != nil {
		return "", err
	}

	// set vars
	op := cmoptions.GetGlobalCMOptions()
	values := AutoScalerValues{
		Namespace:    op.ComponentDeploy.AutoScaler.ReleaseNamespace,
		APIAddress:   op.ComponentDeploy.BCSAPIGateway,
		Token:        op.ComponentDeploy.Token,
		ReplicaCount: as.Replicas,
	}
	values.Expander = as.AutoScalingOption.Expander
	values.SkipNodesWithLocalStorage = as.AutoScalingOption.SkipNodesWithLocalStorage
	values.IsScaleDownEnable = as.AutoScalingOption.IsScaleDownEnable
	values.MaxEmptyBulkDelete = as.AutoScalingOption.MaxEmptyBulkDelete
	values.ScaleDownUnneededTime = time.Duration(as.AutoScalingOption.ScaleDownUnneededTime) * time.Second
	values.ScaleDownUtilizationThreahold = float64(as.AutoScalingOption.ScaleDownUtilizationThreahold) / 100
	values.ScaleDownGpuUtilizationThreshold = float64(as.AutoScalingOption.ScaleDownGpuUtilizationThreshold) / 100
	values.OkTotalUnreadyCount = as.AutoScalingOption.OkTotalUnreadyCount
	values.MaxTotalUnreadyPercentage = as.AutoScalingOption.MaxTotalUnreadyPercentage
	values.ScaleDownUnreadyTime = time.Duration(as.AutoScalingOption.ScaleDownUnreadyTime) * time.Second
	values.BufferResourceRatio = 1 - float64(as.AutoScalingOption.BufferResourceRatio)/100
	values.MaxGracefulTerminationSec = as.AutoScalingOption.MaxGracefulTerminationSec
	values.ScanInterval = time.Duration(as.AutoScalingOption.ScanInterval) * time.Second
	values.MaxNodeProvisionTime = time.Duration(as.AutoScalingOption.MaxNodeProvisionTime) * time.Second
	values.ScaleUpFromZero = as.AutoScalingOption.ScaleUpFromZero
	values.ScaleDownDelayAfterAdd = time.Duration(as.AutoScalingOption.ScaleDownDelayAfterAdd) * time.Second
	values.ScaleDownDelayAfterDelete = time.Duration(as.AutoScalingOption.ScaleDownDelayAfterDelete) * time.Second
	values.ScaleDownDelayAfterFailure = time.Duration(as.AutoScalingOption.ScaleDownDelayAfterFailure) * time.Second
	for _, v := range as.NodeGroups {
		if len(v.NodeGroupID) != 0 && v.AutoScaling != nil {
			values.Nodes = append(values.Nodes,
				fmt.Sprintf("%d:%d:%s", v.AutoScaling.MinSize, v.AutoScaling.MaxSize, v.NodeGroupID))
		}
	}

	// parse
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, values); err != nil {
		return "", err
	}
	return buf.String(), nil

}