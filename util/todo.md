# Go Lang Util To Do
- - -

## Watcher

Create a file watcher to watch a set of directories and execute a command when a file changes. After searching for watchers, it looks like most implementations use the OS notification, which is broken much of the time.  So, the implementation should probably do a brute force scan once every second to determine changes based on modification date and file size (or hash).


