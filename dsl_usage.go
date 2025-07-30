func (s *SystemUnderTest) PublishArticle(ctx context.Context, opts ...articleOption) {
	// driver
	articlePublication := article{
		articleID:    ids.GenerateArticleID(),
		language:     defaultLanguage,
		title:        defaultEnglishArticleTitle,
		content:      defaultEnglishArticleBlockContent,
		imageURL:     defaultImageURL,
		authorID:     defaultAuthorAuth0ID,
		authorHidden: defaultAuthorHidden,
	}
	for _, opt := range opts {
		opt(&articlePublication)
	}
	req := buildRequest(articlePublication)
	s.protocolDriver.PublishArticle(ctx, articlePublication.articleID, req)
}
