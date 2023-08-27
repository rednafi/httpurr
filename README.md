<div align="left">
    <h1>ᗢ httpurr</h1>
    <i>HTTP status codes on speed dial</i>
    <div align="right">
</div>

---

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
    ==========

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
    ==========

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
    httpurr -code 410
    ```

    ```txt
    ᗢ httpurr
    ==========

    Description
    -----------

    The HyperText Transfer Protocol (HTTP) 410 Gone client error response code
    indicates that access to the target resource is no longer available at the
    origin server and that this condition is likely to be permanent.

    If you don't know whether this condition is temporary or permanent, a 404 status
    code should be used instead.

    Status
    ------

    410 Gone

    Source
    ------

    https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410
    ```
