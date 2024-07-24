# db42

> Package `db42` is an experimental, in-process SQL database engine inspired by SQLite.
Although `db42` uses the SQLite file format, it is written entirely from scratch in Go
with minimal external dependencies. It aims to support features like row-level locking
and the subset of SQL essential for monolithic web applications like
[praja-dev/porgs](https://github.com/praja-dev/porgs). 

## Initial design decisions

**Positioning:**
- Does not try to emulate `SQLite`. Instead, aims to provide a simple, reliable, and
  performant database engine tailored for monolithic web applications written in Go lang.
- Support only a subset of `SQL` statements, defined by the requirement for `porgs`.
- Excellent support for row-level locking.


**Database file format:**
- Use the [SQLite file format](https://www.sqlite.org/fileformat.html). Database files created
  by `db42` should open in SQLite but the reverse is not required.

**Behavior:**
- When a database file is opened, acquire an exclusive non-blocking lock, allowing other
  processes to read but not write to the file. Application using db42 is expected to
  open the database file at startup and maintain the lock until shutdown.
- Use only WAL (Write-Ahead Logging) mode.

**Hygiene:**
- Depend only on the Go standard library and `golang.org/x/*` packages.


## Release plan

Version 1 of `db42` is when it can replace SQLite in [praja-dev/porgs](https://github.com/praja-dev/porgs).
This is at least a year away.
