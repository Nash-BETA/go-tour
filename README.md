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
