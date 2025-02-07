/*
MetalSoft REST API

Testing NetworkFabricsAPIService

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

func Test_sdk_NetworkFabricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkFabricsAPIService CreateNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkFabricsAPI.CreateNetworkFabric(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricsAPIService DeleteNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId float32

		httpRes, err := apiClient.NetworkFabricsAPI.DeleteNetworkFabric(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricsAPIService GetNetworkFabricById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId float32

		resp, httpRes, err := apiClient.NetworkFabricsAPI.GetNetworkFabricById(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricsAPIService GetNetworkFabrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NetworkFabricsAPI.GetNetworkFabrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkFabricsAPIService UpdateNetworkFabric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkFabricId float32

		resp, httpRes, err := apiClient.NetworkFabricsAPI.UpdateNetworkFabric(context.Background(), networkFabricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
