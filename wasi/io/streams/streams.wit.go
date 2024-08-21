// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package streams represents the imported interface "wasi:io/streams@0.2.0".
//
// WASI I/O is an I/O abstraction API which is currently focused on providing
// stream types.
//
// In the future, the component model is expected to add built-in stream types;
// when it does, they are expected to subsume this API.
package streams

import (
	ioerror "github.com/mossaka/wasi-http-with-wasm-tools-go/wasi/io/error"
	"github.com/mossaka/wasi-http-with-wasm-tools-go/wasi/io/poll"
	"github.com/ydnar/wasm-tools-go/cm"
)

// StreamError represents the imported variant "wasi:io/streams@0.2.0#stream-error".
//
// An error for input-stream and output-stream operations.
//
//	variant stream-error {
//		last-operation-failed(error),
//		closed,
//	}
type StreamError cm.Variant[uint8, ioerror.Error, ioerror.Error]

// StreamErrorLastOperationFailed returns a [StreamError] of case "last-operation-failed".
//
// The last operation (a write or flush) failed before completion.
//
// More information is available in the `error` payload.
func StreamErrorLastOperationFailed(data ioerror.Error) StreamError {
	return cm.New[StreamError](0, data)
}

// LastOperationFailed returns a non-nil *[ioerror.Error] if [StreamError] represents the variant case "last-operation-failed".
func (self *StreamError) LastOperationFailed() *ioerror.Error {
	return cm.Case[ioerror.Error](self, 0)
}

// StreamErrorClosed returns a [StreamError] of case "closed".
//
// The stream is closed: no more input will be accepted by the
// stream. A closed output-stream will return this error on all
// future operations.
func StreamErrorClosed() StreamError {
	var data struct{}
	return cm.New[StreamError](1, data)
}

// Closed returns true if [StreamError] represents the variant case "closed".
func (self *StreamError) Closed() bool {
	return self.Tag() == 1
}

// InputStream represents the imported resource "wasi:io/streams@0.2.0#input-stream".
//
// An input bytestream.
//
// `input-stream`s are *non-blocking* to the extent practical on underlying
// platforms. I/O operations always return promptly; if fewer bytes are
// promptly available than requested, they return the number of bytes promptly
// available, which could even be zero. To wait for data to be available,
// use the `subscribe` function to obtain a `pollable` which can be polled
// for using `wasi:io/poll`.
//
//	resource input-stream
type InputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "input-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self InputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_InputStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]input-stream
//go:noescape
func wasmimport_InputStreamResourceDrop(self0 uint32)

// BlockingRead represents the imported method "blocking-read".
//
// Read bytes from a stream, after blocking until at least one byte can
// be read. Except for blocking, behavior is identical to `read`.
//
//	blocking-read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) BlockingRead(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingRead((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-read
//go:noescape
func wasmimport_InputStreamBlockingRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

// BlockingSkip represents the imported method "blocking-skip".
//
// Skip bytes from a stream, after blocking until at least one byte
// can be skipped. Except for blocking behavior, identical to `skip`.
//
//	blocking-skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) BlockingSkip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.blocking-skip
//go:noescape
func wasmimport_InputStreamBlockingSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Read represents the imported method "read".
//
// Perform a non-blocking read from the stream.
//
// When the source of a `read` is binary data, the bytes from the source
// are returned verbatim. When the source of a `read` is known to the
// implementation to be text, bytes containing the UTF-8 encoding of the
// text are returned.
//
// This function returns a list of bytes containing the read data,
// when successful. The returned list will contain up to `len` bytes;
// it may return fewer than requested, but not more. The list is
// empty when no bytes are available for reading at this time. The
// pollable given by `subscribe` will be ready when more bytes are
// available.
//
// This function fails with a `stream-error` when the operation
// encounters an error, giving `last-operation-failed`, or when the
// stream is closed, giving `closed`.
//
// When the caller gives a `len` of 0, it represents a request to
// read 0 bytes. If the stream is still open, this call should
// succeed and return an empty list, or otherwise fail with `closed`.
//
// The `len` parameter is a `u64`, which could represent a list of u8 which
// is not possible to allocate in wasm32, or not desirable to allocate as
// as a return value by the callee. The callee may return a list of bytes
// less than `len` in size while more bytes are available for reading.
//
//	read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) Read(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamRead((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.read
//go:noescape
func wasmimport_InputStreamRead(self0 uint32, len0 uint64, result *cm.Result[cm.List[uint8], cm.List[uint8], StreamError])

