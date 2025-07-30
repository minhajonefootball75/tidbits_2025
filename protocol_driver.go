func (d *ProtocolDriver) PublishArticle(ctx context.Context, articleID ids.ID, body ArticlePublishRequest, token string) {
	t := d.T

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("%s/v1/articles/%s", d.networkAPIURL, articleID.String()), io.NopCloser(bytes.NewReader(mustMarshal(body))))
	require.NoError(t, err)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	require.NoError(t, err, "failed to publish article")
	require.Equal(t, http.StatusOK, resp.StatusCode, "unexpected status code for publish article")
}
