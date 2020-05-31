# note: call scripts from /scripts
.PHONY: all clean-build

.PHONY: all
all: clean-build

.PHONY: clean-build
clean-build:
	@ ./scripts/clean.sh
	@ ./scripts/build.sh
	@echo ">>>> Done..."