# Container fundamentals

This is based on the talk and [code](https://github.com/lizrice/containers-from-scratch/blob/master/main.go)
from Liz Rice.

## Prepare a chroot file system

These commands extract the image for a Debian container, so we can later use it
to transform our process on an Ubuntu machine to a Debian machine.

```sh
docker run -it debian sh

# install ps and other tools
apt update && apt install procps

docker export container_name > output.tar
ls 
mkdir debianfs && cd debianfs 
mv debian.tar debianfs
tar xvf debian.tar
```

## What to show?

1. Change hostname in container while having `CLONE_NEWUTS` set.
   It will **only** change the hostname inside the container.

   ```sh
   go run main.go /bin/bash
   hostname
   # set hostname to "container"
   hostname container

   # should now ouput "container
   hostname

   # also run hostname on the host to see it hasn't changed 
   ```


2. Add `CLONE_NEWPID` to give the newly spanned child a new process ID starting from `0`.


3. Show `ls /proc`.
   The process shares the same `/proc` file system with the host.
   Set up the chroot mount with enabling (follow the upper documentation on how to prepare it first):
   
   ```go
   	syscall.Chroot("/home/aps/share/gocker/debianfs")
	   syscall.Chdir("/")
   ```

   Start again with `go run main.go /bin/bash` and execute `cat /etc/os-release`.
   We are now Debian instead of Ubuntu.


   Set container to sleep with `sleep 1000`.
   Then find the sleep process from the host with `ps -C sleep`.
   Show what we know about the process with `ls /proc/ID`.
   How does the container know it's root? -> `ls -l /proc/ID/root`


   Does `ps` work now?
   No it is a pseudo file system and is coming from the kernel.
   Add the following:

   ```go
   syscall.Mount("proc", "proc", "proc", 0, "")
   ```

   Now `ps` should only show the processes started in the container. 