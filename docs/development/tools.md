# Development Tools

This document lists all external tools required to develop, build, test, and maintain the **TrueFlow** project.

> **Prerequisites**
>
> - Go 1.24+ (or the version specified in `go.mod`)
> - Git
> - PostgreSQL

---

# Required Tools

## Task

**Purpose**

Task is the project's command runner. It provides a simple and consistent interface for common development workflows
such as building, testing, formatting, database migrations, dependency management, and documentation generation.

All routine development tasks should be executed through the project's `Taskfile.yml`.

**Installation**

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

**Verify**

```bash
task --version
```

---

## Air

**Purpose**

Provides live reloading during development by automatically rebuilding and restarting the application whenever source
files change.

**Installation**

```bash
go install github.com/air-verse/air@latest
```

---

## GolangCI-Lint

**Purpose**

Runs multiple Go linters with a unified configuration to ensure code quality and consistency.

**Installation**

```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

**Verify**

```bash
golangci-lint version
```

---

## Goose

**Purpose**

Database migration tool used to manage PostgreSQL schema changes.

**Installation**

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

**Verify**

```bash
goose -version
```

---

## Swag

**Purpose**

Generates OpenAPI (Swagger) documentation directly from Go annotations.

**Installation**

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

**Verify**

```bash
swag --version
```
