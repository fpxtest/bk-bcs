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

package namespace

import (
	"os"
	"testing"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/auth/iam"
)

var (
	AppCode   = os.Getenv("APP_CODE")
	AppSecret = os.Getenv("APP_SECRET")

	GateWayHost = os.Getenv("IAM_HOST")
)

var opts = &iam.Options{
	SystemID:    iam.SystemIDBKBCS,
	AppCode:     AppCode,
	AppSecret:   AppSecret,
	External:    false,
	GateWayHost: GateWayHost,
	Metric:      false,
	Debug:       true,
}

func newBcsNamespacePermCli() (*BCSNamespacePerm, error) {
	cli, err := iam.NewIamClient(opts)
	if err != nil {
		return nil, err
	}

	return NewBCSNamespacePermClient(cli), nil
}

func TestPerm_CanCreateNamespace(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanCreateNamespace(username, projectID, clusterID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanViewNamespace(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanViewNamespace(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanListNamespace(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanListNamespace(username, projectID, clusterID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanUpdateNamespace(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanUpdateNamespace(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanDeleteNamespace(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanDeleteNamespace(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanCreateNamespaceScopedResource(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanCreateNamespaceScopedResource(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanViewNamespaceScopedResource(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanViewNamespaceScopedResource(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanUpdateNamespaceScopedResource(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanUpdateNamespaceScopedResource(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}

func TestPerm_CanDeleteNamespaceScopedResource(t *testing.T) {
	cli, err := newBcsNamespacePermCli()
	if err != nil {
		t.Fatal(err)
	}

	var (
		projectID = os.Getenv("PROJECT_ID")
		clusterID = os.Getenv("CLUSTER_ID")
		username  = os.Getenv("PERM_USERNAME")
	)
	allow, url, err := cli.CanDeleteNamespaceScopedResource(username, projectID, clusterID, "bcs-system")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(allow, url)
}
