1. go 接口原理
2. go ...语法糖实现
3. dataframe
4. func copy(dst, src []Type) int
5.指针好还是结构体好 
```text
    type parent struct {
        note string
    }
    type child struct {
        *parent
        id   int
        name string
}
```