# UrlShortner
A simple URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result.

## URL shortener example 
https://bitly.com/
## API with decent architecture.
 
 # Endpoints:
    1. Http method Post: http://localhost:8192/short
        Send the URL within the request body 
        example:
            {
	            "url":"https://github.com/VishwasMallikarjuna/UrlShortener"
            }
        The response will be shortened URL
    2. Http method Get: http://localhost:8192/original/{shortendpayload}
        example:
               http://localhost:8192/original/sF10y
            The response will ORIGINAL URL
    3. Http method Get: http://localhost:8192/top3Dns
        The response that returns the top 3 domain names that have been shortened the most number of times

## To run locally go installed
	go run main.go

## You can pull the docker image by using the command
    docker pull vishwasmallikarjuna/urlrepo:latest

## You can run the image by using the command
    docker run --rm -p 8192:8192 vishwasmallikarjuna/urlrepo:latest
