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

#### How Bearer tokens work within this project

##### The Login endpoint
- User hits the `/auth/login` endpoint with a username and password
- Login handler gets called, which then calls the `VerifyCreditials()` method of the AuthService Interface
- `VerifiyCredentials()` calls the `CheckCredentials()` method of the `AuthRepository` Interface 
- `CheckCredentials()` queries the DB to get the `username`, `customer_id`, `role`, and list of `account_ids` (if any) for the user (assuming the username and password is correct) and returns the user.
- The user then gets passed back to the service interface `VerifyCredentials()` which generates the mapClaims and the token (MapClaims is just a goLang map with Claims- [the admin or user info])
- Once mapClaims have been generated the token is generated by calling `jwt.NewWithClaims()` and passing in the signing method
- Then the token is signed using the secret and returned to the login handler

##### The Verify endpoint



##### Types of Users
- There are two types of User roles: admin and user
- Admin can access any account
- User can only access user specific accounts 


