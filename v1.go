//go:build v1
// +build v1

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ProductNumber      int    `json:"productNumber"`
	ProductDescription string `json:"productDescription"`
}

var products = map[int]Product{
	1: {
		ProductNumber:      1,
		ProductDescription: "Apple",
	},
	2: {
		ProductNumber:      2,
		ProductDescription: "Huawei",
	},
	3: {
		ProductNumber:      3,
		ProductDescription: "Samsung",
	},
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		productNumber := c.Query("productNumber")
		if productNumber == "" {
			c.JSON(400, gin.H{
				"message": "productNumber is required",
			})
			return
		}

		id, err := strconv.Atoi(productNumber)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid productNumber",
			})
			return
		}

		c.JSON(200, products[id])
	})

	r.GET("/detail", func(c *gin.Context) {
		productNumber := c.Query("productNumber")
		if productNumber == "" {
			c.JSON(400, gin.H{
				"message": "productNumber is required",
			})
			return
		}

		id, err := strconv.Atoi(productNumber)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid productNumber",
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
