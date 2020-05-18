package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//把服务单独拿了出来了，之后itemsaver和worker都可以复用
//建立服务器,service这里是interface,随便啥类型都可以
func ServeRpc(host string,service interface{})error{
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err!=nil{
		return err
	}
	//端口成功监听后给个反馈
	log.Printf("Listening on %s",host)

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
