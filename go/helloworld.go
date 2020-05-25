package main

import "fmt"

//関数外で定義した変数はパッケージ変数といい、同一パッケージならどこからでも参照できる。
var zz = 123

func hello() {
	fmt.Printf("helloworld\n")
}

func calculation() {
	//var [変数名] [型]
	var a int
	a = 5

	//var [変数名]:=[変数値]　|| var [変数名]=[変数値]で型推論され、型の指定の必要がない。

	var b = 3
	var c string
	c = "hi takashi"
	var d = "eiko"

	fmt.Printf("n=%d\n", a)  //5
	fmt.Printf("n=%d\n", b)  //3
	fmt.Printf("n=%s\n", c)  //takashi
	fmt.Printf("n=%s\n", d)  //eiko
	fmt.Printf("n=%d\n", zz) //123

	zz = zz + b

	fmt.Printf("n=%d\n", zz) //126
	/*変数の型メモ

		  論理値型(bool)
		  数値型(int)
		      符号付整数型
		          int8
		          int16
		          int32
		          int64
		      符号なし整数型
		          uint8(byte)
		          uint16
		          uint32
		          uint64
		  浮動小数点型
		      float32
		      float64
		  複素数型
		      complex64
		      complex128
		  rune型(rune)
		  文字列型(string)

		  上記の表示
	論理値(bool) 	%t
	符号付き整数(int, int8など) 	%d
	符号なし整数(uint, uint8など) 	%d
	浮動小数点数(float64など) 	%g
	複素数(complex128など) 	%g
	文字列(string) 	%s
	チャネル(chan) 	%p
	ポインタ(pointer) 	%p


	万能なすべての型に使える　%v

	*/

}

func array() {

	/*配列の宣言
	①var 変数名　[長さ]型
	②var 変数名 [長さ]型 = [大きさ]型{初期値1,初期値n
	③変数名 := [...]型{初期値1,初期値n}}
	*/

	/*for i < 5 {
		d[i] = i * 5
		fmt.Printf("n=%s\n", d[i])
		i++
	}*/
}

func test() {
	var a = 0

	/*if(a == 0){}でもいいけど()つけなくてもいい
	if a<0{
		~~~~~
	}
	else {
		~~~~~~
	}
	このelseの位置はエラーになる。
	*/
	if a == 0 {
		hello()
	} else {
		fmt.Printf("flse\n")
	}

}

func main() {
	//画面表示関数
	//hello()

	//演算関数
	//calculation()

	//配列
	//array()

	//分岐、繰り返し
	//test()
}
