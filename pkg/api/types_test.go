package api

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

var unmarshalled = &OpenShiftManagedCluster{
	ID:       "id",
	Location: "location",
	Name:     "name",
	Plan: &ResourcePurchasePlan{
		Name:          "plan.name",
		Product:       "plan.product",
		PromotionCode: "plan.promotionCode",
		Publisher:     "plan.publisher",
	},
	Tags: map[string]string{
		"tags.k1": "v1",
		"tags.k2": "v2",
	},
	Type: "type",
	Properties: &Properties{
		ProvisioningState: "properties.provisioningState",
		OpenShiftVersion:  "properties.openShiftVersion",
		PublicHostname:    "properties.publicHostname",
		NetworkProfile: &NetworkProfile{
			VnetCIDR:   "properties.networkProfile.vnetCidr",
			PeerVnetID: "properties.networkProfile.peerVnetId",
		},
		RouterProfiles: []RouterProfile{
			{
				Name:            "properties.routerProfiles.0.name",
				PublicSubdomain: "properties.routerProfiles.0.publicSubdomain",
				FQDN:            "properties.routerProfiles.0.fqdn",
			},
			{
				Name:            "properties.routerProfiles.1.name",
				PublicSubdomain: "properties.routerProfiles.1.publicSubdomain",
				FQDN:            "properties.routerProfiles.1.fqdn",
			},
		},
		FQDN: "properties.fqdn",
		AuthProfile: &AuthProfile{
			IdentityProviders: []IdentityProvider{
				{
					Name: "properties.authProfile.identityProviders.0.name",
					Provider: &AADIdentityProvider{
						Kind:     "AADIdentityProvider",
						ClientID: "properties.authProfile.identityProviders.0.provider.clientId",
						Secret:   "properties.authProfile.identityProviders.0.provider.secret",
						TenantID: "properties.authProfile.identityProviders.0.provider.tenantId",
					},
				},
			},
		},
		AgentPoolProfiles: []AgentPoolProfile{
			{
				Count:      1,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				Role:       "properties.agentPoolProfiles.0.role",
			},
			{
				Name:       "properties.agentPoolProfiles.0.name",
				Count:      2,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				OSType:     "properties.agentPoolProfiles.0.osType",
				Role:       "properties.agentPoolProfiles.0.role",
			},
			{
				Name:       "properties.agentPoolProfiles.0.name",
				Count:      1,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				OSType:     "properties.agentPoolProfiles.0.osType",
				Role:       "master",
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "properties.servicePrincipalProfile.clientId",
			Secret:   "properties.servicePrincipalProfile.secret",
		},
	},
}

var marshalled = []byte(`{
	"id": "id",
	"location": "location",
	"name": "name",
	"plan": {
		"name": "plan.name",
		"product": "plan.product",
		"promotionCode": "plan.promotionCode",
		"publisher": "plan.publisher"
	},
	"tags": {
		"tags.k1": "v1",
		"tags.k2": "v2"
	},
	"type": "type",
	"properties": {
		"provisioningState": "properties.provisioningState",
		"openShiftVersion": "properties.openShiftVersion",
		"publicHostname": "properties.publicHostname",
		"fqdn": "properties.fqdn",
		"networkProfile": {
			"vnetCidr": "properties.networkProfile.vnetCidr",
			"peerVnetId": "properties.networkProfile.peerVnetId"
		},
		"routerProfiles": [
			{
				"name": "properties.routerProfiles.0.name",
				"publicSubdomain": "properties.routerProfiles.0.publicSubdomain",
				"fqdn": "properties.routerProfiles.0.fqdn"
			},
			{
				"name": "properties.routerProfiles.1.name",
				"publicSubdomain": "properties.routerProfiles.1.publicSubdomain",
				"fqdn": "properties.routerProfiles.1.fqdn"
			}
		],
		"agentPoolProfiles": [
			{
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 2,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "master"
			}
		],
		"authProfile": {
			"identityProviders": [
				{
					"name": "properties.authProfile.identityProviders.0.name",
					"provider": {
						"kind": "AADIdentityProvider",
						"clientId": "properties.authProfile.identityProviders.0.provider.clientId",
						"secret": "properties.authProfile.identityProviders.0.provider.secret",
						"tenantId": "properties.authProfile.identityProviders.0.provider.tenantId"
					}
				}
			]
		},
		"servicePrincipalProfile": {
			"clientId": "properties.servicePrincipalProfile.clientId",
			"secret": "properties.servicePrincipalProfile.secret"
		}
	}
}`)

func TestMarshal(t *testing.T) {
	b, err := json.MarshalIndent(unmarshalled, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, marshalled) {
		t.Errorf("json.MarshalIndent returned unexpected result\n%s\n", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	var oc *OpenShiftManagedCluster
	err := json.Unmarshal(marshalled, &oc)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(oc, unmarshalled) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", oc)
	}
}
