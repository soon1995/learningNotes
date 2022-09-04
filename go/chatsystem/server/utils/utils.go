package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"alsoon.com/go_code/chatsystem/common/message"
)

type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	// 2. read from connection
	_, err = this.Conn.Read(this.Buffer[:4])
	if err != nil {
		// err = errors.New("conn.Read err:")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buffer[:4])

	n, err := this.Conn.Read(this.Buffer[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("conn.Read length incorrect:")
		return
	}

	err = json.Unmarshal(this.Buffer[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json.Unmarshal err:")
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buffer[:4], pkgLen) // marshal pkgLen into binary

	n, err := this.Conn.Write(this.Buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("c.Write failed. err:", err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("c.Write failed. err:", err)
		return
	}

	return
}
