BINARY := gotestx
DIST := dist
RELEASE_DIR := releases

VERSION ?=

.PHONY: build test clean dist checksums changelog release-notes prepare-release publish release

build:
	go build -o $(BINARY) .

test:
	go test ./...

clean:
	rm -rf $(DIST)
	rm -f $(BINARY)

dist: clean
	mkdir -p $(DIST)

	GOOS=linux GOARCH=amd64 go build -o $(DIST)/$(BINARY)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o $(DIST)/$(BINARY)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o $(DIST)/$(BINARY)-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o $(DIST)/$(BINARY)-windows-amd64.exe

checksums: dist
	cd $(DIST) && sha256sum * > checksums.txt

changelog:
	git cliff --config cliff.toml --output CHANGELOG.md

release-notes:
	mkdir -p $(RELEASE_DIR)
	git cliff \
		--config cliff.toml \
		--tag $(VERSION) \
		--template .cliff/release.tpl \
		--output $(RELEASE_DIR)/release-notes-$(VERSION).md

prepare-release:
	@if [ -z "$(VERSION)" ]; then \
		echo "ERROR: VERSION is required"; \
		echo "Usage: make prepare-release VERSION=vX.Y.Z"; \
		exit 1; \
	fi

	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "ERROR: git working tree is dirty"; \
		echo "Please commit or stash changes before preparing a release"; \
		exit 1; \
	fi

	@if git rev-parse "$(VERSION)" >/dev/null 2>&1; then \
		echo "ERROR: tag $(VERSION) already exists"; \
		exit 1; \
	fi

	@echo "Preparing release $(VERSION)"

	$(MAKE) changelog
	$(MAKE) release-notes VERSION=$(VERSION)

	git add CHANGELOG.md $(RELEASE_DIR)/release-notes-$(VERSION).md
	git commit -S -m "docs(release): prepare $(VERSION)"

	@echo ""
	@echo "Release prepared."
	@echo "Next steps:"
	@echo "  git push"
	@echo "  git tag -s $(VERSION) -m \"GoTestX $(VERSION)\""
	@echo "  git push origin $(VERSION)"
	@echo "  make publish VERSION=$(VERSION)"

publish: dist checksums
	@if [ -z "$(VERSION)" ]; then \
		echo "ERROR: VERSION is required"; \
		exit 1; \
	fi

	@echo "Publishing release $(VERSION)"

	gh release create $(VERSION) \
		--title "GoTestX $(VERSION)" \
		--notes-file $(RELEASE_DIR)/release-notes-$(VERSION).md \
		$(DIST)/*

release: prepare-release publish