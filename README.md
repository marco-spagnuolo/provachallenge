# BUILD (ITA) 

nell'esempio si usa:
  golang ver 16.3 
  macOS ver 10.14

## Se si ha golang già installato :

go run main.go 

## Se non si ha golang installato :

./provachallenge

# unit test 

go test 

![Schermata 2023-02-10 alle 04 02 36](https://user-images.githubusercontent.com/81595718/217993317-bbaa1a5b-3be2-407e-b052-f9a24f92555e.png)


# Deployment 

## build image from Dockerfile 

docker build --rm -t challengedocker:alpha .

## run container 

docker run challengedocker:alpha 



# BUILD (ENG VERSION)


in the example we use:
    Golang version 16.3
    macOS version 10.14

## If you have golang already installed

go run main.go

## If you don't have golang installed (assumed we are using linux)

./trychallenge

# unit test

go to the test

![Schermata 2023-02-10 alle 04 02 36](https://user-images.githubusercontent.com/81595718/217993317-bbaa1a5b-3be2-407e-b052-f9a24f92555e.png)


# Distribution

## create an image from Dockerfile

docker build --rm -t challengedocker:alpha .

## run container

docker run challengedocker:alpha
