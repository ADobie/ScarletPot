package mysql

import (
    "fmt"
    "net"
    "scarletpot/utils/pool"
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


func Start() {
    wg, poolX := pool.New(1)
    defer poolX.Release()
    serverAddr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:3306")
    listener, _ := net.ListenTCP("tcp", serverAddr)

    for {
        wg.Add(1)
        poolX.Submit(func () {
            conn, err := listener.Accept()
            if err != nil {
                fmt.Println("zh-CN","Mysql", "127.0.0.1", "Mysql 连接失败",err)
            }
            go connectionHandler(conn)
            wg.Done()
        })
    }
}

func connectionHandler(conn net.Conn) {
    // 结束后关闭连接
    defer conn.Close()
    connFrom := conn.RemoteAddr().String()
    fmt.Println("收到来自",connFrom,"的连接")
    fmt.Println("发送握手包..")
    _, err := conn.Write(handshakePack)
    if err != nil {
       fmt.Println("握手包发送失败..",err)
    }

    _, err = conn.Write(okPack)
    if err != nil {
       fmt.Println("ok包发送失败",err)
    }

}