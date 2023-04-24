package main

import (
    "encoding/base64"
    "fmt"
)

func textToBinary(text string) ([]byte, error) {
    binaryData := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
  /*
    make 关键字用于创建一个指定类型的对象，并返回其指针
    只能用于创建 slice、map 和 channel
    make(T, args)
    其中 T 表示要创建的对象类型，args 表示指定的参数
    base64.StdEncoding 是一个用于 base64 编码的标准编码器
    EncodedLen 函数是 StdEncoding 类型的一个方法，用于计算将 n 个字节编码为 base64 格式所需的输出字节数
    函数签名--->func (enc *Encoding) EncodedLen(n int) int
    enc 表示 base64 编码器的实例，n 表示要编码的字节数。EncodedLen 函数返回值为将 n 个字节编码为 base64 格式所需的输出字节数
  */
    base64.StdEncoding.Encode(binaryData, []byte(text))
    return binaryData, nil
}

func main() {
    text := "hello world"
    binaryData, err := textToBinary(text)
    if err != nil {
        fmt.Println("Failed to convert text to binary:", err)
        return
    }
    fmt.Println(binaryData)
}
