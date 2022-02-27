# goLangAuth
Authentication MicroService for [GoLangPracticeProject](https://github.com/jesserahman/goLangPracticeProject)
This microservice will be used for logging into the application and returning a bearer token to the user.
This microservice will also be verifying the validity of the bearer token when trying to hit any endpoint with the [GoLangPracticeProject](https://github.com/jesserahman/goLangPracticeProject)

## Running the application
- `go run main.go`

### Getting auth token
Go to `localhost:{port-in-env-file}/`

- POST `/auth/login` to get Bearer token

#### Hitting Verify endpoint 
- GET `/auth/verify`
Example `http://localhost:8081/auth/verify?routeName=GetAllCustomers&token=aaaa.bbbb.cccc`

Sample JSON examples:
<h4> Customers </h4>

``` 
POST /customer
{
    "username" : "John Doe",
    "password" : "SuperSecretPassword",
}
```


