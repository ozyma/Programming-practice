# October 2nd, 2018 - Exploring pointers

Pointers are used as reference point of memory locations, and what is stored as memory points. Today we will be used for basic practice as well as understanding and commenting on lines where pointers can be explored.

## daily practice app
Today we will work on a simple practice app that uses structs and allow us to print change values of those structs.

Progress and issues will be reported below:

I put a field `fillings` with the type of `[]string`, and I had trouble initializing a new variable with that array type. The proper way to initialize my variable, `crabbymelt`, was the following:

```go
type grilledcheese struct {
	bread    string
	cheese   string
	fillings []string
	price    float32
}

crabbymelt := grilledcheese{"sourdough", "mozzerella", []string{"crab meat", "old bay", "scallions"}, 8.99}
```

We can see that I needed to input a `[]string` type of variables directly into my initializtion parameters for the `crabbymelt` variable. I can't use: `{"crab meat", "old bay", "scallions"}` It won't won't work, because I have not let the new `grilledcheese` object know that I'm sending these variables into memory as a `[]string` arrangement. It will only try to send in the variables in sequence, not a `[]string` sequence.