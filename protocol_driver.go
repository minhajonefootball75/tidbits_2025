func (d *ProtocolDriver) CreateArticle(ctx context.Context, language language.Code, integrationID int) ids.ID {
	t := d.T

	e := createArticleRequest{Language: language}
	url := fmt.Sprintf("%s/v1/articles", d.networkAPIURL)
	b := io.NopCloser(bytes.NewReader(mustMarshal(e)))
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, b)

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	var response createArticleResponse
	_ = json.NewDecoder(resp.Body).Decode(&response)

	return response.ID
}