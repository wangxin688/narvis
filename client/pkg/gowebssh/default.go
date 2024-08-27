package gowebssh

import (
	"io"
	"log"
	"time"
)

var (
	DefaultTerm       = TermXterm
	DefaultTimeout    = time.Second * 15
	DefaultLogger     = log.New(io.Discard, "[webssh] ", log.Ltime|log.Ldate)
	DefaultBufferSize = uint32(8192)
)