// Skip represents the imported method "skip".
//
// Skip bytes from a stream. Returns number of bytes skipped.
//
// Behaves identical to `read`, except instead of returning a list
// of bytes, returns the number of bytes consumed from the stream.
//
//	skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) Skip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.skip
//go:noescape
func wasmimport_InputStreamSkip(self0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Subscribe represents the imported method "subscribe".
//
// Create a `pollable` which will resolve once either the specified stream
// has bytes available to read or the other end of the stream has been
// closed.
// The created `pollable` is a child resource of the `input-stream`.
// Implementations may trap if the `input-stream` is dropped before
// all derived `pollable`s created with this function are dropped.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self InputStream) Subscribe() (result poll.Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_InputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]input-stream.subscribe
//go:noescape
func wasmimport_InputStreamSubscribe(self0 uint32) (result0 uint32)

// OutputStream represents the imported resource "wasi:io/streams@0.2.0#output-stream".
//
// An output bytestream.
//
// `output-stream`s are *non-blocking* to the extent practical on
// underlying platforms. Except where specified otherwise, I/O operations also
// always return promptly, after the number of bytes that can be written
// promptly, which could even be zero. To wait for the stream to be ready to
// accept data, the `subscribe` function to obtain a `pollable` which can be
// polled for using `wasi:io/poll`.
//
//	resource output-stream
type OutputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "output-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [resource-drop]output-stream
//go:noescape
func wasmimport_OutputStreamResourceDrop(self0 uint32)

// BlockingFlush represents the imported method "blocking-flush".
//
// Request to flush buffered output, and block until flush completes
// and stream is ready for writing again.
//
//	blocking-flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingFlush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamBlockingFlush((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-flush
//go:noescape
func wasmimport_OutputStreamBlockingFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// BlockingSplice represents the imported method "blocking-splice".
//
// Read from one stream and write to another, with blocking.
//
// This is similar to `splice`, except that it blocks until the
// `output-stream` is ready for writing, and the `input-stream`
// is ready for reading, before performing the `splice`.
//
//	blocking-splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingSplice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-splice
//go:noescape
func wasmimport_OutputStreamBlockingSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// BlockingWriteAndFlush represents the imported method "blocking-write-and-flush".
//
// Perform a write of up to 4096 bytes, and then flush the stream. Block
// until all of these operations are complete, or an error occurs.
//
// This is a convenience wrapper around the use of `check-write`,
// `subscribe`, `write`, and `flush`, and is implemented with the
// following pseudo-code:
//
//	let pollable = this.subscribe();
//	while !contents.is_empty() {
//	// Wait for the stream to become writable
//	pollable.block();
//	let Ok(n) = this.check-write(); // eliding error handling
//	let len = min(n, contents.len());
//	let (chunk, rest) = contents.split_at(len);
//	this.write(chunk  );            // eliding error handling
//	contents = rest;
//	}
//	this.flush();
//	// Wait for completion of `flush`
//	pollable.block();
//	// Check for any errors that arose during `flush`
//	let _ = this.check-write();         // eliding error handling
//
//	blocking-write-and-flush: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteAndFlush(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamBlockingWriteAndFlush((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteAndFlush(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// BlockingWriteZeroesAndFlush represents the imported method "blocking-write-zeroes-and-flush".
//
// Perform a write of up to 4096 zeroes, and then flush the stream.
// Block until all of these operations are complete, or an error
// occurs.
//
// This is a convenience wrapper around the use of `check-write`,
// `subscribe`, `write-zeroes`, and `flush`, and is implemented with
// the following pseudo-code:
//
//	let pollable = this.subscribe();
//	while num_zeroes != 0 {
//	// Wait for the stream to become writable
//	pollable.block();
//	let Ok(n) = this.check-write(); // eliding error handling
//	let len = min(n, num_zeroes);
//	this.write-zeroes(len);         // eliding error handling
//	num_zeroes -= len;
//	}
//	this.flush();
//	// Wait for completion of `flush`
//	pollable.block();
//	// Check for any errors that arose during `flush`
//	let _ = this.check-write();         // eliding error handling
//
//	blocking-write-zeroes-and-flush: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteZeroesAndFlush(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingWriteZeroesAndFlush((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.blocking-write-zeroes-and-flush
//go:noescape
func wasmimport_OutputStreamBlockingWriteZeroesAndFlush(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])

