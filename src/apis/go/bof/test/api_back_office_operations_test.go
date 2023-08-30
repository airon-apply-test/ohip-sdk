/*
OPERA Cloud Back Office Operations API

Testing BackOfficeOperationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package bof

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/applyinnovations/ohip-sdk/bof"
)

func Test_bof_BackOfficeOperationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BackOfficeOperationsApiService ClearBackOfficeOperationsServiceCache", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BackOfficeOperationsApi.ClearBackOfficeOperationsServiceCache(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackOfficeOperationsApiService GetBusinessDate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.BackOfficeOperationsApi.GetBusinessDate(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackOfficeOperationsApiService PingBackOfficeOperationsService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BackOfficeOperationsApi.PingBackOfficeOperationsService(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
