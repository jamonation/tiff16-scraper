require 'nokogiri'
require 'open-uri'
require 'json'

urls_file = File.read("urls.json")
urls = JSON.parse(urls_file)
films = Array.new

for url in urls do
  tiffDOM = Nokogiri::HTML(open(url), nil, 'utf-8')
  film = Hash.new
  film["name"] = tiffDOM.css("body.film #wrap .container h1").text || ""
  film["director"] = tiffDOM.css("#director .credit-content").text || ""
  film["countries"] = tiffDOM.css("span.quick-info .countries").text || ""
  film["runtime"] = tiffDOM.css("span.quick-info .runtime").text || ""
  film["premiere"] = tiffDOM.css("span.quick-info .premiere").text || ""
  film["year"] = tiffDOM.css("span.quick-info .year").text || ""
  film["language"] = tiffDOM.css("span.quick-info .language").text || ""
  film["pitch"] = tiffDOM.css(".pitch").text || ""
  film["production"] = tiffDOM.css("#productionCompany .credit-content").text || ""
  film["producers"] = tiffDOM.css("#producers .credit-content").text || ""
  film["screenplay"] = tiffDOM.css("#screenplay .credit-content").text || ""
  film["cinematographers"] = tiffDOM.css("#cinematographers .credit-content").text || ""
  film["editors"] = tiffDOM.css("#editors .credit-content").text || ""
  film["score"] = tiffDOM.css("#originalScore .credit-content").text || ""
  film["sound"] = tiffDOM.css("#sound .credit-content").text || ""
  film["cast"] = tiffDOM.css("#cast .credit-content").text || ""
  films.push(film)
end

File.open("films/films.json","w") do |f|
  f.write(JSON.pretty_generate(films))
end
