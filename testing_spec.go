package main

import (
	"testing"
)

func TestPublishingService(t *testing.T) {
	t.Parallel()
	// START OMIT

	t.Run(`
		GIVEN a content studio draft article
		WHEN the article is published
		THEN the article is available for the fans
	`, func(t *testing.T) {

		articleID := ids.GenerateArticleID()
		sut := dsl.NewSystemUnderTest(t, dsl.SutWithEditedArticle(articleID))
		defer sut.Close()

		sut.PublishArticle(t.Context(),
			dsl.ArticleWithID(articleID),
			dsl.ArticleWithDefaultAuthentication())

		sut.AssertArticlePublished(t.Context(), articleID)
	})
	// END OMIT

}
