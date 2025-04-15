/*
MetalSoft REST API

Testing UsersAPIService

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

func Test_sdk_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService AddUserDelegate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32
		var delegateId float32

		resp, httpRes, err := apiClient.UsersAPI.AddUserDelegate(context.Background(), userId, delegateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AddUserSshKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.AddUserSshKey(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ArchiveUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.ArchiveUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ChangeUserAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.ChangeUserAccount(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUserAuthorized", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.CreateUserAuthorized(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUserSshKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32
		var keyId float32

		httpRes, err := apiClient.UsersAPI.DeleteUserSshKey(context.Background(), userId, keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserChildDelegates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserChildDelegates(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserConfiguration(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserLimits", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserLimits(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserParentDelegates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserParentDelegates(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserPermissions(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserSshKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32
		var keyId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserSshKey(context.Background(), userId, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserSshKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserSshKeys(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserSuspendReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.GetUserSuspendReasons(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService RemoveUserDelegate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32
		var delegateId float32

		resp, httpRes, err := apiClient.UsersAPI.RemoveUserDelegate(context.Background(), userId, delegateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService SuspendUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.SuspendUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UnarchiveUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.UnarchiveUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UnsuspendUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		httpRes, err := apiClient.UsersAPI.UnsuspendUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.UpdateUserConfig(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserLimits", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.UpdateUserLimits(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserMeta", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.UpdateUserMeta(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId float32

		resp, httpRes, err := apiClient.UsersAPI.UpdateUserPermissions(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
