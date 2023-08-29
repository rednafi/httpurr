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

* Filter the status codes by categories:

    ```sh
    httpurr -list -cat 2
    ```

    ```txt
    ᗢ httpurr
    ==========

    Status Codes
    ------------

    ------------------ 2xx ------------------

    200    OK
    201    Created
    202    Accepted
    203    Non-Authoritative Information
    204    No Content
    205    Reset Content
    206    Partial Content
    207    Multi-Status
    208    Already Reported
    226    IM Used
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
* To publish a new version, create a new [release] with a [tag], and the [CI] will take care
of the rest.

[cover-img]: https://github.com/rednafi/httpurr/assets/30027932/1c8e01fc-e943-4adf-b212-56584ff99f5d
[binary]: https://github.com/rednafi/httpurr/releases/latest
[tag]: https://github.com/rednafi/httpurr/tags
[release]: https://github.com/rednafi/httpurr/releases/new
[CI]: ./.github/workflows/release.yml
