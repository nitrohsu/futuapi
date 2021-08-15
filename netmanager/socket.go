package netmanager

import (
	"bufio"
	"log"
	"net"
	"nitrohsu.com/futu/protocol"
)

type NetManager struct {
	Host        string
	Port        string
	ReadBuffer  chan *protocol.Message
	WriteBuffer chan *protocol.Message
	reader      *bufio.Reader
	writer      *bufio.Writer
	quit        chan bool
	connection  net.Conn
}

func (netMgr *NetManager) Connect() error {

	netMgr.quit = make(chan bool, 1)
	netMgr.WriteBuffer = make(chan *protocol.Message, 50)
	netMgr.ReadBuffer = make(chan *protocol.Message, 100)

	conn, err := net.Dial("tcp", netMgr.Host+":"+netMgr.Port)
	if err != nil {
		return err
	}

	netMgr.writer = bufio.NewWriter(conn)
	netMgr.reader = bufio.NewReader(conn)
	netMgr.connection = conn

	go netMgr.ReadTCP()
	go netMgr.WriteTCP()

	return nil
}

func (netMgr *NetManager) Close() {
	netMgr.connection.Close()

	close(netMgr.WriteBuffer)
	close(netMgr.ReadBuffer)
	close(netMgr.quit)
}

func (netMgr *NetManager) WriteTCP() error {
	for {
		select {
		case ftProto := <-netMgr.WriteBuffer:
			_, err := ftProto.Write(netMgr.writer)
			if err != nil {
				return err
			}
		case <-netMgr.quit:
			log.Printf("tcp write exit")
			return nil
		}
	}
	return nil
}

func (netMgr *NetManager) ReadTCP() error {
	for {
		ftProto := &protocol.Message{}
		_, err := ftProto.Read(netMgr.reader)
		if err != nil {
			return err
		} else {
			netMgr.ReadBuffer <- ftProto
		}
	}
	return nil
}
