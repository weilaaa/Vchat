package utils

import (
	"Vchat/common/message"
	"crypto/sha256"
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

	hash1 := make([]byte, 32)
	_, err = this.Conn.Read(hash1)
	if err != nil {
		if err == io.EOF {
			fmt.Printf("connection has closed by %s\n", this.Conn.RemoteAddr().String())
			return
		}
		err = errors.New("hash read failed")
	}

	cache := make([]byte, 32)
	total := 0
	for {

		buf := make([]byte, 32)
		n, err := this.Conn.Read(buf)
		if err != nil {
			fmt.Println("data read failed")
		}
		total += n
		cache = append(cache, buf[:n]...)
		if n < 32 {
			break
		}
	}

	fmt.Println("received:", total, "hash:", hash1)

	hash2 := sha256.Sum256(cache[32:total+32])
	if string(hash1) != string(hash2[:32]) {
		fmt.Println("data transfer abnormal")
		return
	}

	// unmarshal data to a message
	// !!! Attention, the second para should be a reference
	err = json.Unmarshal(cache[32:total+32], &mes)
	if err != nil {
		fmt.Println("message unmarshal failed", err)
		return
	}

	return

}

func (this *Transfer) WritePKG(data []byte) (err error) {
	hash := sha256.Sum256(data)

	// ship the container
	n, err := this.Conn.Write(hash[:32])
	if n != 32 || err != nil {
		log.Fatalln("hash sent failed", err)
	}

	// send 32 bit every time
	ptr := 0
	total := 0
	for {
		if len(data) <= 32 {
			n, err = this.Conn.Write(data)
			if err != nil {
				log.Fatalln("data sent failed", err)
			}
			total = n
			break
		}
		if ptr+32 >= len(data) {
			n, err = this.Conn.Write(data[ptr:])
			if err != nil {
				log.Fatalln("data sent failed", err)
			}
			total += n
			break
		}
		n, err = this.Conn.Write(data[ptr : ptr+32])
		if err != nil {
			log.Fatalln("data sent failed", err)
		}
		total += n
		ptr += 32
	}
	fmt.Println("sent:", total, "hash:", hash)
	return
}
