/*
MetalSoft REST API

Testing VMInstanceGroupAPIService

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

func Test_sdk_VMInstanceGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VMInstanceGroupAPIService ApplyVMTypeOnVMInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32
		var vmTypeId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.ApplyVMTypeOnVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId, vmTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService CreateVMInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.CreateVMInstanceGroup(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService CreateVMInterfaceOnVMInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.CreateVMInterfaceOnVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService DeleteVMInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		httpRes, err := apiClient.VMInstanceGroupAPI.DeleteVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetInfrastructureVMInstanceGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetInfrastructureVMInstanceGroup(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetInfrastructureVMInstanceGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetInfrastructureVMInstanceGroups(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetVMInstanceGroupConfigInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupConfigInfo(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetVMInstanceGroupInterfaceInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32
		var vmInstanceGroupInterfaceId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupInterfaceInfo(context.Background(), infrastructureId, vmInstanceGroupId, vmInstanceGroupInterfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetVMInstanceGroupInterfaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupInterfaces(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService GetVMInstanceGroupVMInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.GetVMInstanceGroupVMInstances(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService PatchVMInstanceGroupMeta", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.PatchVMInstanceGroupMeta(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VMInstanceGroupAPIService UpdateVMInstanceGroupConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var vmInstanceGroupId float32

		resp, httpRes, err := apiClient.VMInstanceGroupAPI.UpdateVMInstanceGroupConfig(context.Background(), infrastructureId, vmInstanceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
