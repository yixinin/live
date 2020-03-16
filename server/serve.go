package server

import (
	"fmt"
	"go-lib/ip"
	"go-lib/log"
	"go-lib/utils"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"go-lib/registry"
)

const ()

func (s *Server) Serve() error {

	for _, v := range s.conf.Protocol {
		if v == ProtocolTCP {
			go s.ServeSocket()
			s.stopTcp = make([]chan bool, 0, 10)
		}
		if v == ProtocolTCP {
			s.ServeWs()
			s.stopWs = make([]chan bool, 0, 10)
		}
	}
	go s.ServeHttp()
	go s.handleCreateRoom()
	var srv = &registry.Service{
		Name:    "live-chat.voip",
		Version: "1.0",
		Nodes: []*registry.Node{
			&registry.Node{
				Id:      utils.UUID(),
				Address: ip.LocalIP,
			},
		},
	}
	s.Registry.Register(srv, registry.RegisterTTL(time.Minute))
	return nil
}

func (s *Server) ServeSocket() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr
	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", s.conf.ListenIp, s.conf.TcpPort))
	log.Warn("listen socket/tcp in", tcpAddr.String())
	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		netconn, err := tcpListener.Accept()
		if err != nil {
			log.Errorf("accept conn error: %v", err)
			continue
		}
		// conn := core.NewConn(netconn, 4*1024)
		log.Warn("remote:", netconn.RemoteAddr().String(), "local:", netconn.LocalAddr().String())
		go s.handleTcp(netconn)
	}
}

func (s *Server) ServeWs() {
	var addr = fmt.Sprintf("%s:%s", s.conf.ListenIp, s.conf.HttpPort)
	http.HandleFunc("/live/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Warn("upgrade:", err)
			return
		}
		s.handleWs(conn)
	})
	log.Warn("listen ws in", addr+"/live")
}

func (s *Server) ServeHttp() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
		http.HandleFunc("/live/http", func(w http.ResponseWriter, r *http.Request) {
			var buf, err = ioutil.ReadAll(r.Response.Body)
			if err != nil {
				w.Write([]byte(""))
				return
			}
			s.handleHttp(buf)
		})
	}()
	var addr = fmt.Sprintf("%s:%s", s.conf.ListenIp, s.conf.HttpPort)
	log.Warn("listen http in", addr)
	log.Faltal(http.ListenAndServe(addr, nil))
}
