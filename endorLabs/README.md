# ObjectDB with Redis

This is a simple Go project that demonstrates an ObjectDB implementation using Redis as the data store. The project provides an interface and a Redis-based implementation for storing, retrieving, and deleting objects.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Requirements

- Go 1.16 or higher
- Redis server

## Installation

1. Clone the repository:
```bash
git clone https://github.com/YG007/golang-examples
```

2. Change to the project directory:
```bash
cd endorLabs
```

3. Install the dependencies:
```bash
go mod download
```


4. Update the Redis server configuration:

Open the `main.go` file and modify the Redis connection options in the `NewRedisObjectDB` function to match your Redis server configuration.

5. Build and run the project:
```bash
go run main.go
```


## Usage

The project provides a simple command-line program that demonstrates the usage of the ObjectDB with Redis.

When you run the program, it will store and retrieve `Person` and `Animal` objects using the RedisObjectDB implementation. The program demonstrates various operations such as storing objects, retrieving objects by ID and name, listing objects, and deleting objects.


## Testing

The project includes unit tests for the ObjectDB implementation. You can run the tests by executing the following command:
```bash
go test
```


Make sure you have a Redis server running before running the tests.

## Contributing

Contributions to this project are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

Please make sure to follow the existing coding style and add appropriate tests for any new functionality.

## License

TBD





