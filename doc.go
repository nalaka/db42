// Package db42 is an experimental, in-process SQL database engine inspired by SQLite.
// Although db42 uses the SQLite file format, it is written entirely from scratch in Go
// with minimal external dependencies. It aims to support features like row-level locking
// and the subset of SQL essential for monolithic web applications like praja-dev/porgs.
package db42
