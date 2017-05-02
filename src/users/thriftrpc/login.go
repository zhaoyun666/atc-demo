package thriftrpc

import (
	"fmt"
	"github.com/adolphlxm/atc-demo/src/users/thriftrpc/gen/micro"
	"time"
)

type MicroHandler struct {
	//atc.Handler
}

func (this *MicroHandler) CallBack(callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
	return
}

func (this *MicroHandler) Put(s *micro.Article) (err error) {
	fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%s\tAuthor:%s\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}
