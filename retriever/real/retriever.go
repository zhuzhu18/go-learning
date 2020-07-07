package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct{
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) Get(url string) string{
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}else{
		result, err := httputil.DumpResponse(resp, true)
		resp.Body.Close()
		if err != nil{
			panic(err)
		}else{
			return string(result)
		}
	}
}