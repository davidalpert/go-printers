PROJECTNAME=go-printers

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# go versioning flags
VERSION=$(shell sbot get version)

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)

# ---------------------- targets -------------------------------------

.PHONY: default
default: help

.PHONY: cit
cit: clean vale test-unit ## clean build and test-unit

.PHONY: version
version: ## show current version
	echo ${VERSION}

.PHONY: clean
clean: ## clean build output
	rm -rf ./bin

.PHONY: vale
vale: ## run linting rules against markdown files
ifeq ("$(GITHUB_ACTIONS)","true")
	echo "GITHUB_ACTIONS is true; assuming vale has been run by a previous step"
else
	vale README.md CONTRIBUTING.md # we don't valedate CHANGELOG.md as that reflects historical commit summaries
endif

.PHONY: guard
guard: ## run ruby-guard
	bundle exec guard

.PHONY: gen
gen: ## invoke go generate
	@CGO_ENABLED=1 go generate ./...

test: test-unit ## run unit tests

.PHONY: test-unit
test-unit: ## run unit tests
	go test -v ./...

.PHONY: doctor
doctor: ## run doctor.sh to sort out development dependencies
	./.tools/doctor.sh

.PHONY: changelog
changelog: ## Generate/update CHANGELOG.md
	git-chglog --output CHANGELOG.md

.PHONY: preview-release-notes
preview-release-notes: ## preview release notes (generates RELEASE_NOTES.md)
	git-chglog --output RELEASE_NOTES.md --template .chglog/RELEASE_NOTES.tpl.md "v$(shell sbot get version)"

.PHONY: preview-release
preview-release: preview-release-notes ## preview release (using goreleaser --snapshot)
	goreleaser release --snapshot --rm-dist --release-notes RELEASE_NOTES.md

eq = $(and $(findstring $(1),$(2)),$(findstring $(2),$(1)))

.PHONY: tag-release
tag-release:
	$(if $(call eq,0,$(shell git diff-files --quiet; echo $$?)),, \
		$(error There are unstaged changes; clean your working directory before releasing.) \
	)
	$(if $(call eq,0,$(shell git diff-index --quiet --cached HEAD --; echo $$?)),, \
		$(error There are uncomitted changes; clean your working directory before releasing.) \
	)
	$(eval next_version := $(shell sbot predict version --mode ${BUMP_TYPE}))
	# echo "Current Version: ${VERSION}"
	# echo "   Next Version: ${next_version}"
ifdef FAST
	$(MAKE) test-unit VERSION=$(next_version)
else
	$(MAKE) cit VERSION=$(next_version)
endif
	git-chglog --next-tag v$(next_version) --output CHANGELOG.md
	git add -f CHANGELOG.md
	git commit --message "release notes for v$(next_version)"
	sbot release version --mode ${BUMP_TYPE}
	git show --no-patch --format=short v$(next_version)

SEMVER_TYPES := major minor patch
BUMP_TARGETS := $(addprefix release-,$(SEMVER_TYPES))
.PHONY: $(BUMP_TARGETS)
$(BUMP_TARGETS): ## bump version
	$(eval BUMP_TYPE := $(strip $(word 2,$(subst -, ,$@))))
	$(MAKE) tag-release BUMP_TYPE=$(BUMP_TYPE)

.PHONY: help
help: Makefile
	@echo
	@echo " ${PROJECTNAME} ${VERSION} - available targets:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	printf "\033[36m%-30s\033[0m %s\n" '----------' '------------------'
	@echo $(BUMP_TARGETS) | tr ' ' '\n' | sort | sed -E 's/((.+)\-(.+))/\1: ## \2 \3 version/' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
