//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsGet.json
func ExampleScopeAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewScopeAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"subscriptions/subscriptionC",
		"subscriptionCAssignment",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsPut.json
func ExampleScopeAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewScopeAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"subscriptions/subscriptionC",
		"subscriptionCAssignment",
		armmanagednetwork.ScopeAssignment{
			Properties: &armmanagednetwork.ScopeAssignmentProperties{
				AssignedManagedNetwork: to.Ptr("/subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsDelete.json
func ExampleScopeAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewScopeAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"subscriptions/subscriptionC",
		"subscriptionCAssignment",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsList.json
func ExampleScopeAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewScopeAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("subscriptions/subscriptionC",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}