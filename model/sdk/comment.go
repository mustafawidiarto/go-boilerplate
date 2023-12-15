package sdk

import (
	"strings"
	"time"
)

type Comments []Comment
type Comment struct {
	ID           int64          `json:"id"`
	RoomID       int64          `json:"room_id"`
	Email        string         `json:"email"`
	IDStr        string         `json:"id_str"`
	Message      string         `json:"message"`
	RoomIDStr    string         `json:"room_id_str"`
	RoomName     string         `json:"room_name"`
	Timestamp    string         `json:"timestamp"`
	Username     string         `json:"username"`
	UniqueTempID string         `json:"unique_temp_id"`
	UserExtras   *UserExtras    `json:"user_extras"`
	Extras       *CommentExtras `json:"extras"`
	Payload      interface{}    `json:"payload"`
	IsDeleted    bool           `json:"is_deleted"`

	CommentBeforeID    int64  `json:"comment_before_id"`
	CommentBeforeIDStr string `json:"comment_before_id_str"`
	CreatedAt          string `json:"created_at"`
	DisableLinkPreview bool   `json:"disable_link_preview"`
	Text               string `json:"text"`
	Type               string `json:"type"`
	UnixNanoTimestamp  int64  `json:"unix_nano_timestamp"`
	UnixTimestamp      int64  `json:"unix_timestamp"`
}

type UserExtras struct {
	IsCustomer bool   `json:"is_customer"`
	Type       string `json:"type"`
}

type CommentExtras struct {
	Order  *Order      `json:"order"`
	Action interface{} `json:"action"`
}

type Order struct {
	Text         string       `json:"text"`
	CatalogID    string       `json:"catalog_id"`
	ProductItems ProductItems `json:"product_items"`
}

type ProductItems []ProductItem
type ProductItem struct {
	Currency          string `json:"currency"`
	ItemPrice         int64  `json:"item_price"`
	ProductRetailerID string `json:"product_retailer_id"`
	Quantity          int    `json:"quantity"`
}

func (c *Comment) GetTimestamp() (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", c.Timestamp)
}

func (c *Comment) IsWhatsappOrderMessage() bool {
	return strings.EqualFold(c.Message, "Order message from Whatsapp catalog")
}

func (cs Comments) IsLastMessageFromCustomer() bool {
	if len(cs) == 0 {
		return true
	}

	for _, v := range cs {
		if strings.ToLower(v.Username) == "system" {
			continue
		}
		if v.SenderIsBot() {
			continue
		}

		return v.RoomName == v.Username
	}

	return true
}

func (cs Comments) GetLastMessageFromCustomer() *Comment {
	for _, v := range cs {
		if v.RoomName == v.Username {
			return &v
		}
	}

	return nil
}

func (cs Comments) GetLastMessageFromAgent() *Comment {
	for _, v := range cs {
		if strings.ToLower(v.Username) == "system" {
			continue
		}

		if v.RoomName != v.Username {
			return &v
		}
	}

	return nil
}

func (cs Comments) GetLastOrderMessageFromCustomer() *Comment {
	for _, v := range cs {
		if v.IsMessageFromCustomer() && v.IsWhatsappOrderMessage() {
			return &v
		}
	}

	return nil
}

func (c Comment) IsMessageFromCustomer() bool {
	return strings.EqualFold(c.RoomName, c.Username)
}

func (c *Comment) SenderIsBot() bool {
	if c.Extras != nil {
		if action, ok := c.Extras.Action.(string); ok {
			if action == "bot_reply" {
				return true
			}
		}
	}
	return false
}
