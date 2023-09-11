# Bind mounts (and linter example)

`COPY` creates an extra layer in the image, which takes up time and space. You even can get rid of such layers by using
`RUN --mount=type=bind`, which can bind mounting from the build context, stage or an image.

`bind` is the default mount type, so you can omit it. The following two lines are equivalent:

```Dockerfile
RUN --mount=type=bind,source=. .
RUN --mount=source=.,target=. .
```

