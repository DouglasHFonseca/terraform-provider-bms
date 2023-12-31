/*
AdServer Creative Management

Testing CreativeGroupsAPIService

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

func Test_openapi_CreativeGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CreativeGroupsAPIService CreateCreativeGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.CreativeGroupsAPI.CreateCreativeGroup(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreativeGroupsAPIService DeleteCreativeGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var creativeGroupId string

		httpRes, err := apiClient.CreativeGroupsAPI.DeleteCreativeGroup(context.Background(), accountId, creativeGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreativeGroupsAPIService GetCreativeGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var creativeGroupId string

		resp, httpRes, err := apiClient.CreativeGroupsAPI.GetCreativeGroup(context.Background(), accountId, creativeGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreativeGroupsAPIService GetCreativeGroupAvailableCreatives", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var creativeGroupId string

		resp, httpRes, err := apiClient.CreativeGroupsAPI.GetCreativeGroupAvailableCreatives(context.Background(), accountId, creativeGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreativeGroupsAPIService ListCreativeGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.CreativeGroupsAPI.ListCreativeGroups(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CreativeGroupsAPIService PatchCreativeGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var creativeGroupId string

		resp, httpRes, err := apiClient.CreativeGroupsAPI.PatchCreativeGroup(context.Background(), accountId, creativeGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
