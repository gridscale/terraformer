// Copyright 2021 The Terraformer Authors.
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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gridscale/gsclient-go/v3"
)

type NetworkGenerator struct {
	GridscaleService
}

func (g NetworkGenerator) createResources(networkList []gsclient.Network) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for idx, network := range networkList {
		if network.Properties.PublicNet {
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			network.Properties.ObjectUUID,
			fmt.Sprintf("%s-%d", network.Properties.Name, idx),
			"gridscale_network",
			"gridscale",
			[]string{}))
	}
	return resources
}

func (g *NetworkGenerator) InitResources() error {
	client := g.generateClient()
	networkList, err := client.GetNetworkList(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(networkList)
	return nil
}
