# lpar

In Golang add your custom fields to map fields easily

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

// output
// is_active: true
// IDs: [1 2]
// {"IDs":"[1 2]","is_active":"true"}
```

Json Marshal for structs
```go
type a struct {
    B string `json:"b"`
}

ab := a{B: "asdasd"}

fmt.Println(lpar.Param("ab", ab).AsJsonString())
// output
// {"ab":"{\"b\":\"asdasd\"}"}
```

With Buffer
```go
bf := bytes.Buffer{}
bf.WriteString("buffer string")

fmt.Println(lpar.Param("test", bf).AsString())
// output
// test: buffer string
```

Log
```go
log.Error(err, 
    lpar.Param("productID", productID).
    With("categoryID", categoryID),
)
```