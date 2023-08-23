/*
OPERA Cloud Block Configuration API

Testing BlockConfigApiService

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

func Test_openapi_BlockConfigApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlockConfigApiService DeleteBlockConfigServiceCache", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.DeleteBlockConfigServiceCache(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService DeleteBlockSalesAllowanceRange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.DeleteBlockSalesAllowanceRange(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService DeleteBlockStatusCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockStatusCode string

		resp, httpRes, err := apiClient.BlockConfigApi.DeleteBlockStatusCode(context.Background(), blockStatusCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService DeleteWashSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var washScheduleCode string
		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.DeleteWashSchedule(context.Background(), washScheduleCode, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetBlockSalesAllowance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.GetBlockSalesAllowance(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetBlockStatusCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.GetBlockStatusCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetNextBlockStatusCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.GetNextBlockStatusCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetSalesManagerGoals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesManagerId string
		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.GetSalesManagerGoals(context.Background(), salesManagerId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetSalesManagerGoalsMultipleHotelIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesManagerId string

		resp, httpRes, err := apiClient.BlockConfigApi.GetSalesManagerGoalsMultipleHotelIds(context.Background(), salesManagerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetSalesManagers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.GetSalesManagers(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetSalesManagersMultipleHotelIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.GetSalesManagersMultipleHotelIds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService GetWashSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.GetWashSchedule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PingBlockConfigService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.PingBlockConfigService(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PostBlockStatusCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockConfigApi.PostBlockStatusCode(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PostSalesManagerGoals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesManagerId string
		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.PostSalesManagerGoals(context.Background(), salesManagerId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PostWashSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.PostWashSchedule(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PutBlockStatusCode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var blockStatusCode string

		resp, httpRes, err := apiClient.BlockConfigApi.PutBlockStatusCode(context.Background(), blockStatusCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PutNextBlockStatusCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var srcBlockStatusCode string

		resp, httpRes, err := apiClient.BlockConfigApi.PutNextBlockStatusCodes(context.Background(), srcBlockStatusCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PutSalesManagerGoals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesGoalId string
		var salesManagerId string
		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.PutSalesManagerGoals(context.Background(), salesGoalId, salesManagerId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService PutWashSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var washScheduleCode string
		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.PutWashSchedule(context.Background(), washScheduleCode, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService RemoveSalesManagerGoal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesGoalId string
		var salesManagerId string

		resp, httpRes, err := apiClient.BlockConfigApi.RemoveSalesManagerGoal(context.Background(), salesGoalId, salesManagerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService RemoveSalesManagerGoals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var salesManagerId string

		resp, httpRes, err := apiClient.BlockConfigApi.RemoveSalesManagerGoals(context.Background(), salesManagerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockConfigApiService SetBlockSalesAllowanceRange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BlockConfigApi.SetBlockSalesAllowanceRange(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
