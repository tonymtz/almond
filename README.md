# Almond

[ ![Codeship Status for tonymtz/almond](https://app.codeship.com/projects/4f4db340-2342-0135-1185-229ae24d69ea/status?branch=master)](https://app.codeship.com/projects/221937)
[![codecov](https://codecov.io/gh/tonymtz/almond/branch/master/graph/badge.svg?token=sf0qEs1thZ)](https://codecov.io/gh/tonymtz/almond)

## Installation

### Requirements

- Golang v1.8
- Docker v17.03

#### Postgresql

1. Let docker do the magic: `$ docker-compose up` 🐳
2. Ready to rock on port `5433`

Note: This is not following default port for postgres, however the migrations script is aware of this. Worry not.

#### Golang

golang packages required for development:

- [govendor](https://github.com/kardianos/govendor): `$ go get -u github.com/kardianos/govendor`
- [goose](https://bitbucket.org/liamstask/goose/): `$ go get bitbucket.org/liamstask/goose/cmd/goose`

Install vendors:

```bash
$ govendor sync
```

Run migrations (requires postgres service, check docker composer):

```bash
$ goose up
```

### Compiling

```bash
$ go run main.go
```

Open your browser and navigate to [http://localhost:3000](http://localhost:3000)

### Testing

```bash
$ govendor test +local -cover
```

Ready to hack!
