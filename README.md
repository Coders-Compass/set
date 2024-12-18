# Go Set Implementation

[![Go Code Quality & Tests](https://github.com/Coders-Compass/set/workflows/Go%20Code%20Quality%20&%20Tests/badge.svg)](https://github.com/Coders-Compass/set/actions)
[![Go Reference](https://pkg.go.dev/badge/coderscompass.org/set.svg)](https://pkg.go.dev/coderscompass.org/set)
[![Go Report Card](https://goreportcard.com/badge/coderscompass.org/set)](https://goreportcard.com/report/coderscompass.org/set)
[![Security Policy](https://img.shields.io/badge/security-policy-brightgreen.svg)](SECURITY.md)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

## What this project does

This package provides a generic interface and an in-memory implementation of finite sets in Go. It serves both as a practical utility and as a teaching tool for understanding set theory concepts through code.

Key features:
- Generic implementation supporting any comparable type
- O(1) insertions using hash-based storage
- Efficient set operations (intersection, equality checking)
- Comprehensive test coverage
- Clear documentation connecting theory to implementation

## Why it's useful

Our set implementation helps you:
- Manage collections of unique elements efficiently
- Perform set operations with guaranteed performance characteristics
- Learn set theory through practical code examples
- See how mathematical concepts translate to Go code

## Getting started

### Installation

```bash
go get coderscompass.org/set
```

### Basic Usage

```go
package main

import (
	"fmt"

	"coderscompass.org/set"
)

func main() {
	// Create sets
	hobbits := set.NewHashSet[string]()
	fellowship := set.NewHashSet[string]()

	// Add elements
	hobbits.Insert("Frodo")
	hobbits.Insert("Sam")
	fellowship.Insert("Frodo")
	fellowship.Insert("Gandalf")

	// Find intersection
	hobbitFellows := hobbits.Intersection(fellowship)
	// Result contains {"Frodo"}

	expected := set.NewHashSet[string]()
	expected.Insert("Frodo")

	if !hobbitFellows.Equals(expected) {
		fmt.Println("Expected hobbitFellows to be equal to expected")
	} else {
		fmt.Println("hobbitFellows is equal to expected")
	}
}

```

### Performance Guarantees

| **Operation**    | **Time Complexity** |
|------------------|---------------------|
| Insert           | O(1)                |
| Intersection     | O(m + n)            |
| Equals           | O(min(n, m))        |

## Getting help

- [API Documentation](https://pkg.go.dev/coderscompass.org/set)
- [GitHub Issues](https://github.com/Coders-Compass/set/issues)
- [Contributing Guide](CONTRIBUTING.md)
- [Security Policy](SECURITY.md)

## Project maintenance

### Current Status
- Version: v0.1.1
- Go Version: >= 1.23
- Maintenance: Active

### Contributing
We welcome contributions! See our [Contributing Guide](CONTRIBUTING.md) for details on:
- Code style and standards
- Development setup
- Test requirements
- Pull request process

### Security
For security concerns, please review our [Security Policy](SECURITY.md).

## Learning Resources

- [Set Theory Book](https://coderscompass.org/books/set-theory-for-beginners?utm_campaign=presentation&utm_source=github): Comprehensive guide to set theory concepts
- [pkg.go.dev Documentation](https://pkg.go.dev/coderscompass.org/set): API reference

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please [contact us](https://coderscompass.org/contact).
