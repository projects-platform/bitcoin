## BitCoin


## 什么是区块链


## Hash 函数

&ensp;&ensp;```Hash 函数``` 是一种算法，用于将任意长度的二进制数据映射为固定长度的二进制数据。

&ensp;&ensp;一个好的 ```Hash 函数``` 应该具备以下特点：

- 确定性：对同一个输入数据，无论进行多少次 ```Hash``` 计算，每次都会得到形同结果。

- 单向性：可以对已知的输入数据很容易的计算 ```Hash 值``` ,但是通过已知的 ```Hash 值``` 逆向推算输入数据却很难。

- 隐秘性: 一个好的 ```Hash 函数``` 是抗暴力破解的，不能通过已知的 ```Hash 值```  扭向推算输入数据。

- 抗篡改：对一个数据块，哪怕只是更改一个 ```byte 位``` 都是对其 ```Hash 值``` 进行了很大的更改。

- 抗碰撞：很难出现系统的数据块。


&ensp;&ensp;```Hash 函数``` 的实现

- ```MD 系列```：最为常用的为 ```MD5```，由于 ```MD5``` 被发现越来越多的问题，因此在较高的安全性上，不推荐使用。

- ```SHA 系列```： 推荐 ```SHA256``` 、 ```SHA3```。

&ensp;&ensp;各个语言中 ```Hash 函数``` 的使用

- Golang

```golang
func calcHash(data string) string {
    hashByte := sha256.Sum256([]byte(data))
    hashData := hex.EncodeToString(hashByte[:])
    return hashData
}
```