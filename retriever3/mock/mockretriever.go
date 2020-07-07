package mock

import "fmt"
type Retriever struct {
	Contents string
}

func (r *Retriever)String()string{          // 实现系统接口String()string函数
	return fmt.Sprintf("Retriever: {contents = %s}", r.Contents)
}
func (r *Retriever)Post(url string, form map[string]string)string{
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever)Get(url string)string{
	return r.Contents
}