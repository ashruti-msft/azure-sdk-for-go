//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolRestorePoints.json
func ExampleSQLPoolRestorePointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLPoolRestorePointsClient().NewListPager("Default-SQL-SouthEastAsia", "testserver", "testDatabase", nil)
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
		// page.RestorePointListResult = armsynapse.RestorePointListResult{
		// 	Value: []*armsynapse.RestorePoint{
		// 		{
		// 			Name: to.Ptr("131546477590000000"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Synapse/workspaces/testserver/sqlPools/testDatabase/restorePoints/131546477590000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsynapse.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
		// 				RestorePointLabel: to.Ptr("mylabel1"),
		// 				RestorePointType: to.Ptr(armsynapse.RestorePointTypeDISCRETE),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("131553636140000000"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Synapse/workspaces/testserver/sqlPools/testDatabase/restorePoints/131553636140000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsynapse.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T03:40:14Z"); return t}()),
		// 				RestorePointLabel: to.Ptr("mylabel2"),
		// 				RestorePointType: to.Ptr(armsynapse.RestorePointTypeDISCRETE),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("131553619750000000"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Synapse/workspaces/testserver/sqlPools/testDatabase/restorePoints/131553619750000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsynapse.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T03:12:55Z"); return t}()),
		// 				RestorePointType: to.Ptr(armsynapse.RestorePointTypeDISCRETE),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolRestorePoints.json
func ExampleSQLPoolRestorePointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLPoolRestorePointsClient().BeginCreate(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", armsynapse.CreateSQLPoolRestorePointDefinition{
		RestorePointLabel: to.Ptr("mylabel"),
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
	// res.RestorePoint = armsynapse.RestorePoint{
	// 	Name: to.Ptr("131546477590000000"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Synapse/workspaces/testserver/sqlPools/testDatabase/restorePoints/131546477590000000"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsynapse.RestorePointProperties{
	// 		RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
	// 		RestorePointLabel: to.Ptr("mylabel"),
	// 		RestorePointType: to.Ptr(armsynapse.RestorePointTypeDISCRETE),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsGet.json
func ExampleSQLPoolRestorePointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLPoolRestorePointsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testws", "testpool", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armsynapse.RestorePoint{
	// 	Name: to.Ptr("131546477590000000"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Synapse/workspaces/testws/sqlPools/testpool/restorePoints/131546477590000000"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsynapse.RestorePointProperties{
	// 		RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
	// 		RestorePointLabel: to.Ptr("mylabel"),
	// 		RestorePointType: to.Ptr(armsynapse.RestorePointTypeDISCRETE),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsDelete.json
func ExampleSQLPoolRestorePointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSQLPoolRestorePointsClient().Delete(ctx, "Default-SQL-SouthEastAsia", "testws", "testpool", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}