dep:
	go install github.com/goreleaser/goreleaser@latest

release:
	goreleaser

snapshot:
	goreleaser --snapshot

clean:
	rm -r dist
