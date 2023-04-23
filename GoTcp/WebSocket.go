package Socket

import (
    "fmt"
    "log"
    "net/url"

    "github.com/gorilla/websocket"
)

func main() {
    // 创建websocket连接
    u := url.URL{Scheme: "ws", Host: "localhost:8080"}
    log.Printf("connecting to %s", u.String())

    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
  /*
    Dial函数是Dialer结构体的一个方法,用于在 Websocket 客户端和服务器之间建立连接
    两个形参:urlStr string,和requestHeader http.Header
    urlStr 是 Websocket 服务器的 URL 地址，requestHeader 是 HTTP 请求头，可以为 nil
    返回 Websocket 连接对象
  */
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer conn.Close()

    // 发送请求并解析返回的response
    request := "hello"
    err = conn.WriteMessage(websocket.TextMessage, []byte(request))
  /*
    返回指向Conn结构体的指针,代表了客户端与 Websocket 服务器之间的连接
    Conn结构体中有一个WriteMessage方法，用于向服务器发送 Websocket 消息
    WriteMessage方法将消息类型和消息内容写入连接的写缓冲区中
    等待发送到 Websocket 服务器。如果写缓冲区已满，WriteMessage方法将会阻塞
    直到有足够的空间来写入消息。
    在发送完消息后，WriteMessage方法将返回一个可能的错误
    如果返回 nil，则表示消息已成功发送到 Websocket 服务器。如果返回一个错误，则表示发送失败
    messageType是一个 int 类型的参数
    用于指定发送的消息类型
    可以是websocket.TextMessage或websocket.BinaryMessage之一
    分别表示文本消息和二进制消息
    data是一个[]byte 类型的参数，表示要发送的消息内容
    对于文本消息，需要将消息内容转换为[]byte 类型
    对于二进制消息，可以直接使用[]byte 类型的数据
  */
    if err != nil {
        log.Println("write:", err)
        return
    }

    _, message, err := conn.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        return
    }

    fmt.Printf("received message: %s\n", message)
}
