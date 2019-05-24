package utils

/*import (
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

	// use hash to check date completion later
	// judge if message is losing data

	hash1 := make([]byte, 32)
	_, err = this.Conn.Read(hash1)
	if err != nil {
		if err == io.EOF {
			fmt.Printf("connection has closed by %s\n", this.Conn.RemoteAddr().String())
			return
		}
		err = errors.New("hash read failed")
	}

	// transfer []uint32 to uint32
	//var pkgLen uint32
	//pkgLen = binary.BigEndian.Uint32(buf[:4])

	// receive mess content from client
	// try to read it by for loop
	buf := make([]byte, 1024*8)

	n, err := this.Conn.Read(buf)
	if err != nil {
		fmt.Println("data read failed")
	}

	fmt.Println("received:", n, "hash:", hash1)

	hash2 := sha256.Sum256(buf[:n])
	if string(hash1) != string(hash2[:32]) {
		fmt.Println("data transfer abnormal")
		return
	}

	// unmarshal data to a message
	// !!! Attention, the second para should be a reference
	err = json.Unmarshal(buf[:n], &mes)
	if err != nil {
		fmt.Println("message unmarshal failed", err)
		return
	}

	return

}

func (this *Transfer) WritePKG(data []byte) (err error) {
	// send length of pkg

	hash := sha256.Sum256(data)

	// ship the container
	n, err := this.Conn.Write(hash[:32])
	if n != 32 || err != nil {
		log.Fatalln("hash sent failed", err)
	}

	n, err = this.Conn.Write(data)
	if err != nil {
		log.Fatalln("data sent failed", err)
	}

	fmt.Println("sent:", n, "hash:", hash)

	return
}
*/