# Fox Db

Fox db is a small and fast in-memory key value database capable of concurrent reads and writes. It also persists the database by writing the gob of database on disk.

### Setting up
```
# Clone the repository
$ git clone git@github.com:OmairK/fox.git

$ cd fox
```

```
# Test if the server is up & running
$ nc -v 0.0.0.0 8000
> CHECK
YIP YIP
```
#### Docker
* Pre-requisites
	* Docker
```
# Build the docker image
$ docker build -t fox_db .

# Run the container with docker image
$ docker run --name=fox_db -p 8000:8000 fox 
```

#### Source
* Pre-requisites
	* [Go](https://golang.org/dl/)
```
$ go build fox.go
$ ./fox.go
```
```
# Test if the server is up & running
$ nc -v 0.0.0.0 8000
> CHECK
YIP YIP
```