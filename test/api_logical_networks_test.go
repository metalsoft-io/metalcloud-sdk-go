/*
MetalSoft REST API

Testing LogicalNetworksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func Test_sdk_LogicalNetworksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogicalNetworksAPIService CreateLogicalNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogicalNetworksAPI.CreateLogicalNetwork(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogicalNetworksAPIService DeleteLogicalNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logicalNetworkId float32

		httpRes, err := apiClient.LogicalNetworksAPI.DeleteLogicalNetwork(context.Background(), logicalNetworkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogicalNetworksAPIService GetAllLogicalNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogicalNetworksAPI.GetAllLogicalNetworks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogicalNetworksAPIService GetLogicalNetworkById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logicalNetworkId float32

		resp, httpRes, err := apiClient.LogicalNetworksAPI.GetLogicalNetworkById(context.Background(), logicalNetworkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogicalNetworksAPIService UpdateLogicalNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var logicalNetworkId float32

		resp, httpRes, err := apiClient.LogicalNetworksAPI.UpdateLogicalNetwork(context.Background(), logicalNetworkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
