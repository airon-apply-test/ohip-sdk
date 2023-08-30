/*
OPERA Cloud Price Availability Rate Async API

Testing AvailabilityAsyncApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package parasync

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/applyinnovations/ohip-sdk/parasync"
)

func Test_parasync_AvailabilityAsyncApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AvailabilityAsyncApiService GetRestrictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string
		var extSystemCode string
		var requestId string

		resp, httpRes, err := apiClient.AvailabilityAsyncApi.GetRestrictions(context.Background(), hotelId, extSystemCode, requestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailabilityAsyncApiService GetRestrictionsProcessStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId string
		var extSystemCode string
		var hotelId string

		httpRes, err := apiClient.AvailabilityAsyncApi.GetRestrictionsProcessStatus(context.Background(), requestId, extSystemCode, hotelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvailabilityAsyncApiService PostRestrictionsProcess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string
		var extSystemCode string

		resp, httpRes, err := apiClient.AvailabilityAsyncApi.PostRestrictionsProcess(context.Background(), hotelId, extSystemCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
