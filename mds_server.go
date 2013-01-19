package main

import (
	"fmt"
	"net"
)
/*
  定义两个常量 server_ip和 server_port
*/
const (
	server_ip="127.0.0.1"
	server_port=19781
)
/*
  入口主函数
*/
func main() {
	server()
}
/*
 MDS-SERVER原理是自架一台服务器，程序模拟MDS服务的简单回应。客户端使用黑莓断续膏软件
 保持模拟在线。
*/
func server() {
	exit:=make(chan bool)
	ip:=net.ParseIP(server_ip)
	addr:=net.UDPAddr{ip,server_port}

	/*
	 使用go的goroutine实现多线程并发支持
	*/
	 go func () {
	 	//监听UDP端口
	 	listen,err:=net.ListenUDP("udp",&addr)
        if err!=nil {
        	fmt.Println("初始化失败",err.Error())
        	exit<-true
        	return
        }
        fmt.Println("正在监听...")
        defer listen.Close()
        for {
        	data:=make([]byte,4096)
        	read,remoteAddr,err:=listen.ReadFromUDP(data)
        	if err!=nil {
        		fmt.Println("读取数据失败！",err)
        		continue
        	}
        	fmt.Println(read,remoteAddr)
        	sendData:=[]byte{0x10,0x8,0,0,0,0,data[3],data[4],data[5],data[6],data[11]-128}
        	_,err=listen.WriteToUDP(sendData,remoteAddr)
        	if err!=nil {
        		return
        		fmt.Println("发送数据失败!",err)
        	}
        }
	 }();
	 <-exit
	 fmt.Println("服务器关闭！")
}
