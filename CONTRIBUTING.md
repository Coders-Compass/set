# Contributing to Go Set Implementation

First off, thank you for considering contributing to our Go Set implementation! This is a teaching and demonstration project that aims to help others understand both set theory and Go programming concepts.

## Why Read These Guidelines?

Following these guidelines helps to communicate that you respect the time of the developers maintaining and developing this open source project. In return, we'll reciprocate that respect by addressing your issues, assessing changes, and helping you finalize your pull requests.

## What We're Looking For

We love contributions that:

- Fix bugs in the implementation
- Improve performance of existing operations
- Add new set operations (with tests!)
- Improve documentation and examples
- Add educational content connecting implementation to set theory
- Fix typos or improve code readability
- Add more test cases

## What We're NOT Looking For

- Implementation of infinite sets (this package focuses on finite sets)
- Breaking changes to the existing API without discussion
- Changes that don't maintain our current performance characteristics
- Features that would be better suited as a separate package

## Ground Rules

### Expectations

- Create issues for any major changes and enhancements before implementing them
- Keep pull requests focused on a single change
- Write and update tests for all changes
- Maintain Go 1.23+ compatibility
- Follow existing code style and conventions
- Be welcoming to newcomers and encourage diverse contributions
- Keep performance characteristics in mind

### Technical Requirements

- Ensure all tests pass: `go test ./...`
- Maintain 100% test coverage for new code
- Run `go fmt` before committing
- Address all `go vet` warnings
- Pass all linter checks (`golangci-lint run`)
- Update documentation for any changed functionality

## Your First Contribution

Unsure where to begin? Here are some ways to start:

- **Beginner issues** - Look for issues tagged with `good-first-issue`
- **Documentation** - Help improve our godoc comments or examples
- **Tests** - Add test cases for edge scenarios
- **Bug fixes** - Look for issues tagged with `bug`

Never contributed to open source before? Check out:
- [Make your first open source contribution in four easy steps](https://github.com/readme/guides/first-oss-contribution)
- [Finding ways to contribute to open source on GitHub](https://docs.github.com/en/get-started/exploring-projects-on-github/finding-ways-to-contribute-to-open-source-on-github)

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/set.git
   cd set
   ```
3. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/Coders-Compass/set.git
   ```
4. Create a branch:
   ```bash
   git checkout -b my-feature
   ```

You might want to use SSH urls if you have your SSH key added to GitHub.

### Minor Changes

Small contributions (typo fixes, small documentation updates) can be submitted directly through GitHub's UI. For code changes, please submit a pull request.

## How to Report a Bug

> [!CAUTION]
> If you find a security vulnerability, do NOT open an issue. Email security@coderscompass.org instead.

When filing an issue, make sure to answer these questions:

1. What version of Go are you using (`go version`)?
2. What version of our package are you using?
3. What did you do?
4. What did you expect to see?
5. What did you see instead?

## How to Suggest a Feature

Open an issue that:
1. Clearly describes the feature
2. Explains why it would be valuable
3. Outlines possible implementation approaches
4. Considers performance implications
5. Describes how it fits with set theory concepts

## Code Review Process

The core team reviews Pull Requests on a regular basis. You can expect:

1. Initial response within 3 business days
2. Detailed review within a week
3. Follow-up reviews as needed

After feedback has been given, we expect responses within two weeks. After two weeks of inactivity, we may close the PR.

## Development Conventions

### Code Style

- Follow standard Go conventions
- Use meaningful variable names
- Keep functions focused and small
- Comment complex algorithms
- Use godoc format for package documentation

### Commit Messages

Format:
```
type(scope): description

[optional body]
[optional footer]
```

Types:
- feat: New feature
- fix: Bug fix
- docs: Documentation
- test: Adding tests
- perf: Performance improvements
- refactor: Code change that neither fixes a bug nor adds a feature

### Issue Labels

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `good-first-issue`: Good for newcomers
- `help-wanted`: Extra attention is needed
- `performance`: Performance-related issues
- `documentation`: Documentation improvements

## Community

- GitHub Issues: Primary communication channel
- Pull Requests: For code reviews and discussion
- [Website](https://coderscompass.org): Blog posts and articles
- Code of Conduct: Our [CoC](/CODE_OF_CONDUCT.md)

## Additional Resources

- [Set Theory Book](https://coderscompass.org/books/set-theory-for-beginners)
- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
