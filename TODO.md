Open and close a database file

- Open the file with an exclusive lock so no other process can open it
- Open(string) (*db42.DB, error)
- Close(*db42.DB) error

Check locking behavior on Mac, Windows, and Linux

