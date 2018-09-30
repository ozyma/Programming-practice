# September 29th, 2018 - Code Refreshing

Today I will be going over the effective Golang on Golang's website, and training myself to write idiomatic Go that is also to standard.

By keeping myself to the standard spec, I can make the code that I write readable by other Golang Developers who practice the spec.

Mistakes I made, and how I fixed them:

the variable type **always** comes after the variable name. I made this mistake when I was creating a struct today. I tried the following:
```go
type Struct car {
	name  string
	speed int32
	cost  int32
}
```
This threw errors, because car is not a data type, and struct is a predetermined Go Syntax that can't be used as a new type. Also `Struct` was capitalized, and I don't need to export it.

Therefore, by putting the `struct` after car, I make a new `type` called `car` that is of the variable type `struct`:

```go
type car struct {
	name  string
	speed int32
	cost  int32
}
```
This ran correctly, and then let me correctly work with my new car type:
```go
	corvette := car{"corvette", 32, 26000}
	fmt.Println(corvette)  //prints: {corvette 32 26000}
        fmt.Println(corvette.name)  //prints: corvette
	fmt.Println(corvette.cost) //prints: 26000
	fmt.Println(corvette.speed) //prints: 32
```