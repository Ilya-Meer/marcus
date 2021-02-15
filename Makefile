BINDIR    = /usr/local/bin
EXEC      = marcus
BIN		  = $(BINDIR)/$(EXEC)

# Install executable
install:
	cp $(EXEC) $(BINDIR)
	chmod 755 $(BIN)

# Uninstall executable
uninstall:
	rm -f $(BIN)

# Build executable
build:
	go build