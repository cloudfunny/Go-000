# 学习笔记

- [Go进阶训练营笔记03-error: 错误处理最佳实践](https://lailin.xyz/post/go-training-03.html)

## 作业说明

### 题目

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
  
### 解答

目录结构如下所示，在示例当中 repository 层，不对外返回底层 sql(orm) 错误，因为考虑到后续这一层代码可能会被拆分成为微服务，那时候就没有 sql 了，所以直接返回在 errcode 包定义的业务错误

```bash
.
├── cmd
│   └── http
│       └── main.go
├── domain
│   └── user.go
├── pkg
│   └── errcode
│       ├── code.go
│       └── error.go
├── test.db
└── user
    ├── api # 外部接口
    │   └── api.go
    ├── repository # 类似 dao 层
    │   └── sqlite.go
    └── usecase # 业务逻辑
        └── usecase.go
```

关键代码如下所示

```go
// Login 用户登录
func (r *repository) Login(username, password string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where(domain.User{Name: username, Password: password}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(errcode.UserLogin, "ErrRecordNotFound")
	}

	if err != nil {
		return nil, errors.Wrapf(errcode.DBQuery, "db errors is: %+v", err)
	}

	return &user, nil
}
```

### 运行

```bash
go run Week02/work/cmd/http/main.go
```