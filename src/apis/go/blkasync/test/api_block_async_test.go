/*
OPERA Cloud Block Reservation Asynchronous API

Testing BlockAsyncApiService

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

func Test_openapi_BlockAsyncApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlockAsyncApiService GetBlockAllocationSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string
		var hotelId string
		var extSystemCode string

		resp, httpRes, err := apiClient.BlockAsyncApi.GetBlockAllocationSummary(context.Background(), requestId, hotelId, extSystemCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAsyncApiService GetBlockAllocationSummaryProcessStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string
		var extSystemCode string
		var hotelId string

		httpRes, err := apiClient.BlockAsyncApi.GetBlockAllocationSummaryProcessStatus(context.Background(), requestId, extSystemCode, hotelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAsyncApiService StartBlockAllocationSummaryProcess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string
		var extSystemCode string

		httpRes, err := apiClient.BlockAsyncApi.StartBlockAllocationSummaryProcess(context.Background(), hotelId, extSystemCode).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
