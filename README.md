# tiff16-scraper
This is a Ruby and Nokogiri application that scrapes the TIFF.net website for festival 2016 film data. 

## Copyright Notice
As per TIFF Terms and Conditions, this data is being used soley for educational, non-commercial purposes. Read the full terms and conditions of the TIFF site and data before modifying for your own use http://tiff.net/terms/

## How it Works
* This application runs on your local machine, it reads in a json file of URLs that represent the urls of the TIFF16 films. * That urls.json file is generate through javascript in the browser on TIFF.net to parse out all the URLS on the "at the festival" filter of films.
* Then the application hits each of those URLs, during witch Nokogiri uses CSS selectors to grab the element the data is in, and then using the .text() api of Nokogiri returns plain text.
* This plain text is then parsed into a piece of the films hash keys.
* All the keys are then put into the films array.
* The films array then is wrote back to a JSON file, using the JSON and File gems.

## System Specifications
This process takes about 2 minutes on MacBook Air (13-inch, Early 2015) 1.6 GHz Intel Core i5 and at 250Mbps internet connection. This process will take much longer on a slower internet connection and slower processor.

### To Do
* Make the URLs.json file with Nokogiri, requires site interaction to load all the data before parsing the dom for the urls of films.
