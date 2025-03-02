/*
MangaUpdates API

Testing GenreAPIService

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

func Test_openapi_GenreAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenreAPIService AddGenre", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GenreAPI.AddGenre(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenreAPIService DeleteGenre", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.GenreAPI.DeleteGenre(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenreAPIService RetrieveGenreById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.GenreAPI.RetrieveGenreById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenreAPIService RetrieveGenres", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GenreAPI.RetrieveGenres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenreAPIService UpdateGenre", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.GenreAPI.UpdateGenre(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
