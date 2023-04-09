//go:build v2
// +build v2

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
}

var products = map[int]Product{
	1: {
		ProductID:   1,
		ProductName: "Apple",
	},
	2: {
		ProductID:   2,
		ProductName: "Huawei",
	},
	3: {
		ProductID:   3,
		ProductName: "Samsung",
	},
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		productID := c.Query("productId")
		if productID == "" {
			c.JSON(400, gin.H{
				"message": "productId is required",
			})
			return
		}
		id, err := strconv.Atoi(productID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid productId",
			})
			return
		}

		c.JSON(200, products[id])
	})

	r.GET("/product", func(c *gin.Context) {
		productID := c.Query("productId")
		if productID == "" {
			c.JSON(400, gin.H{
				"message": "productId is required",
			})
			return
		}
		id, err := strconv.Atoi(productID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid productId",
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"data": map[string]interface{}{
				"info": products[id],
			},
		})
	})

	return r
}
