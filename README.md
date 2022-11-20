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