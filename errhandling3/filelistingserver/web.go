package main

import (
	"errhandling3/filelistingserver/filelisting"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter,
						request *http.Request)error

func errWrapper(handler appHandler)func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request){
		err:=handler(writer, request)
		if err!=nil{
			code:=http.StatusOK
			switch{
				log.Warn("Error handling request: %s", err.Error())
			case os.IsNotExist(err):
					code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main(){
	http.HandleFunc("/E%3A/", errWrapper(filelisting.HandleFileList))
	err:=http.ListenAndServe(":8888", nil)
	if err!=nil{
		panic(err)
	}
}