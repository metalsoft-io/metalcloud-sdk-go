/*
MetalSoft REST API

Testing FirmwareCatalogAPIService

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

func Test_sdk_FirmwareCatalogAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirmwareCatalogAPIService CreateFirmwareCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirmwareCatalogAPI.CreateFirmwareCatalogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareCatalogAPIService DeleteFirmwareCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var firmwareCatalogId float32

		httpRes, err := apiClient.FirmwareCatalogAPI.DeleteFirmwareCatalog(context.Background(), firmwareCatalogId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareCatalogAPIService GetFirmwareCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var firmwareCatalogId float32

		resp, httpRes, err := apiClient.FirmwareCatalogAPI.GetFirmwareCatalog(context.Background(), firmwareCatalogId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareCatalogAPIService GetFirmwareCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirmwareCatalogAPI.GetFirmwareCatalogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirmwareCatalogAPIService UpdateFirmwareCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var firmwareCatalogId float32

		resp, httpRes, err := apiClient.FirmwareCatalogAPI.UpdateFirmwareCatalog(context.Background(), firmwareCatalogId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
