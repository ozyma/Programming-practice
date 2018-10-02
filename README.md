# golang-practice

The purpose of this repository is to refamiliarize myself with Golang, practice new concepts, and create foundations for my Golang education.

**Basic REST concepts:
REST uses only four basic commands for its protocol over HTTP:**

- `GET`: Clients can request the status of a resource by making an HTTP GET request to the server, using the resource's URI (Uinque Resource Identifier). REST requires that thsi operation does not produce any side effect to the resource's status (nullipotent)
- `PUT`: Creates a new resource. Since the client does not know the next invoice number, the URI can be: http://www.mysite.com/invoice/841 (for example) is (and must be) idempotent. Invoice 841 must not be created multiple times if clients call that PUT requests several times.
- `POST`: POST requires POST client requests to update the corresponding resource with information provided by the client, or to create this resource if it does not exist. This operation is not idempotent.
- `DELETE`: This operation removes the resource forever. It is idempotent.

The format (JSON, XML, ...) used to return representations of resources is set in the media type of the server..

To handle success or errors issues, HTTP REST recommends using one of the HTTP status Codes. https://en.m.wikipedia.org/wiki/List_of_HTTP_status_codes

**Everything finally fucking clicked today with this project, as I annotated this block of code:**

```go
//Create a new book function uses pointer receiver so reduce expense of reading. we don't need a copy of the request bogging down our throughput
func createBook(w http.ResponseWriter, r *http.Request) {
	//formats the HTTP headers we will send back when we write our response
	w.Header().Set("Content-Type", "application/json")
	//new variable, of the custom type Book that we made (the struct above)
	var book Book
	//takes the request sent to us (in json) and Decodes it into our &book variable we declared
	_ = json.NewDecoder(r.Body).Decode(&book)
	//randomly generates an id for our book up to the number 100000000, then applies to the "ID" of our book variable (mock id - not safe)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	//takes all the data we've randomly generated and Decoded from the JSON request sent, and appends it into our []books array of Book type objects
	books = append(books, book)
	//Writes a response (variable "w") to the requester with a JSON encoding of the book variable we created through this process
	json.NewEncoder(w).Encode(book)
}
```

In the function `createBook()` I am passing in two **optional** parameters. An http request from an outside source (`r *http.Request`, making `r` my variable for the `*http.Request` variable type), and then choosing to respond back to the requester (optional) using the `w http.ResponseWriter` (`w` being the variable for the `http.ResponseWriter` variable type). We use a pointer receiver for `http.Request` because it's less expensive than having to make a copy of the request we are receiving. We could potentially receive a hefty HTTP request and

## HTTP protocol
HTTP sends shit via codes... **I need to memorize these codes**, but here are the most commonly referred to/used:
- `200` OK
- `400` Bad Request
- `401` Unauthorized
- `403` Forbidden
- `404` Not Found
- `500` Internal Server Error
- `503` Service Unavailable
- `550` Permission Denied

They most commonly send using protocol version 1.1, or:

`HTTP/1.1 200 ok` (HTTP protocol version 1.1, the request went ok)

This information also exists in the `headers` of the HTTP requests and responses, something very important to know. The headers will prepare the response/request of my Go program for the type of information that it might receive. Hence:

```go
    w.Header().Set("Content-Type", "application/json")
```

in the above code. It will be up to the requester to send their correct `HTTP Headers` if they want to interact with my program, and it's also up to me to specify what headers I will need. REST API's for `json` will essentially always be `application/json`. In my case, I am setting the `Content-Type` Header to `application/json` so that when I send back my response, that server will know I am sending content in my response meant to be read as `json`

![](2018-09-30-REST-API-testing/README-img/Golang-HTTP-data-flow.png)

## bbolt

On the 1st of October, I did not get as far with the understanding of the database as I would have liked. I will return to it.

## Correctly initializing struct variables that have an `[]array` type

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

## From the Golang website on pointer vs. value receivers
Should I define methods on values or pointers? ¶
```go
func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value
```
For programmers unaccustomed to pointers, the distinction between these two examples can be confusing, but the situation is actually very simple. When defining a method on a type, the receiver (s in the above examples) behaves exactly as if it were an argument to the method. Whether to define the receiver as a value or as a pointer is the same question, then, as whether a function argument should be a value or a pointer. There are several considerations.

First, and most important, does the method need to modify the receiver? If it does, the receiver must be a pointer. (Slices and maps act as references, so their story is a little more subtle, but for instance to change the length of a slice in a method the receiver must still be a pointer.) In the examples above, if `pointerMethod` **modifies the fields of `s`, the caller will see those changes**, but `valueMethod` **is called with a copy of the caller's argument (that's the definition of passing a value)**, so changes it makes will be invisible to the caller.

By the way, in Java method receivers are always pointers, although their pointer nature is somewhat disguised (and there is a proposal to add value receivers to the language). It is the value receivers in Go that are unusual.

Second is the consideration of efficiency. **If the receiver is large, a big struct for instance, it will be much cheaper to use a pointer receiver.**

Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used. See the section on method sets for details.

For types such as basic types, slices, and small structs, a value receiver is very cheap so unless the semantics of the method requires a pointer, a value receiver is efficient and clear.