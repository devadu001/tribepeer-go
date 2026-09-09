package tribepeer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultPartner = "https://tribepeer.com/api/partner/v1"
	DefaultProduct = "https://tribepeer.com/api/product/v1"
)

type Client struct {
	ClientID        string
	ClientSecret    string
	PartnerBase     string
	ProductBase     string
	InstitutionUUID string
	HTTP            *http.Client

	partnerToken string
	userToken    string

	Users       *Users
	Tribes      *Tribes
	Materials   *Materials
	Submissions *Submissions
	Communities *Communities
	AI          *AI
	Campus      *Campus
}

func New(clientID, clientSecret string) *Client {
	c := &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		PartnerBase:  strings.TrimRight(DefaultPartner, "/"),
		ProductBase:  strings.TrimRight(DefaultProduct, "/"),
		HTTP:         &http.Client{Timeout: 30 * time.Second},
	}
	c.Users = &Users{c: c}
	c.Tribes = &Tribes{c: c}
	c.Materials = &Materials{c: c}
	c.Submissions = &Submissions{c: c}
	c.Communities = &Communities{c: c}
	c.AI = &AI{c: c}
	c.Campus = &Campus{c: c}
	return c
}

func (c *Client) Token(ctx context.Context) (map[string]any, error) {
	if c.ClientID == "" || c.ClientSecret == "" {
		return nil, fmt.Errorf("clientId and clientSecret are required for partner auth")
	}
	data, err := c.request(ctx, http.MethodPost, c.PartnerBase+"/auth/token", map[string]any{
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	}, "")
	if err != nil {
		return nil, err
	}
	if tok, ok := data["access_token"].(string); ok {
		c.partnerToken = tok
	}
	return data, nil
}

func (c *Client) PartnerToken() string { return c.partnerToken }
func (c *Client) UserToken() string    { return c.userToken }

func (c *Client) SetUserToken(token string) { c.userToken = token }

func (c *Client) Me(ctx context.Context) (map[string]any, error) {
	return c.partner(ctx, http.MethodGet, "/me", nil)
}

func (c *Client) UpdateBranding(ctx context.Context, body map[string]any) (map[string]any, error) {
	return c.partner(ctx, http.MethodPatch, "/me/branding", body)
}

func (c *Client) partner(ctx context.Context, method, path string, body map[string]any) (map[string]any, error) {
	if c.partnerToken == "" {
		if _, err := c.Token(ctx); err != nil {
			return nil, err
		}
	}
	return c.request(ctx, method, c.PartnerBase+path, body, c.partnerToken)
}

func (c *Client) product(ctx context.Context, method, path string, body map[string]any) (map[string]any, error) {
	return c.request(ctx, method, c.ProductBase+path, body, c.userToken)
}

func (c *Client) request(ctx context.Context, method, href string, body map[string]any, token string) (map[string]any, error) {
	var rdr io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(raw)
	}

	req, err := http.NewRequestWithContext(ctx, method, href, rdr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if c.InstitutionUUID != "" {
		req.Header.Set("X-Institution-Uuid", c.InstitutionUUID)
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := map[string]any{}
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &data); err != nil {
			data = map[string]any{"message": string(raw)}
		}
	}

	if res.StatusCode >= 400 {
		msg := firstString(data, "message", "error")
		if msg == "" {
			msg = fmt.Sprintf("TribePeer request failed (%d)", res.StatusCode)
		}
		if res.StatusCode == 402 {
			upgrade, _ := data["upgrade"].(map[string]any)
			return nil, &PaymentRequiredError{Status: 402, Message: msg, Body: data, Upgrade: upgrade}
		}
		return nil, &Error{Status: res.StatusCode, Message: msg, Body: data}
	}

	return data, nil
}

func firstString(m map[string]any, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k].(string); ok && v != "" {
			return v
		}
	}
	return ""
}

func withQuery(href string, q map[string]string) string {
	vals := url.Values{}
	for k, v := range q {
		if v != "" {
			vals.Set(k, v)
		}
	}
	if encoded := vals.Encode(); encoded != "" {
		if strings.Contains(href, "?") {
			return href + "&" + encoded
		}
		return href + "?" + encoded
	}
	return href
}
