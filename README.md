<div align="left">
    <h1>ᗢ httpurr</h1>
    <i>HTTP status codes on speed dial</i>
    <div align="right">
</div>

---

![img][cover-img]

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

[cover-img]: https://user-images.githubusercontent.com/30027932/263548019-7de08764-5030-4d65-95d5-166d226bc7d9.png
