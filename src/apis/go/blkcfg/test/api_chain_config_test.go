/*
OPERA Cloud Block Configuration API

Testing ChainConfigApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ChainConfigApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChainConfigApiService DeleteBlockCancellationReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockCancellationReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteBlockCancellationReasons(context.Background(), blockCancellationReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteBlockLostBookingCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockLostBookingCodesId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteBlockLostBookingCodes(context.Background(), blockLostBookingCodesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteBlockRateOverrideReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockRateOverrideReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteBlockRateOverrideReasons(context.Background(), blockRateOverrideReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteBlockRefusedReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockRefusedReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteBlockRefusedReasons(context.Background(), blockRefusedReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteChainConfigServicesCache", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteChainConfigServicesCache(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteDestinationCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var destinationCodesId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteDestinationCodes(context.Background(), destinationCodesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService DeleteReservationMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationMethodsId string

		resp, httpRes, err := apiClient.ChainConfigApi.DeleteReservationMethods(context.Background(), reservationMethodsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetBlockCancellationReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetBlockCancellationReasons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetBlockLostBookingCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetBlockLostBookingCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetBlockRateOverrideReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetBlockRateOverrideReasons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetBlockRefusedReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetBlockRefusedReasons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetDestinationCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetDestinationCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService GetReservationMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.GetReservationMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PingChainConfigServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PingChainConfigServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostBlockCancellationReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostBlockCancellationReasons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostBlockLostBookingCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostBlockLostBookingCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostBlockRateOverrideReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostBlockRateOverrideReasons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostBlockRefusedReasonspostBlo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostBlockRefusedReasonspostBlo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostDestinationCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostDestinationCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PostReservationMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChainConfigApi.PostReservationMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutBlockCancellationReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockCancellationReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutBlockCancellationReasons(context.Background(), blockCancellationReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutBlockLostBookingCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockLostBookingCodesId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutBlockLostBookingCodes(context.Background(), blockLostBookingCodesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutBlockRateOverrideReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockRateOverrideReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutBlockRateOverrideReasons(context.Background(), blockRateOverrideReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutBlockRefusedReasons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockRefusedReasonsId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutBlockRefusedReasons(context.Background(), blockRefusedReasonsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutDestinationCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var destinationCodesId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutDestinationCodes(context.Background(), destinationCodesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChainConfigApiService PutReservationMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationMethodsId string

		resp, httpRes, err := apiClient.ChainConfigApi.PutReservationMethods(context.Background(), reservationMethodsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
