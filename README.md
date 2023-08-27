<div align="left">
    <h1>ᗢ httpurr</h1>
    <strong><i> >> HTTP status codes on speed dial << </i></strong>
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

* Or elsewhere, go install:

    ```sh
    go install github.com/rednafi/httpurr/cmd/httpurr
    ```

* Else, download the appropriate [binary] for your CPU arch and add it to the `$PATH`.

## Quickstart

* Print the options:

    ```sh
    httpurr -help
    ```

    ```txt
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

## Development

* Clone the repo.
* Go to the root directory and run:
    ```sh
    make init
    ```
* Run the linter:
    ```sh
    make lint
    ```
* Run the tests:
    ```sh
    make test
    ```

[cover-img]: https://github.com/rednafi/httpurr/assets/30027932/1c8e01fc-e943-4adf-b212-56584ff99f5d
[binary]: https://github.com/rednafi/httpurr/releases/latest
