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

var defaultFirewallNameList = map[string]int{
	"Admin Server":    0,
	"Blocked Server":  0,
	"Database Server": 0,
	"Mail Server":     0,
	"Web Server":      0,
}

type FirewallGenerator struct {
	GridscaleService
}

func (g FirewallGenerator) createResources(firewallList []gsclient.Firewall) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for idx, firewall := range firewallList {
		if _, isDefault := defaultFirewallNameList[firewall.Properties.Name]; isDefault {
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			firewall.Properties.ObjectUUID,
			fmt.Sprintf("%s-%d", firewall.Properties.Name, idx),
			"gridscale_firewall",
			"gridscale",
			[]string{}))
	}
	return resources
}

func (g *FirewallGenerator) InitResources() error {
	client := g.generateClient()
	firewallList, err := client.GetFirewallList(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(firewallList)
	return nil
}
