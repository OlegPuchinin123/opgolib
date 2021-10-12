/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"bytes"
	"errors"
	"net"
	"os"
)

type OPClient struct {
	c            *net.TCPConn
	hex_log      *os.File
	recv_bytes   *bytes.Buffer
	message_size int
}

func Client_new() *OPClient {
	var (
		client *OPClient
	)
	client = new(OPClient)
	client.hex_log = nil
	client.recv_bytes = nil
	client.message_size = 256
	return client
}

func (client *OPClient) Client_connect(addr_port string) error {
	var (
		addr *net.TCPAddr
		e    error
	)
	addr, e = net.ResolveTCPAddr("tcp4", addr_port)
	if e != nil {
		return e
	}
	client.c, e = net.DialTCP("tcp4", nil, addr)
	if e != nil {
		return e
	}
	return nil
}

func (client *OPClient) Client_set_log(log_file_name string) error {
	var (
		e error
	)
	client.hex_log, e = os.Create(log_file_name)
	return e
}

func (client *OPClient) Client_log(gpb *GPB) {
	if client.hex_log != nil {
		client.hex_log.Write(gpb.buf)
	}
}

func (client *OPClient) Client_send(gpb *GPB) error {
	if gpb == nil {
		return errors.New("Empty buffer.")
	}
	client.Client_log(gpb)
	client.c.Write(gpb.buf)
	return nil
}

func (client *OPClient) Client_recv(size int) (*GPB, error) {
	var (
		buf   []byte
		gpb   *GPB
		count int
		e     error
	)
	buf = make([]byte, size)
	count, e = client.c.Read(buf)
	if e != nil {
		return nil, e
	}
	gpb = NewGPBBuf(buf[:count])
	client.Client_log(gpb)
	if client.recv_bytes != nil {
		client.recv_bytes.Write(gpb.buf)
	}
	return gpb, nil
}

func (client *OPClient) Client_enable_recv_bytes(status bool) {
	if status {
		client.recv_bytes = bytes.NewBuffer(nil)
	} else {
		if client.recv_bytes != nil {
			client.recv_bytes.Reset()
			client.recv_bytes = nil
		}
	}
}

func (client *OPClient) Client_get_obtained() *bytes.Buffer {
	return client.recv_bytes
}

func (client *OPClient) Client_recv_msg() (*GPB, error) {
	var (
		gpb *GPB
		e   error
	)
	gpb, e = client.Client_recv(client.message_size)
	if e != nil {
		return nil, e
	}
	if len(gpb.buf) != client.message_size {
		return nil, errors.New("Bad message size.")
	}
	return gpb, nil
}

func (client *OPClient) Client_close() {
	if client.hex_log != nil {
		client.hex_log.Close()
	}
	client.c.Close()
}

func (client *OPClient) Client_close_write() {
	client.c.CloseWrite()
}

func (client *OPClient) Client_close_read() {
	client.c.CloseRead()
}
