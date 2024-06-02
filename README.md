# zota-project

Installation
To install this project, you can follow these steps:

Clone the repository:
```bash
git clone <repository_url>
```

Navigate into the project directory:

```bash
cd <project_director
```

Initialization
Before running the project, you need to set up some environment variables. You can export the following environment variables in your terminal:

```bash
export PORT=<port_number>
export MERCHANT_ID=<merchant_id>
export SECRET_KEY=<secret_key>
export CURRENCY=<currency>
export ENDPOINT_ID=<endpoint_id>
export BASE_URL=<base_url>
```

Make sure to replace <port_number>, <merchant_id>, <secret_key>, <currency>, <endpoint_id>, and <base_url> with the appropriate values for your setup.

Usage
Once you have set up the environment variables, you can start the project using the following Golang command:

```bash
go run main.go
```

This will start the project, and it will be accessible at the specified port.

Running Tests
To run the tests for this project, you can use the following command:

```bash
go test ./...
```

This command will run all the tests in the project.

Contributing
If you would like to contribute to this project, feel free to fork the repository and submit a pull request with your changes. Make sure to follow the contribution guidelines outlined in the repository.