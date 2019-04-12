package utils

import (
	"Vchat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [1024 * 4]byte
}

func (this *Transfer) ReadPKG() (mes message.Message, err error) {

	// use hash to check date completion later
	// judge if message is losing data
	buf := make([]byte, 1024*4)
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		if err == io.EOF {
			fmt.Printf("connection has closed by %s\n", this.Conn.RemoteAddr().String())
			return
		}
		err = errors.New("length read failed")
	}

	// transfer []uint32 to uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// receive mess content from client
	n, err := this.Conn.Read(buf[:pkgLen])
	if int(pkgLen) != n || err != nil {
		fmt.Println("data read failed")
	}

	// unmarshal data to a message
	// !!! Attention, the second para should be a reference
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("message unmarshal failed", err)
		return
	}

	return

}

func (this *Transfer) WritePKG(data []byte) (err error) {
	// send length of pkg
	var buf [4]byte
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		log.Fatalln("length sent failed", err)
	}

	// ship the container
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		log.Fatalln("data sent failed", err)
	}

	return
}
