```markdown
# URL Shortener

A simple URL shortening service implemented in Go.

## Requirements

- Go (version 1.18 or higher)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/url-shortener.git
   ```

2. Navigate to the project directory:

   ```sh
   cd url-shortener
   ```

3. Build the project:

   ```sh
   go build
   ```

4. Run the project:

   ```sh
   ./url-shortener
   ```

## Usage

### Shorten a URL:

Send a POST request to `http://localhost:3000/short-it` with the `url` parameter set to the URL you want to shorten.

### Redirect to a shortened URL:

Send a GET request to `http://localhost:3000/short/{key}` where `{key}` is the shortened URL key.

## Example

### Shorten a URL:

```sh
curl -X POST -d "url=https://example.com" http://localhost:3000/short-it
```

### Redirect to a shortened URL:

```sh
curl -i http://localhost:3000/short/abc123
```

## License

This project is licensed under the MIT License.

## Contributing

If you'd like to contribute to the project, please fork the repository and submit a pull request.

## Credits

- [go-chi/chi](https://github.com/go-chi/chi)
- [lithammer/shortuuid](https://github.com/lithammer/shortuuid)
```

This format includes headers, code blocks, and links, making it easier to read and follow the instructions.Here is the formatted README.md file for your URL Shortener project:

```markdown
# URL Shortener

A simple URL shortening service implemented in Go.

## Requirements

- Go (version 1.18 or higher)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/url-shortener.git
   ```

2. Navigate to the project directory:

   ```sh
   cd url-shortener
   ```

3. Build the project:

   ```sh
   go build
   ```

4. Run the project:

   ```sh
   ./url-shortener
   ```

## Usage

### Shorten a URL:

Send a POST request to `http://localhost:3000/short-it` with the `url` parameter set to the URL you want to shorten.

### Redirect to a shortened URL:

Send a GET request to `http://localhost:3000/short/{key}` where `{key}` is the shortened URL key.

## Example

### Shorten a URL:

```sh
curl -X POST -d "url=https://example.com" http://localhost:3000/short-it
```

### Redirect to a shortened URL:

```sh
curl -i http://localhost:3000/short/abc123
```

## License

This project is licensed under the MIT License.

## Contributing

If you'd like to contribute to the project, please fork the repository and submit a pull request.

## Credits

- [go-chi/chi](https://github.com/go-chi/chi)
- [lithammer/shortuuid](https://github.com/lithammer/shortuuid)
```
