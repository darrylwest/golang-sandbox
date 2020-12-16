.SILENT:
.PHONY: build dist edit format lint test watch run web

## help: this help file
help:
	@( echo "" && echo "Makefile targets..." && echo "" )
	@( cat Makefile | grep '^##' | sed -e 's/##/ -/' | sort && echo "" )

