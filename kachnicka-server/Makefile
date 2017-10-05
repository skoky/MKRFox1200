VERSION = $$(git rev-parse --abbrev-ref HEAD 2> /dev/null | rev | cut -d/ -f1 | rev)
GOSDKBIN = $$(dirname -- "$$(readlink $$(which gcloud))")

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: dev
dev:
	$(GOSDKBIN)/dev_appserver.py app.yaml
