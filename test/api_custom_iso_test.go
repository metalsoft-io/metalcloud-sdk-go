/*
MetalSoft REST API

Testing CustomIsoAPIService

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

func Test_sdk_CustomIsoAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomIsoAPIService BootCustomIsoIntoServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customIsoId float32
		var serverId float32

		resp, httpRes, err := apiClient.CustomIsoAPI.BootCustomIsoIntoServer(context.Background(), customIsoId, serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService CreateCustomIso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomIsoAPI.CreateCustomIso(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService DeleteCustomIso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customIsoId float32

		httpRes, err := apiClient.CustomIsoAPI.DeleteCustomIso(context.Background(), customIsoId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService GetCustomIso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customIsoId float32

		resp, httpRes, err := apiClient.CustomIsoAPI.GetCustomIso(context.Background(), customIsoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService GetCustomIsos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomIsoAPI.GetCustomIsos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService MakeCustomIsoPublic", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customIsoId float32

		resp, httpRes, err := apiClient.CustomIsoAPI.MakeCustomIsoPublic(context.Background(), customIsoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomIsoAPIService UpdateCustomIso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customIsoId float32

		resp, httpRes, err := apiClient.CustomIsoAPI.UpdateCustomIso(context.Background(), customIsoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
