package wsy

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"strings"

	u "github.com/cbot918/liby/util"
)

type ConnHandler struct {
	Conn net.Conn
	Data	Data
	Frame	Frame
	Message string
}
type Data struct{
	Host 										[]string
	Connection 							[]string
	Pragma 									[]string
	CacheControl						[]string
	UserAgent								[]string
	Upgrade									[]string
	Origin									[]string
	SecWebSocketVersion			[]string
	AcceptEncoding					[]string
	AcceptLanguage					[]string
	SecWebSocketKey					[]string
	SecWebSocketExtensions	[]string
	SecWebSocketAccept				string
}
type Frame struct {
	Fin						byte
	Opcode				byte
	IsMasked			byte
	PayloadLen		byte
	Mask					[]byte
	Payload				[]byte
}

func NewConnHandler(conn net.Conn) *ConnHandler{
	r := new(ConnHandler)
	r.Conn = conn
	return r
}

func (c *ConnHandler) Upgrade(request string) error {
	webSecSocketKey := c.parseJson(request)
	secWebSocketAccept := c.getReturnSec(webSecSocketKey)
	response := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n",secWebSocketAccept)
	err := c.WriteSocket(response)
	return err
}

func (c *ConnHandler) ReadSocket() []byte{
	req := make([]byte, 4096)
	_, err := c.Conn.Read(req)
	u.Checke(err, "conn read failed")
	return req
}

func (c *ConnHandler) WriteSocket(input string) error{
	res := []byte(input)
	_, err := c.Conn.Write(res)
	return err
}

// decode frame
func (c *ConnHandler) DecodeFrame(data []byte) []byte {
	firstByte := data[0]
	secondByte := data[1]

	c.Frame.Fin = firstByte & 0b10000000
	c.Frame.Opcode = firstByte & 0b00001111
	c.Frame.IsMasked = secondByte & 0b10000000
	c.Frame.PayloadLen = secondByte & 0b01111111
	
	// process mask
	mask := []byte{data[2],data[3],data[4],data[5]}
	fmt.Println("mask: ",mask)

	// process payload data
	payload := []byte{}
	for i:=6; i<=int(c.Frame.PayloadLen+6); i++ {
		payload = append(payload, data[i])
	}
	fmt.Println("payload: ",payload)

	// XOR payload and mask
	result := []byte{}
	for i,item := range payload {
		result = append(result, item ^ mask[i%4])
	}

	return result
}

// parse first upgrade request and get SecKey
func (c *ConnHandler) parseJson(firstRequest string) []string{
	var secWebSocketKey []string
	dataArray := strings.Split(strings.ReplaceAll(firstRequest, "\r\n","\n"),"\n")
	
	for index, item := range dataArray {
		if index == 0 { continue }
		key, value := c.lineToKV(item)
		switch strings.Join(key, "") {
		// case "Host": { w.Data.Host = value }
		// case "Connection" : { w.Data.Connection = value }
		case "Sec-WebSocket-Key" : { c.Data.SecWebSocketKey = value; secWebSocketKey = value }
		}
	}
	return secWebSocketKey
}
func (c *ConnHandler) lineToKV(line string) ([]string,[]string){
	return c.readBefore(":",line),c.ReadAfter(":",line)
}
func (c *ConnHandler) ReadAfter(targetChar string, content string) []string{
	var buf []string
	var flag bool
	for _, char := range content {
		if string(char) == targetChar { flag = true; continue }
		if (flag){ buf = append(buf, string(char))}
	}
	return buf
}
func (c *ConnHandler) readBefore(targetChar string, content string) []string{
	var buf []string
	for _, char := range content {
		if string(char) != targetChar {
			buf = append(buf, string(char))	
		} else { break }
	}
	return buf
}

// get encoded SecKey 
func (c *ConnHandler) getReturnSec(webSecSocketkey []string) string {
	secWebSocketKey := strings.TrimSpace(strings.Join(webSecSocketkey, ""))
	var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
	h := sha1.New()
	h.Write([]byte(secWebSocketKey))
	h.Write(keyGUID)
	secWebSocketAccept := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return secWebSocketAccept
}

// not fix yet
// func (c *ConnHandler) GetFrameDetail() string{
// 	return fmt.Sprintf("fin: %d\nopcode: %d\nis_masked: %d\npayload_len: %d\n",
// 	w.CurrentFrame.Fin,w.CurrentFrame.Opcode,w.CurrentFrame.IsMasked,w.CurrentFrame.PayloadLen)
// }
