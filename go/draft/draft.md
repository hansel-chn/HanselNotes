```
package main  
  
import "fmt"  
  
func main() {  
   a:=test{}  
   a.add1()  
   b:=test{}  
   (&b).add1()  
   fmt.Println(a)  
   fmt.Println(b)  
  
   // considering c  
   //There is a handy exception, though. When the value is addressable, the language takes care of the   //common case of invoking a pointer method on a value by inserting the address operator automatically.   c:=test{}  
   c.add2()  
   d:=test{}  
   (&d).add2()  
   fmt.Println(c)  
   fmt.Println(d)  
  
   // but considering e, there isn't a handy exception. Interface doesn't think that variable e realize testInterface2  
   // even if variable e is addressable   e:=test{}  
  
   testInterface1(e)  
   testInterface2(e)  
   testInterfacemix(e)  
  
   f:=&test{}  
   testInterface1(f)  
   testInterface2(f)  
   testInterfacemix(f)  
}  
  
type test []int  
  
func (t test) add1() {  
   t = append(t, 11111)  
}  
func (t *test) add2() {  
   *t = append(*t, 222222)  
}  
  
type AAA1 interface {  
   add1()  
}  
  
type AAA2 interface {  
   add2()  
}  
  
type AAAmix interface {  
   add1()  
   add2()  
}  
func testInterface1(test AAA1)  {  
  
}  
func testInterface2(test AAA2)  {  
  
}  
  
func testInterfacemix(test AAAmix) {  
}
```

Instead of thinking about "accessing a field" why don't you state what you want or need to do with that type (the action that you would be doing with that field) and implement that as a method instead?

Ask yourself: Do I really need to access the Name or do I have to do XYZ with that Name? If you really really need to access that Name everywhere then use struct embedding or implement a `Name() string` method (that's a getter in Go) but if you answer the latter then that action would be a method you can implement in every type that implements the interface

# For statement confusion
[https://go.dev/blog/loopvar-preview](https://go.dev/blog/loopvar-preview)
[https://go.dev/ref/spec#For_statements](https://go.dev/ref/spec#For_statements)
[https://stackoverflow.com/questions/25919213/why-does-go-handle-closures-differently-in-goroutines](https://stackoverflow.com/questions/25919213/why-does-go-handle-closures-differently-in-goroutines)
[https://go.dev/wiki/CommonMistakes](https://go.dev/wiki/CommonMistakes)
[https://stackoverflow.com/questions/26692844/captured-closure-for-loop-variable-in-go](https://stackoverflow.com/questions/26692844/captured-closure-for-loop-variable-in-go)
[https://www.reddit.com/r/golang/comments/3asu5u/a_little_confused_about_closures/](https://www.reddit.com/r/golang/comments/3asu5u/a_little_confused_about_closures/)
在Go 1.22以前，不管是否涉及到并发，是否使用goroutine，for语句都会面临闭包内函数令人困惑的行为。在Go 1.22后，问题解决但仍要注意细节，见Go spec。
理解Function Literals
理解In a sense, _every_ function is a closure.

Each iteration has its own separate declared variable (or variables) [[Go 1.22](https://go.dev/ref/spec#Go_1.22)]. The variable used by the first iteration is declared by the init statement. The variable used by each subsequent iteration is declared implicitly before executing the post statement and initialized to the value of the previous iteration's variable at that moment.

for and for range修改数组的区别

```
// go version 1.19
package main

import "fmt"

func main() {
	//a := []int{0, 1}
	//for i := 0; i < len(a); i++ {
	//	a = []int{}
	//	fmt.Println(a[i])
	//	a = append(a, i)
	//}

	a := []int{0, 1}
	for i, v := range a {
		a[1] = 1111111
		//a = []int{}
		fmt.Println(v)
		a = append(a, i)
	}

	//a := []int{0}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//	a = append(a, i)
	//}
}
```


# Golang Queue
[https://stackoverflow.com/questions/2818852/is-there-a-queue-implementation](https://stackoverflow.com/questions/2818852/is-there-a-queue-implementation)

# link
1.  [What does an empty select do?](https://stackoverflow.com/questions/18661602/what-does-an-empty-select-do)
2.  [Go concurrency with for loop and anonymous function behaves unexpectedly](https://stackoverflow.com/questions/36776315/go-concurrency-with-for-loop-and-anonymous-function-behaves-unexpectedly)
