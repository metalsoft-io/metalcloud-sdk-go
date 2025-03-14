/*
MetalSoft REST API

Testing ConfigurationAPIService

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

func Test_sdk_ConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigurationAPIService GetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationAPI.GetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService PatchConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filter string

		resp, httpRes, err := apiClient.ConfigurationAPI.PatchConfiguration(context.Background(), filter).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationAPIService PutConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filter string

		resp, httpRes, err := apiClient.ConfigurationAPI.PutConfiguration(context.Background(), filter).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
