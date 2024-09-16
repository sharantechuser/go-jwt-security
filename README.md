# <p style="color:teal">go-jwt-security</p>
Json Web Token(JWT) implementation in Go

This is implemented by using package dgrijalva. See [dgrijalva/jwt-go#216](https://github.com/dgrijalva/jwt-go/issues/216)  for more detail.



# <p style="color:teal">Prerequisites</p>
1. **Go installed**: Ensure you have Go installed on your machine. In this example Go1.22 been used
2. **JWT Library**: Use a JWT library for Go. github.com/dgrijalva/jwt-go is a popular choice.
3. **gin Library**: Use a gin library for Go. github.com/gin-gonic/gin is a popular choice.

# <p style="color:teal">Dependencies</p>
```sh
    #jwt package
    go get github.com/dgrijalva/jwt-go

    #http framework package
    go get github.com/gin-gonic/gin
```


# <p style="color:teal">Build & Run the application</p>
The application run with 8000 port by default
```sh
    #Make command
    make build

    #http framework package
    ./go-jwt-security -port=8000
```


# <p style="color:teal">Testing the Application</p>
```sh

    #Login to get a token:
    curl -X POST http://localhost:8000/login -d '{"username": "Test", "password": "pass"}' -H "Content-Type: application/json"

    #Access protected route
    curl http://localhost:8000/authorise -H "Authorization: YOUR_JWT_TOKEN"



```