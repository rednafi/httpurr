<div align="left">
    <h1>ᗢ httpurr</h1>
    <i>HTTP status codes on speed dial</i>
    <div align="right">
</div>

---

```txt
ᗢ httpurr
=========

Description
-----------

The HyperText Transfer Protocol (HTTP) 505 HTTP Version Not Supported response
status code indicates that the HTTP version used in the request is not supported
by the server.

Status
------

505 HTTP Version Not Supported

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505
```


## Installation

* On MacOS, brew install:

    ```sh
    brew tap rednafi/httpurr https://github.com/rednafi/httpurr \
        && brew install httpurr
    ```

* Or else, go install:

    ```sh
    go install github.com/rednafi/httpurr/cmd/httpurr
    ```

## Usage

* Print the options:

    ```sh
    httpurr -help
    ```

    ```
    ᗢ httpurr
    =========

    Usage of httpurr:
    -code string
            Print the description of an HTTP status code
    -help
            Print usage
    -list
            Print HTTP status codes
    -version
            Print version
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
    httpurr -code 418
    ```

    ```txt
    ᗢ httpurr
    =========

    Description
    -----------

    The HTTP 418 I'm a teapot client error response code indicates that the server
    refuses to brew coffee because it is, permanently, a teapot. A combined
    coffee/tea pot that is temporarily out of coffee should instead return 503. This
    error is a reference to Hyper Text Coffee Pot Control Protocol defined in April
    Fools' jokes in 1998 and 2014.

    Some websites use this response for requests they do not wish to handle, such as
    automated queries.

    Status
    ------

    418 I'm a teapot

    Source
    ------

    https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418
    ```
