# go-merge

Clone the repository.

```sh
git clone https://github.com/chrisbradleydev/go-merge.git .
```

Build Docker image.

```sh
docker build -t go-merge .
```

Run Docker container.

```sh
docker run -v $(pwd):/app -v /app/tmp go-merge
```
