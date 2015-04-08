# Docker Registry Manager
This is a simple Go application for working with the Docker Registry.  Currently,
this only supports v1 of the registry.

# Usage

## Show Help

```
registry-manager -h
```

## Search

```
$> registry-manager search
Number of Repositories: 6
 -  Name: library/busybox
    Tags: 2
    Layers: 10
 -  Name: demo/busybox
    Tags: 1
    Layers: 5
 -  Name: library/swarm
    Tags: 1
    Layers: 9
 -  Name: library/redis
    Tags: 1
    Layers: 19
 -  Name: ehazlett/docker-demo
    Tags: 1
    Layers: 7
 -  Name: test/busybox
    Tags: 1
    Layers: 5
```

This will return all repositories in the registry.

```
$> registry-manager search -q busybox
Number of Repositories: 3
 -  Name: library/busybox
    Tags: 2
    Layers: 10
 -  Name: demo/busybox
    Tags: 1
    Layers: 5
 -  Name: test/busybox
    Tags: 1
    Layers: 5
```

This will search the registry for repositories containing the term `busybox`.

## Inspect Repo

```
$> registry-manager inspect-repo -n busybox
Namespace: library
Repository: busybox
Layers: 
 - ID: 4986bf8c15363d1c5d15512d5266f8777bfba4974ac56e3270e7760f6f0a8125
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) CMD [/bin/sh]
 - ID: 4986bf8c15363d1c5d15512d5266f8777bfba4974ac56e3270e7760f6f0a8125
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) CMD [/bin/sh]
 - ID: ea13149945cb6b1e746bf28032f02e9b5a793523481a0a18645fc77ad53c4ea2
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 2433303
   History: 
    - /bin/sh
    - -c
    - #(nop) ADD file:8cf517d90fe79547c474641cc1e6425850e04abbd8856718f7e4a184ea878538 in /
 - ID: df7546f9f060a2268024c8a230d8639878585defcc1bc6f79d2728a13957871b
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.2.0
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) MAINTAINER Jérôme Petazzoni <jerome@docker.com>
 - ID: 511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158
   Author: 
   Docker Version: 0.4.0
   Size: 0
   History: 
 - ID: 4986bf8c15363d1c5d15512d5266f8777bfba4974ac56e3270e7760f6f0a8125
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) CMD [/bin/sh]
 - ID: 4986bf8c15363d1c5d15512d5266f8777bfba4974ac56e3270e7760f6f0a8125
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) CMD [/bin/sh]
 - ID: ea13149945cb6b1e746bf28032f02e9b5a793523481a0a18645fc77ad53c4ea2
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.4.1
   Size: 2433303
   History: 
    - /bin/sh
    - -c
    - #(nop) ADD file:8cf517d90fe79547c474641cc1e6425850e04abbd8856718f7e4a184ea878538 in /
 - ID: df7546f9f060a2268024c8a230d8639878585defcc1bc6f79d2728a13957871b
   Author: Jérôme Petazzoni <jerome@docker.com>
   Docker Version: 1.2.0
   Size: 0
   History: 
    - /bin/sh
    - -c
    - #(nop) MAINTAINER Jérôme Petazzoni <jerome@docker.com>
 - ID: 511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158
   Author: 
   Docker Version: 0.4.0
   Size: 0
   History: 
```

## Delete Tag

```
$> registry-manager delete-tag -n busybox -t latest
```

## Delete Repo

```
$> registry-manager delete-repo -n foo/test
```
