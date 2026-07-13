package startupaid

import (
	"context"
	"net/http"
	"time"
)

// Channel is a connected social channel.
type Channel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Active   bool   `json:"active"`
}

// ListChannels returns the account's connected social channels. Use the ids with
// SchedulePost.
func (c *Client) ListChannels(ctx context.Context) ([]Channel, error) {
	var out struct {
		Channels []Channel `json:"channels"`
	}
	if err := c.do(ctx, http.MethodGet, "/v1/social/channels", nil, nil, &out); err != nil {
		return nil, err
	}
	return out.Channels, nil
}

// SchedulePostRequest schedules or publishes a social post. Leave ScheduledFor
// nil to publish immediately.
type SchedulePostRequest struct {
	Channels     []string   `json:"channels"`
	Content      string     `json:"content"`
	ImageURL     string     `json:"imageUrl,omitempty"`
	ScheduledFor *time.Time `json:"scheduledFor,omitempty"`
}

// Post is a scheduled or published social post.
type Post struct {
	ID           string `json:"id"`
	Channels     string `json:"channels"`
	Content      string `json:"content"`
	ImageURL     string `json:"imageUrl,omitempty"`
	Status       string `json:"status"`
	ScheduledFor string `json:"scheduledFor,omitempty"`
	PublishedAt  string `json:"publishedAt,omitempty"`
	Error        string `json:"error,omitempty"`
}

// SchedulePost schedules or immediately publishes a post to one or more channels.
func (c *Client) SchedulePost(ctx context.Context, req SchedulePostRequest) (*Post, error) {
	var p Post
	if err := c.do(ctx, http.MethodPost, "/v1/social/schedule", nil, req, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ListPosts returns the account's social posts with delivery status.
func (c *Client) ListPosts(ctx context.Context) ([]Post, error) {
	var out struct {
		Posts []Post `json:"posts"`
	}
	if err := c.do(ctx, http.MethodGet, "/v1/social/posts", nil, nil, &out); err != nil {
		return nil, err
	}
	return out.Posts, nil
}
