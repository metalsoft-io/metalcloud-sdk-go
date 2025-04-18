/*
MetalSoft REST API

Testing NetworkAPIService

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

func Test_sdk_NetworkAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkAPIService CreateInfrastructureNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.NetworkAPI.CreateInfrastructureNetwork(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkAPIService DeleteInfrastructureNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var networkId float32

		httpRes, err := apiClient.NetworkAPI.DeleteInfrastructureNetwork(context.Background(), infrastructureId, networkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkAPIService GetInfrastructureNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var networkId float32

		resp, httpRes, err := apiClient.NetworkAPI.GetInfrastructureNetwork(context.Background(), infrastructureId, networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkAPIService GetInfrastructureNetworks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.NetworkAPI.GetInfrastructureNetworks(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
