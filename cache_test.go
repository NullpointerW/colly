package main

import (
	"fmt"
	"testing"
	"colly/cache"
)

func TestCacheSerach(t *testing.T) {
	

	rs := cache.Rows.Search("书")
	for _, r := range rs {
		fmt.Printf("%#+v \n",r)
	}
     fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("分享各类电子书")
	for _, r := range rs {
		fmt.Printf("%#+v \n",r)
	}
	fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("这只是个测试数据")
	for _, r := range rs {
		fmt.Printf("%#+v \n",r)
	}
	fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("安全")
	for _, r := range rs {
		fmt.Printf("%#+v \n",r)
	}



}