# wire2kratos

A small script to resolve this problem. wire issues - Go workspaces broken:
```
go generate ./...
go: -mod may only be set to readonly or vendor when in workspace mode, but it is set to "mod"
        Remove the -mod flag to use the default readonly value, 
        or set GOWORK=off to disable workspace mode.
xxx/xxx/wire_gen.go:3: running "go": exit status 1
```

```
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject
```

## 🔁 Demo Projects

* [demo1kratos](https://github.com/orzkratos/wire2kratos-demos/tree/main/demo1kratos)
* [demo2kratos](https://github.com/orzkratos/wire2kratos-demos/tree/main/demo2kratos)
