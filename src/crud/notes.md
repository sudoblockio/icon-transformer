
### TTD 

1. Make GetXXXCrud generic
- With interface 
  - https://stackoverflow.com/a/71216396/12642712

```go
type Crud interface {
    BlockCrud | TransactionCrud | ...
}

func GetCrud[c Crud]() c{
    // boilerplate -> Could be another function 
    a := new(p)
    i := any(a)
	
	swich i.(type) {
    case BlockCrud:
        return &blockCrud{}
    case TransactionCrud:
        return &transactionCrud{}
    }
}
```


2. Start loader channel generic 
- Needs to take generic function, interface with methods, and struct



3. Make loader channel to batch upserts 
- Do reflection step in loop 
  - Make struct that 

4. Make loader channels for targetted upserts restricted to less channels 