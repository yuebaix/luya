package spider

import (
	"fmt"
	"github.com/yuebaix/luya/internal/pkg/util"
)

type M3u8Option struct {
	Url string
}

func FindM3u8(m3u8Option *M3u8Option) {
	fmt.Println("Find M3u8")
	fmt.Println(m3u8Option.Url)

	c := util.DefaultCollector()
	c.Visit("https://hackerspaces.org")
}
