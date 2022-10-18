//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DataCollectionEndpointsClient contains the methods for the DataCollectionEndpoints group.
// Don't use this type directly, use NewDataCollectionEndpointsClient() instead.
type DataCollectionEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataCollectionEndpointsClient creates a new instance of DataCollectionEndpointsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataCollectionEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataCollectionEndpointsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DataCollectionEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates or updates a data collection endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
// options - DataCollectionEndpointsClientCreateOptions contains the optional parameters for the DataCollectionEndpointsClient.Create
// method.
func (client *DataCollectionEndpointsClient) Create(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientCreateOptions) (DataCollectionEndpointsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataCollectionEndpointsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DataCollectionEndpointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataCollectionEndpointsClient) createHandleResponse(resp *http.Response) (DataCollectionEndpointsClientCreateResponse, error) {
	result := DataCollectionEndpointsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a data collection endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
// options - DataCollectionEndpointsClientDeleteOptions contains the optional parameters for the DataCollectionEndpointsClient.Delete
// method.
func (client *DataCollectionEndpointsClient) Delete(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientDeleteOptions) (DataCollectionEndpointsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataCollectionEndpointsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataCollectionEndpointsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataCollectionEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the specified data collection endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
// options - DataCollectionEndpointsClientGetOptions contains the optional parameters for the DataCollectionEndpointsClient.Get
// method.
func (client *DataCollectionEndpointsClient) Get(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientGetOptions) (DataCollectionEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataCollectionEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataCollectionEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataCollectionEndpointsClient) getHandleResponse(resp *http.Response) (DataCollectionEndpointsClientGetResponse, error) {
	result := DataCollectionEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all data collection endpoints in the specified resource group.
// Generated from API version 2021-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DataCollectionEndpointsClientListByResourceGroupOptions contains the optional parameters for the DataCollectionEndpointsClient.ListByResourceGroup
// method.
func (client *DataCollectionEndpointsClient) NewListByResourceGroupPager(resourceGroupName string, options *DataCollectionEndpointsClientListByResourceGroupOptions) *runtime.Pager[DataCollectionEndpointsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionEndpointsClientListByResourceGroupResponse]{
		More: func(page DataCollectionEndpointsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionEndpointsClientListByResourceGroupResponse) (DataCollectionEndpointsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionEndpointsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataCollectionEndpointsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionEndpointsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataCollectionEndpointsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataCollectionEndpointsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataCollectionEndpointsClient) listByResourceGroupHandleResponse(resp *http.Response) (DataCollectionEndpointsClientListByResourceGroupResponse, error) {
	result := DataCollectionEndpointsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResourceListResult); err != nil {
		return DataCollectionEndpointsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all data collection endpoints in the specified subscription
// Generated from API version 2021-09-01-preview
// options - DataCollectionEndpointsClientListBySubscriptionOptions contains the optional parameters for the DataCollectionEndpointsClient.ListBySubscription
// method.
func (client *DataCollectionEndpointsClient) NewListBySubscriptionPager(options *DataCollectionEndpointsClientListBySubscriptionOptions) *runtime.Pager[DataCollectionEndpointsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionEndpointsClientListBySubscriptionResponse]{
		More: func(page DataCollectionEndpointsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionEndpointsClientListBySubscriptionResponse) (DataCollectionEndpointsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionEndpointsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataCollectionEndpointsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionEndpointsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataCollectionEndpointsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataCollectionEndpointsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/dataCollectionEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataCollectionEndpointsClient) listBySubscriptionHandleResponse(resp *http.Response) (DataCollectionEndpointsClientListBySubscriptionResponse, error) {
	result := DataCollectionEndpointsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResourceListResult); err != nil {
		return DataCollectionEndpointsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates part of a data collection endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
// options - DataCollectionEndpointsClientUpdateOptions contains the optional parameters for the DataCollectionEndpointsClient.Update
// method.
func (client *DataCollectionEndpointsClient) Update(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientUpdateOptions) (DataCollectionEndpointsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataCollectionEndpointsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DataCollectionEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DataCollectionEndpointsClient) updateHandleResponse(resp *http.Response) (DataCollectionEndpointsClientUpdateResponse, error) {
	result := DataCollectionEndpointsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsClientUpdateResponse{}, err
	}
	return result, nil
}