#### 下载 GORM
```
go get -u gorm.io/gorm
```

#### 模型定义
模型是使用普通结构体定义的,结构体可以包含具有基本Go类型、指针或这些类型的别名，甚至是自定义类型。
```go
package main

import "time"
import "database/sql"

type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // 一个常规字符串字段
  Email        *string        // 一个指向字符串的指针, allowing for null values
  Age          uint8          // 一个未签名的8位整数
  Birthday     *time.Time     // A pointer to time.Time, can be null
  MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
  ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
  CreatedAt    time.Time      // 创建时间（由GORM自动管理）
  UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}
```
在此模型中：
- 具体数字类型如`uint`、`string`和`uint8`直接使用。
- `*string`和`*time.Time`类型的指针表示可空字段。
- 来自`database/sql`包的`sql.NullString`和`sql.NullTime`用于具有更多控制的可空字段。
- `CreatedAt`和`UpdatedAt`是特殊字段，当记录被创建或更新时，`GORM`会自动向内填充当前时间。

除了 GORM 中模型声明的基本特性外，强调下通过 serializer 标签支持序列化也很重要。 此功能增强了数据存储和检索的灵活性，特别是对于需要自定义序列化逻辑的字段。
#### 约定
1. 主键：`GORM`一个名为`id`的字段作为每一个模型的默认主键。
2. 表名：默认情况下，`GORM`将结构体名称转换为`snake_case`并为表名加上复数形式。
3. 列名：`GORM`自动将结构体字段名称转换为`snake_case`作为数据库中的列名。
4. 时间戳字段：`GORM`使用字段`CreatedAt`和`UpdatedAt`来自动跟踪记录的创建和更新时间。
#### 字段级权限控制
```
type User struct {
  Name string `gorm:"<-:create"` // 允许读和创建
  Name string `gorm:"<-:update"` // 允许读和更新
  Name string `gorm:"<-"`        // 允许读和写（创建和更新）
  Name string `gorm:"<-:false"`  // 允许读，禁止写
  Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
  Name string `gorm:"->;<-:create"` // 允许读和写
  Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
  Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
  Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
  Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
}
```


#### 创建/更新时间追踪（纳秒、毫秒、秒、Time）
```
type User struct {
  CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
  UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
  Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间
  Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
  Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}
```