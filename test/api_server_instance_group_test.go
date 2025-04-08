/*
MetalSoft REST API

Testing ServerInstanceGroupAPIService

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

func Test_sdk_ServerInstanceGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServerInstanceGroupAPIService ApplyProfileToServerInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32
		var serverInstanceProfileId int32

		httpRes, err := apiClient.ServerInstanceGroupAPI.ApplyProfileToServerInstanceGroup(context.Background(), serverInstanceGroupId, serverInstanceProfileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService CreateServerInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.CreateServerInstanceGroup(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService DeleteServerInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		httpRes, err := apiClient.ServerInstanceGroupAPI.DeleteServerInstanceGroup(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetInfrastructureServerInstanceGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetInfrastructureServerInstanceGroups(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroup(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupConfig(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupDriveGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupDriveGroups(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupInterface", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32
		var interfaceId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupInterface(context.Background(), serverInstanceGroupId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupInterfaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupInterfaces(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupNetworkConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfiguration(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupNetworkConfigurationLogicalNetworkById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32
		var logicalNetworkId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationLogicalNetworkById(context.Background(), serverInstanceGroupId, logicalNetworkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupNetworkConfigurationLogicalNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupNetworkConfigurationLogicalNetworks(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService GetServerInstanceGroupServerInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.GetServerInstanceGroupServerInstances(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService UpdateServerInstanceGroupConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		resp, httpRes, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupConfig(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerInstanceGroupAPIService UpdateServerInstanceGroupMeta", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverInstanceGroupId int32

		httpRes, err := apiClient.ServerInstanceGroupAPI.UpdateServerInstanceGroupMeta(context.Background(), serverInstanceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
