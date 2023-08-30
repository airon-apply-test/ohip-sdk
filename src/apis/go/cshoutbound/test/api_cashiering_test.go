/*
OPERA Cloud Cashiering API

Testing CashieringApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cshoutbound

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/applyinnovations/ohip-sdk/cshoutbound"
)

func Test_cshoutbound_CashieringApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CashieringApiService ApproveCompPostings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.ApproveCompPostings(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashieringApiService AuthorizeCompRedemptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var membershipId string
		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.AuthorizeCompRedemptions(context.Background(), membershipId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashieringApiService GetCompRedemptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var membershipId string
		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.GetCompRedemptions(context.Background(), membershipId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashieringApiService PostCompRedemptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var membershipId string
		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.PostCompRedemptions(context.Background(), membershipId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashieringApiService ReverseCompPostings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.ReverseCompPostings(context.Background(), hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CashieringApiService ReverseCompRedemptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var membershipId string
		var hotelId string

		resp, httpRes, err := apiClient.CashieringApi.ReverseCompRedemptions(context.Background(), membershipId, hotelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
