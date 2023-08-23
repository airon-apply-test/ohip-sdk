/*
OPERA Cloud Front Desk Operations Service

Testing FOFStatsApiService

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

func Test_openapi_FOFStatsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FOFStatsApiService GetFrontOfficeStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var statisticsCode string
		var hotelId string

		resp, httpRes, err := apiClient.FOFStatsApi.GetFrontOfficeStatistics(context.Background(), statisticsCode, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FOFStatsApiService GetFrontOfficeStatisticsWithDateRange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var statisticsCode string
		var hotelId string

		resp, httpRes, err := apiClient.FOFStatsApi.GetFrontOfficeStatisticsWithDateRange(context.Background(), statisticsCode, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FOFStatsApiService GetReservationQueueStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.FOFStatsApi.GetReservationQueueStatistics(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FOFStatsApiService GetTaskSheetStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.FOFStatsApi.GetTaskSheetStatistics(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
