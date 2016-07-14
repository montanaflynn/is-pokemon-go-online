# Is Pokémon Go Online?

### CLI Usage

If you have `go` installed and `$GOPATH/bin` in your path:

```
go get github.com/montanaflynn/is-pokemon-go-online/examples/pogo-cli
pogo-cli
```

Otherwise use the [binary release](https://github.com/montanaflynn/is-pokemon-go-online/releases/tag/0.1.0), for example on a mac:

```
wget https://github.com/montanaflynn/is-pokemon-go-online/releases/download/0.1.0/pogo-cli-osx -O pogo-cli
chmod +x pogo-cli
./pogo-cli
```

### API Usage

Pretty much the same as the cli except you can set the port:

```
# From source
go get github.com/montanaflynn/is-pokemon-go-online/examples/pogo-cli
pogo-cli :4444 &
```

```
# From release
wget https://github.com/montanaflynn/is-pokemon-go-online/releases/download/0.1.0/pogo-api-osx -O pogo-api
chmod +x pogo-api
./pogo-api :4444
```

There's also a [Dockerfile](https://github.com/montanaflynn/is-pokemon-go-online/blob/master/examples/pogo-api/Dockerfile), if you're into that kind of thing:

```
git clone github.com/montanaflynn/is-pokemon-go-online
cd ./is-pokemon-go-online
docker build -t $(whoami)/pogo ./examples/pogo-api/
docker run -d -p 4444:8080 $(whoami)/pogo
```

In the above examples the web service will be running on [localhost:4444](http://localhost:4444)!
