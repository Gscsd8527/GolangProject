package main

import (
	"GoLearning/Golearning/4.daobao_he_init_diaoyong/pack1"
	"GoLearning/Golearning/4.daobao_he_init_diaoyong/pack2"

	//_ 表示匿名，导入但并不使用，只会使用里面的init方法
	//_ "GoLearning/Golearning/4.daobao_he_init_diaoyong/pack1"

	// . 表示把包当做当前包，不用使用库.方法名， 直接方法名即可
	//. "GoLearning/Golearning/4.daobao_he_init_diaoyong/pack2"

	// 取个别名
	//mytest "GoLearning/Golearning/4.daobao_he_init_diaoyong/pack2"
)

func main() {
	// import导包路径问题和init方法调用
	pack1.PackTest()
	pack2.PackTest()
}
