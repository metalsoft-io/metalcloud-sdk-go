/*
MetalSoft REST API

Testing VMTypeAPIService

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

func Test_sdk_VMTypeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VMTypeAPIService CreateVMType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VMTypeAPI.CreateVMType(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMTypeAPIService DeleteVMType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vmTypeId float32

		httpRes, err := apiClient.VMTypeAPI.DeleteVMType(context.Background(), vmTypeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMTypeAPIService GetVMType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vmTypeId float32

		resp, httpRes, err := apiClient.VMTypeAPI.GetVMType(context.Background(), vmTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMTypeAPIService GetVMTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VMTypeAPI.GetVMTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMTypeAPIService GetVMsByVMType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vmTypeId float32

		resp, httpRes, err := apiClient.VMTypeAPI.GetVMsByVMType(context.Background(), vmTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMTypeAPIService UpdateVMType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vmTypeId float32

		resp, httpRes, err := apiClient.VMTypeAPI.UpdateVMType(context.Background(), vmTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
