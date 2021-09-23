# GolangBackend

Running Tests
`$ go test -v ./...`

Test Coverage
`go test ./... -coverprofile a.out`
`go tool cover -html a.out`

To scan for vulnerability issues. This will produce a results.json file  in the root directory.

`$ go get github.com/securego/gosec/v2/cmd/gosec`

`$ gosec -fmt=json -out=results.json ./...`