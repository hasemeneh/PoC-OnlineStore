package cart

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
)

func (c *cartDomain) AddToCart(ctx context.Context, req *models.AddToCartRequest) (*models.AddToCartResponse, error) {
	resp := &models.AddToCartResponse{}
	temp, err := c.getUserCart(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	temp[fmt.Sprintf("%d", req.ProductID)] = req.ProductRequest
	byteData, _ := json.Marshal(temp)
	_, err = c.redis.Set(fmt.Sprintf(constants.UserCartCacheKey, req.UserID), string(byteData), constants.DefaultRedisTTL*time.Second)
	for _, product := range temp {
		resp.Products = append(resp.Products, product)
	}
	return resp, err
}

func (c *cartDomain) getUserCart(ctx context.Context, userID int64) (map[string]models.ProductRequest, error) {
	resp := make(map[string]models.ProductRequest)
	data, err := c.redis.Get(fmt.Sprintf(constants.UserCartCacheKey, userID))
	if err != nil {
		if err == redis.Nil {
			return resp, nil
		}
		return resp, err
	}
	byteData := []byte(data)
	err = json.Unmarshal(byteData, &resp)
	if err != nil {
		log.Println("[getUserCart] Failed to unmarshall user cart because of ", err)
	}
	return resp, nil
}

func (c *cartDomain) GetUserCart(ctx context.Context, userID int64) (*models.AddToCartResponse, error) {
	resp := &models.AddToCartResponse{
		Products: make([]models.ProductRequest, 0),
	}
	temp, err := c.getUserCart(ctx, userID)
	if err != nil {
		return resp, err
	}
	for _, product := range temp {
		resp.Products = append(resp.Products, product)
	}
	return resp, nil
}
