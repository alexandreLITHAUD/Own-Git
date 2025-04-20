# Own-Git 🌳

![Build](https://github.com/alexandreLITHAUD/Own-Git/actions/workflows/launch-tests.yaml/badge.svg)
![Lint](https://github.com/alexandreLITHAUD/Own-Git/actions/workflows/lint-go-code.yaml/badge.svg)
![Tests](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/alexandreLITHAUD/3aff3ab94739bdcdd6a9640f0150eeda/raw/tests.json)
![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/alexandreLITHAUD/3aff3ab94739bdcdd6a9640f0150eeda/raw/coverage.json)

> A simple git copycat in go using cobra to challenge myself and further my understanding of git 🌳

---

## 🚀 Overview

**Own-Git** is a minimalist reimplementation of Git built in Go, designed as a hands-on challenge to understand Git's inner workings and improve my skills with:
- 🧠 Go & [Cobra](https://github.com/spf13/cobra) CLI
- 🧪 Git internals
- 🛠️ CLI architecture & testing
- 🚧 CI/CD pipelines and Go toolchain integration

---

## 📦 Installation

You can install the latest version with Go:

```bash
go install github.com/alexandreLITHAUD/Own-Git@latest
```

Or download a release binary from the [Releases](https://github.com/alexandreLITHAUD/Own-Git/releases) page.

---

## 🧑‍💻 Usage

```bash
own init         # Initialize a new repository
own add file.txt # Stage a file
own commit -m "Message"  # Commit with a message
own status       # Check current status
own version      # Show the current version
```

For full help:

```bash
own help
```

---

## 🧰 Available Commands

| Command     | Description                |
|------------|----------------------------|
| `init`     | Initialize a new repository |
| `add`      | Stage files for commit      |
| `commit`   | Commit changes              |
| `status`   | Show the repo status        |
| `version`  | Show version information    |

> More commands are being added as the project evolves!

---

## 📄 Documentation

The documentation is automatically generated with [Hugo](https://gohugo.io) and hosted on GitHub Pages:

🌐 [https://alexandrelithaud.github.io/Own-Git](https://alexandrelithaud.github.io/Own-Git)

---

## 📈 CI/CD & Tooling

- ✅ Tests & coverage with GitHub Actions
- 🧹 Code linting and formatting
- 🛡️ Security analysis with `gosec`
- 📄 Auto-generated CLI docs with `cobra/doc`
- 🧪 Benchmarking with Go's `testing` package

---

## 🤝 Contributing

Contributions are welcome!  
Feel free to fork, explore, and open a pull request 🙌

---

## 📜 License

Own-Git is licensed under the [MIT License](./LICENSE).
