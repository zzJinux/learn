
BuildKit is the default builder (Docker Desktop and Docker Engine as of version 23.0).
To enable it explicitly, set `DOCKER_BUILDKIT=1`.

`$ docker builder prune -af` - clear your build cache

`$ docker build --target=client --progress=plain . 2> log1.txt` - print process and output a log file

https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/
https://www.docker.com/blog/containerize-your-go-developer-environment-part-2/
https://www.docker.com/blog/containerize-your-go-developer-environment-part-3/

https://docs.docker.com/language/golang/build-images/
https://docs.docker.com/build/guide/mounts/
