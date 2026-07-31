// Package gpgapi provides a minimal hand-written client for the Bitbucket
// Cloud GPG-keys endpoint, which is not covered by the auto-generated
// swagger client.
package gpgapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.bitbucket.org/2.0"

// GpgAccountKey represents a GPG public key on a Bitbucket account.
type GpgAccountKey struct {
	Type              string     `json:"type"`
	Name              string     `json:"name,omitempty"`
	KeyID             string     `json:"key_id,omitempty"`
	Fingerprint       string     `json:"fingerprint,omitempty"`
	ParentFingerprint string     `json:"parent_fingerprint,omitempty"`
	AddedOn           *time.Time `json:"added_on,omitempty"`
	CreatedOn         *time.Time `json:"created_on,omitempty"`
	ExpiresOn         *time.Time `json:"expires_on,omitempty"`
	LastUsed          *time.Time `json:"last_used,omitempty"`
}

type paginatedGpgKeys struct {
	Size   int32          `json:"size,omitempty"`
	Page   int32          `json:"page,omitempty"`
	Values []GpgAccountKey `json:"values,omitempty"`
	Next   string         `json:"next,omitempty"`
}

// Client calls the Bitbucket GPG-keys API on behalf of a single user.
type Client struct {
	httpClient *http.Client
	token      string
	userUUID   string
}

// New returns a Client that authenticates with the given OAuth access token
// and operates on the GPG keys of the user identified by userUUID (the
// Bitbucket account UUID, with or without curly braces).
func New(token, userUUID string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		token:      token,
		userUUID:   userUUID,
	}
}

// List returns all GPG keys on the user's account, following pagination.
func (c *Client) List() ([]GpgAccountKey, error) {
	url := fmt.Sprintf("%s/users/%s/gpg-keys", baseURL, c.userUUID)
	var out []GpgAccountKey
	for url != "" {
		var page paginatedGpgKeys
		next, err := c.getPage(url, &page)
		if err != nil {
			return nil, err
		}
		out = append(out, page.Values...)
		url = next
	}
	return out, nil
}

// Add uploads an ASCII-armoured GPG public key with a user-supplied name.
func (c *Client) Add(name, armoredKey string) (*GpgAccountKey, error) {
	url := fmt.Sprintf("%s/users/%s/gpg-keys", baseURL, c.userUUID)
	body := map[string]string{
		"key":  armoredKey,
		"name": name,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}
	var key GpgAccountKey
	if err := json.Unmarshal(data, &key); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &key, nil
}

// Delete removes the GPG key identified by fingerprint.
func (c *Client) Delete(fingerprint string) error {
	url := fmt.Sprintf("%s/users/%s/gpg-keys/%s", baseURL, c.userUUID, fingerprint)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		data, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}
	return nil
}

// getPage issues a GET for one page of results and returns the URL of the
// next page, or an empty string when there are no more pages.
func (c *Client) getPage(url string, page *paginatedGpgKeys) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %s: %s", resp.Status, string(data))
	}
	if err := json.Unmarshal(data, page); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	return page.Next, nil
}
