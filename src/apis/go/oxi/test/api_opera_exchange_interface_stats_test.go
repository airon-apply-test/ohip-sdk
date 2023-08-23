/*
OPERA Cloud Xchange Interface OXI API

Testing OperaExchangeInterfaceStatsApiService

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

func Test_openapi_OperaExchangeInterfaceStatsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OperaExchangeInterfaceStatsApiService DequeueOXIMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var queueName string
		var hotelId string
		var interfaceId string

		resp, httpRes, err := apiClient.OperaExchangeInterfaceStatsApi.DequeueOXIMessages(context.Background(), queueName, hotelId, interfaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperaExchangeInterfaceStatsApiService GetMessageStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.OperaExchangeInterfaceStatsApi.GetMessageStatistics(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OperaExchangeInterfaceStatsApiService GetOXIMessageStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var queueName string
		var hotelId string

		resp, httpRes, err := apiClient.OperaExchangeInterfaceStatsApi.GetOXIMessageStatistics(context.Background(), queueName, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
