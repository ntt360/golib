/**
* @file writer.go
* @brief log_writer
* @author ligang
* @version 1.0
* @date 2014-10-14
 */

package bizlog

import (
	"bufio"
	//     "fmt"
	"os"
)

/**
* @brief interface log_writer
 */
type i_log_writer interface {
	writeln(s string) (int, error)
	flush()
	free()
}

/**
* @brief get log writer
*
* @param string
* @param int
*
* @return i_log_writer
 */
func newLogWriter(path string, bufsize int) i_log_writer {
	if 0 == bufsize {
		return newNoBufLogWriter(path)
	}
	return newBufLogWriter(path, bufsize)
}

func openLogFile(path string) *os.File {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return f
}

/**
* @name buf_log_writer
* @{ */
type t_buf_log_writer struct {
	f *os.File
	w *bufio.Writer
}

/**
* @brief get new buf_log_writer
*
* @param string
* @param int
*
* @return
 */
func newBufLogWriter(path string, bufsize int) *t_buf_log_writer {
	f := openLogFile(path)
	return &t_buf_log_writer{
		f: f,
		w: bufio.NewWriterSize(f, bufsize),
	}
}

func (writer *t_buf_log_writer) writeln(s string) (int, error) {
	return writer.w.WriteString(s + "\n")
}

func (writer *t_buf_log_writer) flush() {
	writer.w.Flush()
}

func (writer *t_buf_log_writer) free() {
	writer.w.Flush()
	writer.f.Close()
	writer.w = nil
}

/**  @} */

/**
* @name no_buf_log_writer
* @{ */

type t_nobuf_log_writer struct {
	f *os.File
}

/**
* @brief get new nobuf_log_writer
*
* @param string
*
* @return
 */
func newNoBufLogWriter(path string) *t_nobuf_log_writer {
	return &t_nobuf_log_writer{
		f: openLogFile(path),
	}
}

func (writer *t_nobuf_log_writer) writeln(s string) (int, error) {
	return writer.f.WriteString(s + "\n")
}

func (writer *t_nobuf_log_writer) flush() {
}

func (writer *t_nobuf_log_writer) free() {
	writer.f.Close()
}

/**  @} */
