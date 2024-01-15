# User Handbook

## Build
`go build -buildmode=plugin -o godoxwithregexp.so cmd/godox-with-regexp/main.go`

## Linter Settings
```yaml
linters-settings:
  custom:
    godoxwithregexp:
      # Path is required
      path: godoxwithregexp.so
      # Description is optional
      description: Tool for detection of TODO and other comment keywords with regular expression.
      # Original-url is optional, and is only used for documentation purposes.
      original-url: github.com/SilverdewBaker/godox-with-regexp

linters:
  enable:
    - godoxwithregexp
```

## Run
`golangci-lint run`
