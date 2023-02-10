# BUILD

nell'esempio si usa:
  golang ver 16.3 
  macOS ver 10.14

## Se si ha golang già installato :

go run main.go 

## Se non si ha golang installato :

./provachallenge

# unit test 

go test 

# Deployment 

## build image from Dockerfile 

docker build --rm -t challengedocker:alpha .

## run container 

docker run challengedocker:alpha 



ENG VERSION
# BUILD

in the example we use:
    Golang version 16.3
    macOS version 10.14

## If you have golang already installed

go run main.go

## If you don't have golang installed (assumed we are using linux)

./trychallenge

# unit test

go to the test

# Distribution

## create an image from Dockerfile

docker build --rm -t challengedocker:alpha .

## run container

docker run challengedocker:alpha
