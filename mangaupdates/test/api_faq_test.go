/*
MangaUpdates API

Testing FaqAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/sergiolaverde0/mangal/mangaupdates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_FaqAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FaqAPIService AddFaqCategory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FaqAPI.AddFaqCategory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService AddFaqQuestion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FaqAPI.AddFaqQuestion(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService DeleteFaqCategory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FaqAPI.DeleteFaqCategory(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService DeleteFaqQuestion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64
		var questionId int64

		resp, httpRes, err := apiClient.FaqAPI.DeleteFaqQuestion(context.Background(), categoryId, questionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService ReorderFaq", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FaqAPI.ReorderFaq(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService RetrieveAllFaqCategoriesAndQuestions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FaqAPI.RetrieveAllFaqCategoriesAndQuestions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService RetrieveAllFaqCategoryQuestions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FaqAPI.RetrieveAllFaqCategoryQuestions(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService RetrieveFaqCategory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FaqAPI.RetrieveFaqCategory(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService RetrieveFaqQuestion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64
		var questionId int64

		resp, httpRes, err := apiClient.FaqAPI.RetrieveFaqQuestion(context.Background(), categoryId, questionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService UpdateFaqCategory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FaqAPI.UpdateFaqCategory(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FaqAPIService UpdateFaqQuestion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId int64
		var questionId int64

		resp, httpRes, err := apiClient.FaqAPI.UpdateFaqQuestion(context.Background(), categoryId, questionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
