/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"bytes"
	"net"
	"os"
)

type OPServer struct {
	c           *net.TCPConn
	ln          *net.TCPListener
	last_client int
}

type ClientStatus struct {
	c          *net.TCPConn
	recv_bytes *bytes.Buffer
	status     int
	hex_log    *os.File
	client_id  int
}

func NewServer() *OPServer {
	var (
		serv *OPServer
	)
	serv = new(OPServer)
	serv.last_client = 0
	return serv
}

func (serv *OPServer) Listen(host_port string) error {
	var (
		e     error
		laddr *net.TCPAddr
	)

	laddr, e = net.ResolveTCPAddr("tcp4", host_port)
	if e != nil {
		return e
	}
	serv.ln, e = net.ListenTCP("tcp4", laddr)
	if e != nil {
		return e
	}
	return nil
}

func (serv *OPServer) Accept() (*ClientStatus, error) {
	var (
		cs *ClientStatus
		e  error
	)
	cs = new(ClientStatus)
	cs.recv_bytes = bytes.NewBuffer(nil)
	cs.hex_log = nil
	cs.c, e = serv.ln.AcceptTCP()
	if e != nil {
		return nil, e
	}
	serv.last_client++
	cs.client_id = serv.last_client
	return cs, nil
}

func (cs *ClientStatus) CS_set_log(file_name string) error {
	var (
		e error
	)
	if cs.hex_log != nil {
		cs.hex_log.Close()
	}
	cs.hex_log, e = os.Create(file_name)
	if e != nil {
		return e
	}
	return nil
}

func (cs *ClientStatus) CS_Recv(size int) (*GPB, error) {
	var (
		buf   []byte
		count int
		e     error
		gpb   *GPB
	)
	buf = make([]byte, size)
	count, e = cs.c.Read(buf)
	if e != nil {
		return nil, e
	}
	cs.hex_log.Write(buf[:count])
	gpb = NewGPBBuf(buf[:count])
	return gpb, nil
}

func (cs *ClientStatus) CS_Send(gpb *GPB) (int, error) {
	var (
		count int
		e     error
	)
	count, e = cs.c.Write(gpb.buf)
	return count, e
}

func (cs *ClientStatus) CS_Close() error {
	if cs.hex_log != nil {
		cs.hex_log.Close()
	}
	return cs.c.Close()
}
