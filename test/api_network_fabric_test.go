/*
MetalSoft REST API

Testing NetworkFabricAPIService

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

func Test_sdk_NetworkFabricAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkFabricAPIService AddNetworkEquipmentsToFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId int32

		resp, httpRes, err := apiClient.NetworkFabricAPI.AddNetworkEquipmentsToFabric(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService CreateNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkFabricAPI.CreateNetworkFabric(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService DeleteNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId float32

		httpRes, err := apiClient.NetworkFabricAPI.DeleteNetworkFabric(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService GetFabricAndNetworkEquipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId int32

		resp, httpRes, err := apiClient.NetworkFabricAPI.GetFabricAndNetworkEquipment(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService GetNetworkFabricById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId float32

		resp, httpRes, err := apiClient.NetworkFabricAPI.GetNetworkFabricById(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService GetNetworkFabrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkFabricAPI.GetNetworkFabrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService RemoveNetworkEquipmentFromFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId int32
		var networkEquipmentId int32

		resp, httpRes, err := apiClient.NetworkFabricAPI.RemoveNetworkEquipmentFromFabric(context.Background(), networkFabricId, networkEquipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricAPIService UpdateNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId int32

		resp, httpRes, err := apiClient.NetworkFabricAPI.UpdateNetworkFabric(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
