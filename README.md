# ᗢ httpurr [under construction]

HTTP status codes on speed dial

## Installation

* Brew install:

    ```sh
    brew tap rednafi/httpurr https://github.com/rednafi/httpurr \
        && brew install httpurr
    ```

* Go install:

    ```sh
    go install  github.com/rednafi/httpurr/cmd/httpurr
    ```

## Usage

* Print the options:

    ```sh
    httpurr -help
    ```

    ```txt
    ᗢ httpurr
    =========

    Usage of httpurr:
    -code string
            Print description of an HTTP status code
    -help
            Print usage
    -list
            List all HTTP status codes
    ```

* List the HTTP status codes:

    ```sh
    httpurr -list
    ```

    ```txt
    ᗢ httpurr
    =========

    Status Codes
    ------------

    ------------------ 1xx ------------------

    100    Continue
    101    Switching Protocols
    102    Processing
    103    Early Hints

    ------------------ 2xx ------------------
    ...
    ```

* Display the description of a status code:

    ```sh
    httpurr -code 429
    ```

    ```txt
    ᗢ httpurr
    =========

    Description
    -----------

    The HTTP 429 Too Many Requests response status code indicates the user has sent
    too many requests in a given amount of time ("rate limiting").

    A Retry-After header might be included to this response indicating how long to
    wait before making a new request.

    Status
    ------

    429 Too Many Requests

    Source
    ------

    https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429
    ```
