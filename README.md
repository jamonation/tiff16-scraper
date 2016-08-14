# tiff16-scraper
This is a Golang & goquery fork of James Wilkinson's https://github.com/jwilkinson/tiff16-scraper application that scrapes the TIFF.net website for festival 2016 film data and renders it as JSON.

## Copyright Notice
As per TIFF Terms and Conditions, this data is being used soley for educational, non-commercial purposes. Read the full terms and conditions of the TIFF site and data before modifying for your own use http://tiff.net/terms/

## How it Works
* This application runs on your local machine, it reads in a json file of URLs that represent the urls of the TIFF16 films
* That urls.json file is generated through javascript in the browser on TIFF.net to parse out all the URLS on the "at the festival" filter of films
* Then the application hits each of those URLs, during which goquery uses CSS selectors to grab the element the data is in, and then using the .text() api of goquery returns plain text
* This plain text is then parsed into a piece of the films hash keys
* All the keys are then put into the films slice
* The films slice then is written back to a JSON file

## System Specifications
This process takes about 5-7 seconds on MacBook Air (13-inch, Early 2015) 1.6 GHz Intel Core i5 and at 5Mbps internet connection.

### To Do
* Generate the URLs.json file from tiff.net list of festival films using goquery
