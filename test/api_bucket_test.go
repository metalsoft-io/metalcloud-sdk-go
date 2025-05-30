/*
MetalSoft REST API

Testing BucketAPIService

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

func Test_sdk_BucketAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BucketAPIService CreateInfrastructureBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.BucketAPI.CreateInfrastructureBucket(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService DeleteBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		httpRes, err := apiClient.BucketAPI.DeleteBucket(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService GetBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.GetBucket(context.Background(), bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService GetBucketConfigInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.GetBucketConfigInfo(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService GetBucketCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.GetBucketCredentials(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService GetInfrastructureBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.GetInfrastructureBucket(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService GetInfrastructureBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32

		resp, httpRes, err := apiClient.BucketAPI.GetInfrastructureBuckets(context.Background(), infrastructureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService UpdateBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.UpdateBucket(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService UpdateBucketMeta", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var infrastructureId float32
		var bucketId float32

		resp, httpRes, err := apiClient.BucketAPI.UpdateBucketMeta(context.Background(), infrastructureId, bucketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
