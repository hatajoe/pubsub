# pubsub

pubsub is realize publish and subscribe data model.  
This is sutra copying from [mattn/go-pubsub](https://github.com/mattn/go-pubsub). (It is not exactly the same)  

# Example

```go
// subscriber function
f := func(i int) {
    fmt.Println(i)
}

// initialize pubsub
ps := pubsub.New()

// to subscribe 
s1 := pubsub.NewSub(f)
s2 := pubsub.NewSub(f)
ps.Sub(s1)
ps.Sub(s2)

// to publish
ps.Pub(1)

// to unsubscribe
ps.UnSub(s1)
```

# Licence

This is licenced under the same terms as mattn/go-pubsub.  
