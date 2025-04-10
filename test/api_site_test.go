/*
MetalSoft REST API

Testing SiteAPIService

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

func Test_sdk_SiteAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SiteAPIService CreateSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SiteAPI.CreateSite(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService DecommissionSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		httpRes, err := apiClient.SiteAPI.DecommissionSite(context.Background(), siteId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetAgents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.GetAgents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.GetSite(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetSiteConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.GetSiteConfig(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetSiteControllerOneLiner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.GetSiteControllerOneLiner(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetSites", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SiteAPI.GetSites(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService GetSitesStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SiteAPI.GetSitesStatistics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService UpdateSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.UpdateSite(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SiteAPIService UpdateSiteConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId float32

		resp, httpRes, err := apiClient.SiteAPI.UpdateSiteConfig(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
