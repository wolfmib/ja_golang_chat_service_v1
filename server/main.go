package main


import {
	"os"
	glog "google.golang.org/grpc/grpclog"
	proto "github.com/wolfmib/ja_golang_chat_service_v1/proto/chat"
}

var ja_log glog.LoggerV2

// 🎬: Initial
func init(){
	//🖨 Setting Loger: all (info, warining, error) message  is going to stdout
	ja_log = glog.NewLoggerV2(os.Stdout,os.Stdout,os.Stdout)
}


// 🕋 Create Connection Structure:
type Connection struct{
	
}

func main(){


}