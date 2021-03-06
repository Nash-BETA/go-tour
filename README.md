# go-tour

## 短い変数宣言
関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
なお、関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。

## 型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 の別名

rune // int32 の別名
// Unicode のコードポイントを表す

float32 float64

complex64 complex128

## 初期値を与えない場合
ゼロ値は型によって以下のように与えられます:

数値型(int,floatなど): 0
bool型: false
string型: "" (空文字列( empty string ))


## 定数
定数( constant )は、 const キーワードを使って変数と同じように宣言します。

定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。


## Defer
defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。

defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません。

## ポインタ
メモリの保存場所

```
i, j := 42
p := &i         //「&」でポインタ（メモリのアドレス）を変数に代入
fmt.Println(p)  //(ex)0xc000016098
fmt.Println(*p) //42 →メモリアドレス「0xc000016098」に入っている内容
*p = 21         //メモリアドレス「0xc000016098」に入っている内容の書き換え
fmt.Println(i)  //上記で書き換えられたので「21」
```

詳しい説明
<a href="https://qiita.com/Sekky0905/items/447efa04a95e3fec217f">ポインタの説明</a>

## 配列の宣言
intの10個の配列を宣言

```
var a [10]int
```

make 関数はゼロ化された配列を割り当て、その配列を指すスライスを返します。

```
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

## スライスの注意
スライスは配列への参照のようなものです。

```
names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
}
fmt.Println(names) //[John Paul George Ringo]

a := names[0:2]
b := names[1:3]
fmt.Println(a, b) //[John Paul] [Paul George]

b[0] = "XXX"
fmt.Println(a, b)  //[John XXX] [XXX George]
fmt.Println(names) //[John XXX George Ringo]
```

## rage
phpでいうところのforeach
```
for key, value := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
}

//keyのみ
//_を使う
for key, _ := range pow {
}
//２行目の省略
for key := range pow {
}

//valueのみ
for _, value := range pow {
}

```
## map

mapリテラルは、structリテラルに似ていますが、 キー ( key )が必要です。

```
//keyがstringでvalueがint
m := map[string]int
```

## クロージャ
クロージャとは、ネストされた関数定義において、内側の関数が外にある変数を使って処理をすることができる

```
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```

## 値レシーバ

```
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

```
