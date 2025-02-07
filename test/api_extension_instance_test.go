/*
MetalSoft REST API

Testing ExtensionInstanceAPIService

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

func Test_sdk_ExtensionInstanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExtensionInstanceAPIService CreateExtensionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.ExtensionInstanceAPI.CreateExtensionInstance(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionInstanceAPIService DeleteExtensionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionInstanceId float32

		httpRes, err := apiClient.ExtensionInstanceAPI.DeleteExtensionInstance(context.Background(), extensionInstanceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionInstanceAPIService GetExtensionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionInstanceId float32

		resp, httpRes, err := apiClient.ExtensionInstanceAPI.GetExtensionInstance(context.Background(), extensionInstanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionInstanceAPIService GetExtensionInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExtensionInstanceAPI.GetExtensionInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionInstanceAPIService UpdateExtensionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionInstanceId float32

		resp, httpRes, err := apiClient.ExtensionInstanceAPI.UpdateExtensionInstance(context.Background(), extensionInstanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
