# lpar

```
go get -u github.com/muhammedikinci/lpar
```

Usage
```go
lpar.Param("IDs", []int{1, 2}).
    With("error", errors.New("my sweet error")).
```

Map to String & Json String
```go
myParameters := lpar.Param("IDs", []int{1, 2}).
	With("is_active", true)

fmt.Println(myParameters.AsString())
fmt.Println(myParameters.AsJsonString())
```

Json Marshal for structs
```go
type a struct {
    B string `json:"b"`
}

ab := a{B: "asdasd"}

fmt.Println(lpar.Param("ab", ab).AsJsonString())
```

Log
```go
log.Error(err, 
    lpar.Param("productID", productID).
    With("categoryID", categoryID),
)
```