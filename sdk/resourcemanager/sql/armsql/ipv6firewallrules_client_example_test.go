//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleList.json
func ExampleIPv6FirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPv6FirewallRulesClient().NewListByServerPager("firewallrulecrudtest-12", "firewallrulecrudtest-6285", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IPv6FirewallRuleListResult = armsql.IPv6FirewallRuleListResult{
		// 	Value: []*armsql.IPv6FirewallRule{
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-2304"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-2304"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-3927"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-3927"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5370"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5370"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5767"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5767"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0002"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0002"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleGet.json
func ExampleIPv6FirewallRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPv6FirewallRulesClient().Get(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-2304", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPv6FirewallRule = armsql.IPv6FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-2304"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-2304"),
	// 	Properties: &armsql.IPv6ServerFirewallRuleProperties{
	// 		EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
	// 		StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleCreate.json
func ExampleIPv6FirewallRulesClient_CreateOrUpdate_createAnIPv6FirewallRuleMaxMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPv6FirewallRulesClient().CreateOrUpdate(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-5370", armsql.IPv6FirewallRule{
		Properties: &armsql.IPv6ServerFirewallRuleProperties{
			EndIPv6Address:   to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
			StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPv6FirewallRule = armsql.IPv6FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-5370"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5370"),
	// 	Properties: &armsql.IPv6ServerFirewallRuleProperties{
	// 		EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
	// 		StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleUpdate.json
func ExampleIPv6FirewallRulesClient_CreateOrUpdate_updateAnIPv6FirewallRuleMaxMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPv6FirewallRulesClient().CreateOrUpdate(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-3927", armsql.IPv6FirewallRule{
		Properties: &armsql.IPv6ServerFirewallRuleProperties{
			EndIPv6Address:   to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
			StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPv6FirewallRule = armsql.IPv6FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-3927"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-3927"),
	// 	Properties: &armsql.IPv6ServerFirewallRuleProperties{
	// 		EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
	// 		StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleDelete.json
func ExampleIPv6FirewallRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIPv6FirewallRulesClient().Delete(ctx, "firewallrulecrudtest-9886", "firewallrulecrudtest-2368", "firewallrulecrudtest-7011", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
