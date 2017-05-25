# Almond

[ ![Codeship Status for tonymtz/almond](https://app.codeship.com/projects/4f4db340-2342-0135-1185-229ae24d69ea/status?branch=master)](https://app.codeship.com/projects/221937)

## Installation

### Requirements

#### Postgresql

1. Install docker-compose
2. let docker do the magic: `$ docker-compose up`
3. Ready to rock on port 5433

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

Ready to hack!
