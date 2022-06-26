## Compared to the original repo

Add `zh-tw` 繁體中文

Change `zh` to `zh-cn`

## ISO 639-1

[![Go Reference](https://pkg.go.dev/badge/github.com/emvi/iso-639-1?status.svg)](https://pkg.go.dev/github.com/iyear/iso-639-1?status)
[![Go Report Card](https://goreportcard.com/badge/github.com/emvi/iso-639-1)](https://goreportcard.com/report/github.com/iyear/iso-639-1)

List of all ISO 639-1 language names, native names and two character codes as well as functions for convenient access.
The lists of all names and codes (`Codes`, `Names`, `NativeNames`, `Languages`) are build in the init function for quick read access. 
For full documentation please read the Godocs.

### Installation

```
go get github.com/iyear/iso-639-1
```

### Example

```
fmt.Println(iso6391.Codes)          // print all codes
fmt.Println(iso6391.Names)          // print all names
fmt.Println(iso6391.NativeNames)    // print all native names
fmt.Println(iso6391.Languages)      // print all language objects {Code, Name, NativeName}

fmt.Println(iso6391.FromCode("en"))             // prints {Code: "en", Name: "English", NativeName: "English"}
fmt.Println(iso6391.Name("en"))                 // prints "English"
fmt.Println(iso6391.NativeName("zh"))           // prints "中文"
fmt.Println(iso6391.CodeFromName("English"))    // prints "en"
fmt.Println(iso6391.ValidCode("en"))            // prints true
// ... see Godoc for more functions
```

## Contribute

[See CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT
