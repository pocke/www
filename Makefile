dep:
	go install github.com/goreleaser/goreleaser/v2@latest

release:
	goreleaser

snapshot:
	goreleaser --snapshot

clean:
	rm -r dist
