package monday

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Item(query *models.ItemQuery) (out []models.Item, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: mondayAPI,
		In:      query,
		Out:     &ResponseData{},
		Params:  nil,
	}

	if err = c.api.callMethod(options); err != nil {
		return
	}
	out = options.Out.(*ResponseData).Data.Item
	return
}
