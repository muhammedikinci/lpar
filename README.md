# lpar

```
go get -u github.com/muhammedikinci/lpar
```

Usage
```go
lpar.Param("IDs", []int{1, 2}).
    With("error", errors.New("my sweet error")).
```

Log
```go
log.Error(err, 
    lpar.Param("productID", productID).
    With("categoryID", categoryID),
)
```

String & Json String
```go
myParameters := lpar.Param("IDs", []int{1, 2}).
	With("is_active", true)

fmt.Println(myParameters)
fmt.Println(myParameters.Json())
```