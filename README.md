# golang-practice

The purpose of this repository is to refamiliarize myself with Golang, practice new concepts, and create foundations for my Golang education.

Basic REST concepts:
REST uses only four basic commands for its protocol over HTTP:

- `GET`: Clients can request the status of a resource by making an HTTP GET request to the server, using the resource's URI (Uinque Resource Identifier). REST requires that thsi operation does not produce any side effect to the resource's status (nullipotent)
- `PUT`: Creates a new resource. Since the client does not know the next invoice number, the URI can be: http://www.mysite.com/invoice/841 (for example) is (and must be) idempotent. Invoice 841 must not be created multiple times if clients call that PUT requests several times.
- `POST`: POST requires POST client requests to update the corresponding resource with information provided by the client, or to create this resource if it does not exist. This operation is not idempotent.
- `DELETE`: This operation removes the resource forever. It is idempotent.

The format (JSON, XML, ...) used to return representations of resources is set in the media type of the server..

To handle success or errors issues, HTTP REST recommends using one of the HTTP status Codes. https://en.m.wikipedia.org/wiki/List_of_HTTP_status_codes
