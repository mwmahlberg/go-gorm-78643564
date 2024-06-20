go-gorm-78643564
================

This is code related to [my answer][answer] to the question [dial tcp 192.168.48.2:3306: connect: connection refused][question].

Prerequisites
-------------

  * a running containerization
  * `docker compose` or an according wrapper

Run
---

```shell
$ git clone git@github.com:mwmahlberg/go-gorm-78643564.git 
Cloning into 'go-gorm-78643564'...
remote: Enumerating objects: 21, done.
remote: Counting objects: 100% (21/21), done.
remote: Compressing objects: 100% (19/19), done.
remote: Total 21 (delta 4), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (21/21), 4.86 KiB | 2.43 MiB/s, done.
Resolving deltas: 100% (4/4), done.
$ cd go-gorm-78643564
$ docker compose up -d
```

[question]: https://stackoverflow.com/questions/78643564
[answer]: https://stackoverflow.com/a/78645702/1296707