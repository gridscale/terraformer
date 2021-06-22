// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gridscale

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gridscale/gsclient-go/v3"
)

type ServerGenerator struct {
	GridscaleService
}

func (g ServerGenerator) createResources(serverList []gsclient.Server) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, server := range serverList {
		resources = append(resources, terraformutils.NewSimpleResource(
			server.Properties.ObjectUUID,
			server.Properties.Name,
			"gridscale_server",
			"gridscale",
			[]string{}))
	}
	return resources
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	serverList, err := client.GetServerList(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(serverList)
	return nil
}
