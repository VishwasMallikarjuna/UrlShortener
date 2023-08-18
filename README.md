# UrlShortner
A simple URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result.

## URL shortener example 
https://bitly.com/
## API with decent architecture.
 
 # Endpoints:
    1.Http method Post: http://localhost:8192/short
        Send the url within the request body 
        example:
            {
	            "url":"https://github.com/VishwasMallikarjuna/UrlShortener"
            }
        Response wil be shortend URL
    2.Http method Get: http://localhost:8192/original/{shortendpayload}
        example:
               http://localhost:8192/original/sF10y
            Response wil ORIGINAL url
    3.Http method Get: http://localhost:8192/top3Dns
        Response that returns top 3 domain names that have been shortened the most number of times

## You can pull the docker image by using commad
    docker pull vishwasmallikarjuna/urlrepo:latest

## you can run the image by using command
    $ docker run --rm -p 8192:8192 vishwasmallikarjuna/urlrepo:latest