// CheckWrite represents the imported method "check-write".
//
// Check readiness for writing. This function never blocks.
//
// Returns the number of bytes permitted for the next call to `write`,
// or an error. Calling `write` with more bytes than this function has
// permitted will trap.
//
// When this function returns 0 bytes, the `subscribe` pollable will
// become ready when this function will report at least 1 byte, or an
// error.
//
//	check-write: func() -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) CheckWrite() (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamCheckWrite((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.check-write
//go:noescape
func wasmimport_OutputStreamCheckWrite(self0 uint32, result *cm.Result[uint64, uint64, StreamError])

// Flush represents the imported method "flush".
//
// Request to flush buffered output. This function never blocks.
//
// This tells the output-stream that the caller intends any buffered
// output to be flushed. the output which is expected to be flushed
// is all that has been passed to `write` prior to this call.
//
// Upon calling this function, the `output-stream` will not accept any
// writes (`check-write` will return `ok(0)`) until the flush has
// completed. The `subscribe` pollable will become ready when the
// flush has completed and the stream can accept more writes.
//
//	flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Flush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamFlush((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.flush
//go:noescape
func wasmimport_OutputStreamFlush(self0 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// Splice represents the imported method "splice".
//
// Read from one stream and write to another.
//
// The behavior of splice is equivelant to:
// 1. calling `check-write` on the `output-stream`
// 2. calling `read` on the `input-stream` with the smaller of the
// `check-write` permitted length and the `len` provided to `splice`
// 3. calling `write` on the `output-stream` with that read data.
//
// Any error reported by the call to `check-write`, `read`, or
// `write` ends the splice and reports that error.
//
// This function returns the number of bytes transferred; it may be less
// than `len`.
//
//	splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) Splice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.splice
//go:noescape
func wasmimport_OutputStreamSplice(self0 uint32, src0 uint32, len0 uint64, result *cm.Result[uint64, uint64, StreamError])

// Subscribe represents the imported method "subscribe".
//
// Create a `pollable` which will resolve once the output-stream
// is ready for more writing, or an error has occured. When this
// pollable is ready, `check-write` will return `ok(n)` with n>0, or an
// error.
//
// If the stream is closed, this pollable is always ready immediately.
//
// The created `pollable` is a child resource of the `output-stream`.
// Implementations may trap if the `output-stream` is dropped before
// all derived `pollable`s created with this function are dropped.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self OutputStream) Subscribe() (result poll.Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_OutputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.subscribe
//go:noescape
func wasmimport_OutputStreamSubscribe(self0 uint32) (result0 uint32)

// Write represents the imported method "write".
//
// Perform a write. This function never blocks.
//
// When the destination of a `write` is binary data, the bytes from
// `contents` are written verbatim. When the destination of a `write` is
// known to the implementation to be text, the bytes of `contents` are
// transcoded from UTF-8 into the encoding of the destination and then
// written.
//
// Precondition: check-write gave permit of Ok(n) and contents has a
// length of less than or equal to n. Otherwise, this function will trap.
//
// returns Err(closed) without writing if the stream has closed since
// the last call to check-write provided a permit.
//
//	write: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Write(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamWrite((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write
//go:noescape
func wasmimport_OutputStreamWrite(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[StreamError, struct{}, StreamError])

// WriteZeroes represents the imported method "write-zeroes".
//
// Write zeroes to a stream.
//
// This should be used precisely like `write` with the exact same
// preconditions (must use check-write first), but instead of
// passing a list of bytes, you simply pass the number of zero-bytes
// that should be written.
//
//	write-zeroes: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) WriteZeroes(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamWriteZeroes((uint32)(self0), (uint64)(len0), &result)
	return
}

//go:wasmimport wasi:io/streams@0.2.0 [method]output-stream.write-zeroes
//go:noescape
func wasmimport_OutputStreamWriteZeroes(self0 uint32, len0 uint64, result *cm.Result[StreamError, struct{}, StreamError])
