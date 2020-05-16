package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//建立服务器,service这里是interface,随便啥类型都可以
func ServeRpc(host string,service interface{})error{
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err!=nil{
		return err
	}

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Printf("accept error:%v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}

//建立客户端
func NewClient(host string) (*rpc.Client, error){
	conn,err:=net.Dial("tcp",host)
	if err!=nil{
		return nil, err
	}
	return jsonrpc.NewClient(conn),nil
}
