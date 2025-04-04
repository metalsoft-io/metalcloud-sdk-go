/*
MetalSoft REST API

Testing ServerDefaultCredentialsAPIService

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

func Test_sdk_ServerDefaultCredentialsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServerDefaultCredentialsAPIService CreateServerDefaultCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServerDefaultCredentialsAPI.CreateServerDefaultCredentials(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerDefaultCredentialsAPIService DeleteServerDefaultCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverDefaultCredentialsId float32

		httpRes, err := apiClient.ServerDefaultCredentialsAPI.DeleteServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerDefaultCredentialsAPIService GetServerDefaultCredentialsCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverDefaultCredentialsId float32

		resp, httpRes, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsCredentials(context.Background(), serverDefaultCredentialsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerDefaultCredentialsAPIService GetServerDefaultCredentialsInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverDefaultCredentialsId float32

		resp, httpRes, err := apiClient.ServerDefaultCredentialsAPI.GetServerDefaultCredentialsInfo(context.Background(), serverDefaultCredentialsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerDefaultCredentialsAPIService GetServersDefaultCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServerDefaultCredentialsAPI.GetServersDefaultCredentials(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServerDefaultCredentialsAPIService UpdateServerDefaultCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serverDefaultCredentialsId float32

		resp, httpRes, err := apiClient.ServerDefaultCredentialsAPI.UpdateServerDefaultCredentials(context.Background(), serverDefaultCredentialsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
