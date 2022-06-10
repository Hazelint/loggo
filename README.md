# Loggo
A logging package for your go projects

## Getting Started

### Installing
Run this command to get the package in your project:
```
go get github.com/Hazelint/loggo
```

### Example
Here is a example on how to use this package:
```
message := "Test message"

loggo.Info(message)
```
This will be the output in the `go.log` file:

```
[2022-06-10 00:11:38] Info: Test message
```

## Version History

* 0.1.2
    * Added new logging levels
    * Updated documentation
* 0.1.1
    * Changes to `.mod` file
* 0.1.0
    * Initial Release

## License

[![GitHub license](https://img.shields.io/github/license/Hazelint/loggo)](https://github.com/Hazelint/loggo/blob/main/LICENSE)

