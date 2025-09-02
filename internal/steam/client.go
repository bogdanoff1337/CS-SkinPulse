package steam

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "time"
)

type Client struct {
    http *http.Client
}

func NewClient() *Client {
    return &Client{
        http: &http.Client{Timeout: 10 * time.Second},
    }
}

func (c *Client) FetchPrice(itemName string, currency int) (*PriceOverview, error) {
    escapedName := url.QueryEscape(itemName)
    apiURL := fmt.Sprintf(
        "https://steamcommunity.com/market/priceoverview/?appid=730&currency=%d&market_hash_name=%s",
        currency,
        escapedName,
    )

    resp, err := c.http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var po PriceOverview
    if err := json.NewDecoder(resp.Body).Decode(&po); err != nil {
        return nil, err
    }

    return &po, nil
}
