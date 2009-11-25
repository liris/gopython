include $(GOROOT)/src/Make.$(GOARCH)

TARG=py
CGOFILES=\
        gopywrapper.go

#CGO_LDFLAGS+=gopywrapper.o -llibpython
CGO_LDFLAGS+=-lpython -u _PyMac_Error /Library/Frameworks/Python.framework/Versions/2.6/Python
CGO_CFLAGS=-I/usr/include/python2.6/ $(CGO_FLAGS)
CFLAGS=-I/usr/include/python2.6/

LDFLAGS=-u _PyMac_Error /Library/Frameworks/Python.framework/Versions/2.6/Python

include $(GOROOT)/src/Make.pkg

%: gopywrapper.o install %.go
	$(GC)  $*.go
	$(LD) -o $@ $*.$O

gopywrapper.o: gopywrapper.c
	gcc -fPIC -O2 -o gopywrapper.o -c gopywrapper.c

gopytest: install gopytest.go
	$(GC) gopytest.go
	$(LD) -o $@ gopytest.$O 

