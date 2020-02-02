# order

  Order can sort the string slice by inner number.

## How To Use
```go
go get github.com/duttty/order

package main

import github.com/duttty/order

func main(){
    s := []string{"dotisff", "我是第123g123", "lkjasd", "spoi1eq", "第12讲但是", "第讲asd12", "萨拉丁就卡死", "qweq24s", "第20讲", "asdads25"}
    sorted := sort.Strings(s)
    fmt.Println(sorted)
}


//[spoi1eq 第12讲但是 第讲asd12 第20讲 qweq24s asdads25 我是第123g123 dotisff lkjasd 萨拉丁就卡死]

```