# 一个简单的生成SQL where条件的pkg

## Installation

``` bash
go get github.com/wuyunhua1987/query-builder
```

## Usage

```go
import builder "github.com/wuyunhua1987/query-builder"

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
// (`id` = ? AND `name` like ? AND `status` IN (?) AND `delete` IS NULL AND (`email` <> ? OR `state` <> ? OR (`phone` IS NOT NULL AND `tel` = ?) OR `del` = ?))
// [1 %a% 1,2 a@b.com 2,3 01 0]
```

## 