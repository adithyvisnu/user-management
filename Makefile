# note: call scripts from /scripts
.PHONY: all clean-build run

.PHONY: all
all: run

.PHONY: clean-build
clean-build:
	@ ./scripts/clean.sh
	@ ./scripts/build.sh
	@echo ">>>> Done..."

.PHONY: run
run:
	@ make clean-build
	@ ./application