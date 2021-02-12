# Memo

## 関数とメソッドの定義

### 関数の定義
```
func functionName () {
  dosomething..
}
```

### メソッドの定義
```
type Point struct{
  X, Y int
}

func (p *Point) addMethod (point *Point) int {
  x = point.X
  y = point.Y
  return x + y
}
```

関数と異なりfuncとメソッド名の間にレシーバーの型とその変数名が必要になる。  
ここでは` (p *Point) `がレシーバーになり、` (point *Point) `が引数になる。
