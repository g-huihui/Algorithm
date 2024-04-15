/**
 * @Author: Gong Yanhui
 * @Description: 继承net.Conn 只读连接
 * @Date: 2024-04-15 15:27
 */

package sni

import (
	"io"
	"net"
	"time"
)

type readOnlyConn struct {
	reader io.Reader
}

func (conn readOnlyConn) Read(p []byte) (n int, err error) {
	return conn.reader.Read(p)
}

func (conn readOnlyConn) Write(p []byte) (n int, err error) {
	return 0, io.ErrClosedPipe
}

func (conn readOnlyConn) Close() error {
	return nil
}

func (conn readOnlyConn) LocalAddr() net.Addr {
	return nil
}

func (conn readOnlyConn) RemoteAddr() net.Addr {
	return nil
}

func (conn readOnlyConn) SetDeadline(t time.Time) error {
	return nil
}

func (conn readOnlyConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (conn readOnlyConn) SetWriteDeadline(t time.Time) error {
	return nil
}
