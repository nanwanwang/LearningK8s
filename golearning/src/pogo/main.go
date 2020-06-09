package main

import (
	"flag"
	"pogo/common/logs"
)
var log=logs.Log
func init()  {
	flag.Parse()
}
func main(){
  if flag.NArg()==0{
  	log.Error("no video url parameter!")
	  return
  }
	url:=flag.Arg(0)
	log.Debug(url)
}
