# September 30th, 2018 - Praticing REST API concepts

Traversy media put out a tutorial on YouTube going over basic REST concepts and building an API using he gorilla/mux package. 

Everything finally fucking clicked today with this project, as I annotated this block of code:

```go
//Create a new book function
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

In the function `createBook()` I am passing in two **optional** parameters. An http request from an outside source (`r *http.Request`, making `r` my variable for the `*http.Request` variable type), and then choosing to respond back to the requester (optional) using the `w http.ResponseWriter` (`w` being the variable for the `http.ResponseWriter` variable type).

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

