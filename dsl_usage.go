func SutWithEditedArticle(articleID ids.ID, opts ...articleOption) func(s *SystemUnderTest) {
	return func(s *SystemUnderTest) {
		article := article{
			articleID: articleID,
			title:     defaultEnglishArticleTitle,
			// .. other defaults
		}

		for _, opt := range opts {
			opt(&article)
		}

		s.dbProtocolDriver.CreateArticle(article)
		s.dbProtocolDriver.CreateArticleRevision(articleRevision{
			articleID: article.articleID,
			// .. other defaults
		})
	}
}