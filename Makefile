EXEC      = marcus
BINDIR    = /usr/local/bin
BIN		  = $(BINDIR)/$(EXEC)

# Install executable
install:
	cp bin/$(EXEC) $(BINDIR)
	chmod 755 $(BIN)

# Uninstall executable
uninstall:
	rm -f $(BIN)

# Build executable
build:
	go build -o bin/marcus ./src