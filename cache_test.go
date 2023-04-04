package main

import (
	"colly/cache"
	"fmt"
	"testing"
)

func TestCacheSerach(t *testing.T) {
	rs := cache.Rows.Search("书")
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("分享各类电子书")
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("这只是个测试数据")
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")
	rs = cache.Rows.Search("安全")
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
}


func TestCacheRegSerach(t *testing.T){
	rs,err := cache.Rows.SearchWithRegexp("书")
	 if err!=nil{
		t.Fatalf("TestCacheRegSerach err :%v",err)
	 }
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")


	rs,err = cache.Rows.SearchWithRegexp(`.*我的.*总结.*`)
	 if err!=nil{
		t.Fatalf("TestCacheRegSerach err :%v",err)
	 }
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")


	rs,err = cache.Rows.SearchWithRegexp(`^我的.*`)
	 if err!=nil{
		t.Fatalf("TestCacheRegSerach err :%v",err)
	 }
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")

	rs,err = cache.Rows.SearchWithRegexp(`.*分享.*IT.*`)
	 if err!=nil{
		t.Fatalf("TestCacheRegSerach err :%v",err)
	 }
	for _, r := range rs {
		fmt.Printf("%#+v \n", r)
	}
	fmt.Println("-----------------------------------------")



	

}