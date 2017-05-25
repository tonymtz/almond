# almond 

## Installation

### Requirements

golang packages required for development:

- [govendor](https://github.com/kardianos/govendor): `$ go get -u github.com/kardianos/govendor`
- [goose](https://bitbucket.org/liamstask/goose/): `$ go get bitbucket.org/liamstask/goose/cmd/goose`

Install vendors:

```bash
$ govendor sync
```

Run migrations:

```bash
$ goose up
```


