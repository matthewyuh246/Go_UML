package main

import "fmt"

type Server interface {
	handle(userId string)
}

type RealServer struct{}

func NewRealServer() Server {
	return &RealServer{}
}

func (rs *RealServer) handle(userId string) {
	fmt.Printf("%sの処理を実行中\n", userId)
}

type Proxy struct {
	server Server
}

func NewProxy(server Server) Server {
	return &Proxy{
		server: server,
	}
}

func (p *Proxy) authorize(userId string) {
	authorizedUserId := []string{"1", "2", "3"}
	authorized := false
	for i := 0; i < len(authorizedUserId); i++ {
		if authorizedUserId[i] == userId {
			authorized = true
		}
	}
	if !authorized {
		panic("操作が許可されていません")
	}
}

func (p *Proxy) handle(userId string) {
	p.authorize(userId)

	fmt.Println("処理を開始します")
	p.server.handle(userId)
	fmt.Println("処理が終了しました")
}

func main() {
	server := NewRealServer()
	proxy := NewProxy(server)

	proxy.handle("2")
}
