# Contributing

I do not consider this project to be finished. It is and usable for production but improvements are always good :)

## Create an Issue First

Your PR is more likely to be accepted if you open an issue first.

* Please open an issue first
* Create a pull request
* Document your change
* go vet, go imports, go fmt and all other usual suspects first
* Be patient, it will maybe take 1h until I have time to check your change ;)

###Submission Checklist

- [ ] `gofmt -s -l .` produces no output
- [ ] `go vet .` produces no output
- [ ] `golint ./...` produces no output
- [ ] `misspell .` produces no output
- [ ] `ineffassign .` produces no output
- [ ] `go test -race` passes
