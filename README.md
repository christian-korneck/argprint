# argprint

a simple troubleshooting tool that repeats the command line arguments that get passed to it. Useful to inspect escaping, etc. Can be used i.e. as a temporary wrapper for any other command:

```
$ alias docker=argprint
$ docker run --rm -ti fedora:34 /bin/bash -c "echo hello world"
-----------------
args: 8
argprint
run
--rm
-ti
fedora:34
/bin/bash
-c
echo hello world
-----------------
$ unalias docker
```