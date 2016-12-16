# Datastore Load/Save test

Requires govendor, install `go get -u github.com/kardianos/govendor`

## Serve
> $ goapp serve app    
> go to http://localhost:8080/do

## Reproduce error
> update to latest app engine pkg version (currently 08a149cfaee099e6ce4be01c0113a78c85ee1dee)   
> $ govendor fetch google.golang.org/appengine/...    
> go to http://localhost:8080/do
