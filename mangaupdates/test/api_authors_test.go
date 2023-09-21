/*
MangaUpdates API

Testing AuthorsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/belphemur/mangal/mangaupdates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_AuthorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthorsAPIService AddAuthor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthorsAPI.AddAuthor(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService DeleteAuthor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.DeleteAuthor(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService DeleteImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.DeleteImage(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService LockAuthorField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64
		var item string

		resp, httpRes, err := apiClient.AuthorsAPI.LockAuthorField(context.Background(), id, item).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService RetrieveAuthor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.RetrieveAuthor(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService RetrieveAuthorLocks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.RetrieveAuthorLocks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService RetrieveAuthorSeries", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.RetrieveAuthorSeries(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService SearchAuthorsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthorsAPI.SearchAuthorsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService UnlockAuthorField", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64
		var item string

		resp, httpRes, err := apiClient.AuthorsAPI.UnlockAuthorField(context.Background(), id, item).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService UpdateAuthor", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.UpdateAuthor(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorsAPIService UpdateImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.AuthorsAPI.UpdateImage(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
