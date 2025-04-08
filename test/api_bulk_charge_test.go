/*
Paystack

Testing BulkChargeAPIService

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

func Test_openapi_BulkChargeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BulkChargeAPIService BulkChargeCharges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var code string

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargeCharges(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BulkChargeAPIService BulkChargeFetch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var code string

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargeFetch(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BulkChargeAPIService BulkChargeInitiate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargeInitiate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BulkChargeAPIService BulkChargeList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargeList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BulkChargeAPIService BulkChargePause", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var code string

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargePause(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BulkChargeAPIService BulkChargeResume", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var code string

		resp, httpRes, err := apiClient.BulkChargeAPI.BulkChargeResume(context.Background(), code).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
