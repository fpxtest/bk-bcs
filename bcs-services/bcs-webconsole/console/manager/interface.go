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

package manager

import (
	"context"
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-webconsole/console/types"
	"github.com/gorilla/websocket"
)

// Manager is an interface
type Manager interface {

	//StartExec container web console
	StartExec(http.ResponseWriter, *http.Request, *types.WebSocketConfig, *websocket.Conn) error

	// GetK8sContext
	GetK8sContext(ctx context.Context, clusterID, username string) (string, error)
	GetK8sContextByContainerID(containerID string) (*types.K8sContextByContainerID, error)
	CleanUserPod()
	WritePodData(data *types.UserPodData)
	ReadPodData(sessionID, projectID, clustersID string) (*types.UserPodData, bool)
}
