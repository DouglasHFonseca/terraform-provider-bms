/*
AdServer Ad Management

Testing ExchangeReviewsAPIService

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

func Test_openapi_ExchangeReviewsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExchangeReviewsAPIService PutExchangeReview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var adId string
		var exchangeId string

		httpRes, err := apiClient.ExchangeReviewsAPI.PutExchangeReview(context.Background(), accountId, adId, exchangeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExchangeReviewsAPIService RefreshExchangeReviewStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var adId string

		httpRes, err := apiClient.ExchangeReviewsAPI.RefreshExchangeReviewStatus(context.Background(), accountId, adId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExchangeReviewsAPIService SendAdForExchangeReview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var adId string

		httpRes, err := apiClient.ExchangeReviewsAPI.SendAdForExchangeReview(context.Background(), accountId, adId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
