package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type T struct {
	Id    int    `json:"id"`
	Pic   string `json:"pic"`
	Title string `json:"title"`
}

func BannerData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": []T{
			{
				Id:    1,
				Pic:   "https://m.360buyimg.com/babel/s1071x321_jfs/t20270109/235883/3/12114/47052/659e3034F21c8d0b7/7b7abb4f8cc2627c.jpg!q70.dpg",
				Title: "",
			}, {
				Id:    2,
				Pic:   "https://m.360buyimg.com/babel/jfs/t1/232236/36/11109/182790/6597bae7Ff23261c1/11f54ec6423324ae.png",
				Title: "",
			}, {
				Id:    3,
				Pic:   "https://m.360buyimg.com/babel/jfs/t20270102/245832/19/1737/93109/6594dd77F6bf121b2/8c8bbd22ed90d51f.png",
				Title: "",
			},
		},
	})
}
