package pubgo

import (
	// "admin/models"
	// "fmt"
	"math"
)

type PageResult struct {
	Pages    []int
	Nums     int
	Previous int
	Next     int
	Last     int
}

func Pages(data int, pageTotal int, pageNum int, nowPage int) (pages []int, nums int, previous int, next int, last int) {
	nums = data
	last = int(math.Ceil(float64(nums) / float64(pageTotal)))
	next = nowPage + 1
	previous = nowPage
	if nowPage != 1 {
		previous = nowPage - 1
	}
	if next >= last {
		next = last
	}
	if pageHave := int(math.Ceil(float64(nums)/float64(pageTotal))) - nowPage; pageHave >= 8 {
		for i := 0; i <= pageNum; i++ {
			pages = append(pages, nowPage+i)
		}
	} else {
		for i := 0; i <= pageHave; i++ {
			pages = append(pages, nowPage+i)
		}
	}
	return
}


func PagesJson(data int, pageTotal int, pageNum int, nowPage int) (pageJson PageResult) {
	pageJson.Nums = data
	pageJson.Last = int(math.Ceil(float64(pageJson.Nums) / float64(pageTotal)))
	pageJson.Next = nowPage + 1
	pageJson.Previous = nowPage
	if nowPage != 1 {
		pageJson.Previous = nowPage - 1
	}
	if pageJson.Next >= pageJson.Last {
		pageJson.Next = pageJson.Last
	}
	if pageHave := int(math.Ceil(float64(pageJson.Nums)/float64(pageTotal))) - nowPage; pageHave >= 8 {
		for i := 0; i <= pageNum; i++ {
			pageJson.Pages = append(pageJson.Pages , nowPage+i)
		}
	} else {
		for i := 0; i <= pageHave; i++ {
			pageJson.Pages  = append(pageJson.Pages , nowPage+i)
		}
	}
	return
}
