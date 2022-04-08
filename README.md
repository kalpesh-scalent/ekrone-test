# ekrone-test


## Run Locally

### Pre-requisites
To successfully run the project on your local machine, you must have `Go` installed on your machine.

### Clone the project

```bash
  git clone https://github.com/kalpesh-scalent/ekrone-test.git
```

### Go to the project directory

```bash
  cd ekrone-test
```

### Environment Variables:

- To check the required environment variables, check th file `.env.example`.
- Do not worry about environment variables, if not set, it will use a default value and project will still run sucessfully

### Start the server
Once all basic setup is completed, you can run following command to start the server:
```bash
  make run-server
```
-----
## Demo
Server will be started on default port `:8080` on your localhost. Hit the following URL on any browser of your choice.
```bash
  http://localhost:8080/v1/projects/{number}
```

- `{number}` value for the URL variable must be a positive integer.
- A negative number or any non-numeric value will throw an error.
- Zero value for number or any other unexpected result will throw an `error` result. It will be shown in response.
- This is a REST API implementation. You can also make use of any REST client, such as Postman, Custom JS calls etc.

-----
## Tests
- Tests are written on Service layer only for demonstration purpose.
- A mock repository package is created for feeding data to tests.
- To run tests, use below mentioned command:
```bash
  make run-tests
```

-----
## Tech Stack

**Programming Language:** Go lang

**supporting packages:** Gorilla mux, uuid by Google, Logrus
