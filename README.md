# hgo
Help Go, es una librería para ayudar a extender la sintaxis de golang 

## If short (condición `SI` corta)

```go
    import h "github.com/jad21/hgo"

    func main()  {
        if h.If(true, 1, 0) == 1{
            print("is true")
        }
    }
```