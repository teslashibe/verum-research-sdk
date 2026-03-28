package contribute

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/teslashibe/verum-research-sdk/anonymize"
)

type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Submit(ctx context.Context, contribution *anonymize.AnonymizedContribution) error {
	body, err := json.Marshal(SubmitRequest{Contribution: *contribution})
	if err != nil {
		return fmt.Errorf("marshal contribution: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/v1/contribute", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("submit contribution: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("research API %d: %s", resp.StatusCode, string(respBody))
	}
	return nil
}

func (c *Client) GetPlatformStats(ctx context.Context) (*PlatformStats, error) {
	return doGet[PlatformStats](ctx, c, "/v1/stats")
}

func (c *Client) GetCompoundFindings(ctx context.Context, compoundID string) ([]Finding, error) {
	result, err := doGet[[]Finding](ctx, c, "/v1/compounds/"+compoundID+"/findings")
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) GetReport(ctx context.Context, slug string) (*Report, error) {
	return doGet[Report](ctx, c, "/v1/reports/"+slug)
}

func (c *Client) SearchReports(ctx context.Context, query string) ([]ReportSummary, error) {
	result, err := doGet[[]ReportSummary](ctx, c, "/v1/reports?q="+query)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func doGet[T any](ctx context.Context, c *Client, path string) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request %s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("research API %d: %s", resp.StatusCode, string(body))
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &result, nil
}
