package main

import (
	"fmt"
	"log"
)

type Bird struct {
	Age int
	Name string
}

func passV(b Bird) {
	b.Age++
	b.Name = "passed" + b.Name
	log.Printf("传入passV方法中Bird: \t %+v Memory Address %p \n", b, &b)
}

func passPtr (b *Bird) {
	b.Age ++
	b.Name = "Ptr" + b.Name
	log.Printf("在passPtr中Bird的值：%+v \t Memory Address %p \t Ptr Address %p \n", *b, b, &b)
	// *b 是一个值
	// b 是被传进来的地址
	// &b 是被传进来的地址（指针）的地址
}

var cat3 = Bird{Age: 1, Name:"Cat3"}
var cat4 = cat3


func main() {

	/*
		第一部分， 搞清楚指针是什么 （复习一下C++）
	 */
	parrot := Bird{Age: 1, Name: "Blue"}

	log.Printf("传入前: \t %+v Memory Address %p \n", parrot, &parrot)
	passV(parrot)
	log.Printf("传入后: \t %+v Memory Address %p \n", parrot, &parrot)
	//方法并不会改变值

	parrot2 := Bird{Age:2, Name: "Green"}

	ptrP2 := &parrot2
	log.Printf("传入前: \t %+v Memory Address %p \n", parrot2, &parrot2)
	passPtr(ptrP2)
	log.Printf("传入后: \t %+v Memory Address %p \n", parrot2, &parrot2)
	//此时由于传入的是地址， 所以， 最后更改生效了没毛病。

	//指针简单的使用原则
		// 1 不想修改变量就传 T， 如果想要修改就传 *T
		// 2 如果变量非常大 副本会影响性能， 只考虑 *T
		// 3 在函数作用域， T或分配到栈， *T会分配到对象， 垃圾回收会有影响

	/*
		第二部分， 什么时候会有副本创建
			似乎只要是赋值， 都会搞出来副本。
	 */

	var cat1 = Bird{Age: 11, Name:"Cat1"}
	var cat2 = cat1
	log.Printf("Values: \t %+v Memory Address %p \n", cat1, &cat1)
	log.Printf("Values: \t %+v Memory Address %p \n", cat2, &cat2)
	//在这里 Cat1 和 Cat2 值虽然相同， 但是地址不同， 也就是改变一个， 并不会改变另一个
	// 在函数外的话, 也没什么区别
	var cat5 = cat3
	log.Printf("Values: \t %+v Memory Address %p \n", cat3, &cat3)
	log.Printf("Values: \t %+v Memory Address %p \n", cat4, &cat4)
	log.Printf("Values: \t %+v Memory Address %p \n", cat5, &cat5)
	cat6 := cat3
	log.Printf("Values: \t %+v Memory Address %p \n", cat6, &cat6)
	passPtr(&cat6)
	log.Printf("ValuesCat 6: \t %+v Memory Address %p \n", cat6, &cat6)
	log.Printf("ValuesCat 3: \t %+v Memory Address %p \n", cat3, &cat3)

	/*
		第三 map， slice， 和 Array
	 */

	// 常见的slices
	log.Printf("Values: \t %+v Memory Address %p \n", cat1, &cat1)
	s := []Bird{cat1}
	s = append(s,cat1)
	cat1.Age=3
	log.Printf("Values s[0]: \t %+v Memory Address %p \n", s[0], &s[0])
	log.Printf("Values s[1]: \t %+v Memory Address %p \n", s[1], &s[1])
	//可以看出， 加入到slice之后， 就跟cat1没关系了， 地址也完全不同了

	//map

	m := make(map[int]Bird)
	m[0] = cat1
	cat1.Age=4
	log.Printf("Value of m[0] is:\t %+v", m[0])
	cat7 := m[0]
	log.Printf("Value of cat 7 is:\t %+v", cat7 )
	//注意了， 其实每个赋值都是copy进去的， 除非是int， 否则没得修改。 所以跟原来的cat没关系了， 于是如果非的要修改map， 那么可以传指针
	mm := make(map[int]*Bird)
	mm[0]= &cat1
	log.Printf("Value of cat 1 is:\t %+v", cat1 )
	log.Printf("Value of mm[0] is:\t %+v", *mm[0])
	mm[0].Age+=1
	log.Printf("Value of cat 1 is:\t %+v", cat1 )
	log.Printf("Value of mm[0] is:\t %+v", *mm[0])
	mm[1] = &Bird{3,"2B"}
	log.Printf("Value of mm[1] is:\t %+v", *mm[1])

	cat1.Age = 1
	a := [2]Bird{cat1}
	cat1.Age = 6
	fmt.Printf("parrot6:\t\t %+v, \t\t内存地址：%p\n", a[0], &a[0])
	cat1.Age = 1
	a[1] = cat1
	cat1.Age = 7
	fmt.Printf("parrot7:\t\t %+v, \t\t内存地址：%p\n", a[1], &a[1])







}
