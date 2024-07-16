URL Shortener
A simple URL shortening service implemented in Go.

Requirements
Go (version 1.18 or higher)
Installation
Clone the repository:
Copy
Insert in Terminal
git clone https://github.com/your-username/url-shortener.git
Navigate to the project directory:
Copy
Insert in Terminal
cd url-shortener
Build the project:
Copy
Insert in Terminal
go build
Run the project:
Copy
Insert in Terminal
./url-shortener
Usage
Shorten a URL:
Send a POST request to http://localhost:3000/short-it with the url parameter set to the URL you want to shorten.
Redirect to a shortened URL:
Send a GET request to http://localhost:3000/short/{key} where {key} is the shortened URL key.
Example
Shorten a URL:
Copy
Insert in Terminal
curl -X POST -d "url=https://example.com" http://localhost:3000/short-it
Redirect to a shortened URL:
Copy
Insert in Terminal
curl -i http://localhost:3000/short/abc123
License
This project is licensed under the MIT License.

Contributing
If you'd like to contribute to the project, please fork the repository and submit a pull request.

Credits
go-chi/chi
lithammer/shortuuidi