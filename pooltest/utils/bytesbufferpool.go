package utils

import (
	"bytes"
	"sync"
)

type bytesBufferPool struct {
	*sync.Pool
}

var (
	Size = 1024
)

func NewBytesBufferPool(cap int) *bytesBufferPool {
	Size = cap
	return &bytesBufferPool{
		&sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, cap))
		}},
	}
}

func (bp *bytesBufferPool) Get() *bytes.Buffer {
	return (bp.Pool.Get()).(*bytes.Buffer)
}

func (bp *bytesBufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	//对象太大，需要进行处理
	if b.Cap() > Size {
		b = bytes.NewBuffer(make([]byte, 0, Size))
	}
	bp.Pool.Put(b)
}
