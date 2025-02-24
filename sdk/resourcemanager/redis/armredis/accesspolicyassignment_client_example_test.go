//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAccessPolicyAssignmentCreateUpdate.json
func ExampleAccessPolicyAssignmentClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessPolicyAssignmentClient().BeginCreateUpdate(ctx, "rg1", "cache1", "accessPolicyAssignmentName1", armredis.CacheAccessPolicyAssignment{
		Properties: &armredis.CacheAccessPolicyAssignmentProperties{
			AccessPolicyName: to.Ptr("accessPolicy1"),
			ObjectID:         to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
			ObjectIDAlias:    to.Ptr("TestAADAppRedis"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CacheAccessPolicyAssignment = armredis.CacheAccessPolicyAssignment{
	// 	Name: to.Ptr("accessPolicyAssignmentName1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 	Properties: &armredis.CacheAccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("accessPolicy1"),
	// 		ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		ObjectIDAlias: to.Ptr("TestAADAppRedis"),
	// 		ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAccessPolicyAssignmentDelete.json
func ExampleAccessPolicyAssignmentClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccessPolicyAssignmentClient().BeginDelete(ctx, "rg1", "cache1", "accessPolicyAssignmentName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAccessPolicyAssignmentGet.json
func ExampleAccessPolicyAssignmentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPolicyAssignmentClient().Get(ctx, "rg1", "cache1", "accessPolicyAssignmentName1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CacheAccessPolicyAssignment = armredis.CacheAccessPolicyAssignment{
	// 	Name: to.Ptr("accessPolicyAssignmentName1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
	// 	Properties: &armredis.CacheAccessPolicyAssignmentProperties{
	// 		AccessPolicyName: to.Ptr("accessPolicy1"),
	// 		ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
	// 		ObjectIDAlias: to.Ptr("TestAADAppRedis"),
	// 		ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAccessPolicyAssignmentList.json
func ExampleAccessPolicyAssignmentClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccessPolicyAssignmentClient().NewListPager("rg1", "cache1", nil)
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
		// page.CacheAccessPolicyAssignmentList = armredis.CacheAccessPolicyAssignmentList{
		// 	Value: []*armredis.CacheAccessPolicyAssignment{
		// 		{
		// 			Name: to.Ptr("accessPolicyAssignmentName1"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName1"),
		// 			Properties: &armredis.CacheAccessPolicyAssignmentProperties{
		// 				AccessPolicyName: to.Ptr("accessPolicy1"),
		// 				ObjectID: to.Ptr("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
		// 				ObjectIDAlias: to.Ptr("TestAADAppRedis1"),
		// 				ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("accessPolicyAssignmentName2"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/accessPolicyAssignments"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1/accessPolicyAssignments/accessPolicyAssignmentName2"),
		// 			Properties: &armredis.CacheAccessPolicyAssignmentProperties{
		// 				AccessPolicyName: to.Ptr("accessPolicy2"),
		// 				ObjectID: to.Ptr("7497c918-11ad-41e7-1b0f-7c518a87d0b0"),
		// 				ObjectIDAlias: to.Ptr("TestAADAppRedis2"),
		// 				ProvisioningState: to.Ptr(armredis.AccessPolicyAssignmentProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
