# 一个简单的生成SQL where条件的pkg

## 安装

``` bash
go get github.com/wuyunhua1987/query-builder
```

## Usage

```go
import "github.com/wuyunhua1987/query-builder/builder"

b := builder.New()

cond := b.And(
    b.Equal("id", 1),
    b.Like("name", "a", true, true),
    b.In("status", "1,2"),
    b.NULL("delete"),
    b.Or(
        b.NotEqual("email", "a@b.com"),
        b.NotIn("state", "2,3"),
        b.And(
            b.NotNULL("phone"),
            b.Equal("tel", "01"),
        ),
        b.Equal("del", 0),
    ),
)

sql, values := cond.Parse()
fmt.Println(sql, values)
// (`id` = ? AND `name` like ? AND `status` IN (?) AND `delete` IS NULL AND (`email` <> ? OR `state` <> ? OR (`phone` IS NOT NULL AND `tel` = ?) OR `del` = ?)) [1 %a% 1,2 a@b.com 2,3 01 0]
```

## 扩展

如果满足不了你的需求，可以很方便的在自己的项目里扩展这个包

1. 首先定义你自己的where条件，实现`Parse() (string, interface{})`接口

```go
// 假设这里有个位计算的操作
type BitBuilder struct {
	col string
	val interface{}
}

func (b *BitBuilder) Parse() (string, interface{}) {
	return fmt.Sprintf("%s&01111 = ?", b.col), b.val
}
```

2. 在`And Or`中使用

```go
b := builder.New()

cond := b.And(
    b.Equal("id", 1),
    b.Like("name", "a", true, true),
    b.In("status", "1,2"),
    b.NULL("delete"),
    b.Or(
        b.NotEqual("email", "a@b.com"),
        b.NotIn("state", "2,3"),
        b.And(
            b.NotNULL("phone"),
            b.Equal("tel", "01"),
        ),
        b.Equal("del", 0),
    ),
    &BitBuilder{"status", 0x1111}, // <======= 在这里使用用自定义的
)

sql, values := cond.Parse()
fmt.Println(sql, values)
// (`id` = ? AND `name` like ? AND `status` IN (?) AND `delete` IS NULL AND (`email` <> ? OR `state` <> ? OR (`phone` IS NOT NULL AND `tel` = ?) OR `del` = ?) AND status&01111 = ?) [1 %a% 1,2 a@b.com 2,3 01 0 4369]
```

3. 如果要扩展`And Or`也是一样的，只需要实现`Operate()`接口和`Parse() (string, interface{})`接口即可