## BitCoin


## 什么是区块链


## Hash 函数

&ensp;&ensp;Hash函数 是一种算法，用于将任意长度的二进制数据映射为固定长度的二进制数据。

&ensp;&ensp;一个好的 Hash函数 应该具备以下特点：

- 确定性：对同一个输入数据，无论进行多少次 Hash 计算，每次都会得到形同结果。

- 单向性：可以对已知的输入数据很容易的计算 Hash值 ,但是通过已知的 Hash值 逆向推算输入数据却很难。

- 隐秘性: 一个好的 Hash函数 是抗暴力破解的，不能通过已知的 Hash值 扭向推算输入数据。

- 抗篡改：对一个数据块，哪怕只是更改一个 byte位 都是对其 Hash值 进行了很大的更改。

- 抗碰撞：很难出现系统的数据块。


&ensp;&ensp;Hash函数 的实现

- MD 系列：最为常用的为 MD5，由于 MD5 被发现越来越多的问题，因此在较高的安全性上，不推荐使用。

- SHA 系列： 推荐 SHA256 、SHA3。

&ensp;&ensp;各个语言中 Hash函数 的使用

- Golang

```golang
func calcHash(data string) string {
    hashByte := sha256.Sum256([]byte(data))
    hashData := hex.EncodeToString(hashByte[:])
    return hashData
}
```
