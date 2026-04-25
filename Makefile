TOOLS := harness
DIST  := dist

.PHONY: all build test clean $(TOOLS)

all: build

build: $(TOOLS)

$(TOOLS):
	$(MAKE) -C $@ build

test:
	@set -e; for tool in $(TOOLS); do $(MAKE) -C $$tool test; done

clean:
	rm -rf $(DIST)/*
	@touch $(DIST)/.gitkeep
