package src

const (
	statusContinue100 = `
Description
-----------

The HTTP 100 Continue informational status response code indicates that
everything so far is OK and that the client should continue with the request or
ignore it if it is already finished.

To have a server check the request's headers, a client must send Expect:
100-continue as a header in its initial request and receive a 100 Continue
status code in response before sending the body.

Status
------

100 Continue

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100
`

	statusSwitchingProtocols101 = `
Description
-----------

The HTTP 101 Switching Protocols response code indicates a protocol to which the
server switches. The protocol is specified in the Upgrade request header
received from a client.

The server includes in this response an Upgrade response header to indicate the
protocol it switched to.

Status
------

101 Switching Protocols

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101
`

	statusProcessing102 = `
Description
-----------

Deprecated: This status code is deprecated. When used, clients may still accept
it, but simply ignore them.

The HTTP 102 Processing informational status response code indicates to client
that a full request has been received and the server is working on it.

This status code is only sent if the server expects the request to take
significant time. It tells the client that your request is not dead yet.

Status
------
102 Processing

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/102
`

	statusEarlyHints103 = `
Description
-----------

Experimental: This is an experimental technology. Check the Browser
compatibility table carefully before using this in production.

The HTTP 103 Early Hints information response may be sent by a server while it
is still preparing a response, with hints about the resources that the server is
expecting the final response will link.

This allows a browser to start preloading resources even before the server has
prepared and sent that final response.

The early hint response is primarily intended for use with the Link header,
which indicates the resources to be loaded.

It may also contain a Content-Security-Policy header that is enforced while
processing the early hint.

A server might send multiple 103 responses, for example, following a redirect.
Browsers only process the first early hint response, and this response must be
discarded if the request results in a cross-origin redirect. Preloaded resources
from the early hint are effectively pre-pended to the Document's head element,
and then followed by the resources loaded in the final response.

Status
------

103 Early Hints

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103
`
	statusOK200 = `
Description
-----------

The HTTP 200 OK success status response code indicates that the request has
succeeded. A 200 response is cacheable by default.

The meaning of a success depends on the HTTP request method:

* GET   : The resource has been fetched and is transmitted in the message body.
* HEAD  : The representation headers are included in the response without any
          message body
* POST  : The resource describing the result of the action is transmitted in the
		  message body
* TRACE : The message body contains the request message as received by the
          server. The successful result of a PUT or a DELETE is often not a 200
	  OK but a 204 No Content (or a 201 Created when the resource is
	  uploaded for the first time).

Status
------

200 OK

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
`
	statusCreated201 = `
Description
-----------

The HTTP 201 Created success status response code indicates that the request has
succeeded and has led to the creation of a resource.

The new resource, or a description and link to the new resource, is effectively
created before the response is sent back and the newly created items are
returned in the body of the message, located at either the URL of the request,
or at the URL in the value of the Location header.

The common use case of this status code is as the result of a POST request.

Status
------

201 Created

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201
`
	statusAccepted202 = `
Description
-----------

The HTTP 202 Accepted response status code indicates that the request has been
accepted for processing, but the processing has not been completed; in fact,
processing may not have started yet.

The request might or might not eventually be acted upon, as it might be
disallowed when processing actually takes place.

202 is non-committal, meaning that there is no way for the HTTP to later send an
asynchronous response indicating the outcome of processing the request. It is
intended for cases where another process or server handles the request, or for
batch processing.

Status
------

202 Accepted

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/202
`
	statusNonAuthoritativeInfo203 = `
Description
-----------

The HTTP 203 Non-Authoritative Information response status indicates that the
request was successful but the enclosed payload has been modified by a
transforming proxy from that of the origin server's 200 OK response.

The 203 response is similar to the value 214, meaning Transformation Applied, of
the Warning header code, which has the additional advantage of being applicable
to responses with any status code.

Status
------

203 Non-Authoritative Information

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/203
`
	statusNoContent204 = `
Description
-----------

The HTTP 204 No Content success status response code indicates that a request
has succeeded, but that the client doesn't need to navigate away from its
current page.

This might be used, for example, when implementing "save and continue editing"
functionality for a wiki site. In this case a PUT request would be used to save
the page, and the 204 No Content response would be sent to indicate that the
editor should not be replaced by some other page.

A 204 response is cacheable by default (an ETag header is included in such a
response).

Status
------

204 No Content

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204
`
	statusResetContent205 = `
Description
-----------

The HTTP 205 Reset Content response status tells the client to reset the
document view, so for example to clear the content of a form, reset a canvas
state, or to refresh the UI.

Status
------

205 Reset Content

Source
------
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/205
`
	statusPartialContent206 = `
Description
-----------

The HTTP 206 Partial Content success status response code indicates that the
request has succeeded and the body contains the requested ranges of data, as
described in the Range header of the request.

If there is only one range, the Content-Type of the whole response is set to the
type of the document, and a Content-Range is provided.

If several ranges are sent back, the Content-Type is set to multipart/
byteranges and each fragment covers one range, with Content-Range and
Content-Type describing it.

Status
------

206 Partial Content

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206
`
	statusMultiStatus207 = `
Description
-----------

Note: The ability to return a collection of resources is part of the WebDAV
protocol (it may be received by web applications accessing a WebDAV server).
Browsers accessing web pages will never encounter this status code.

The HTTP 207 Multi-Status response code indicates that there might be a mixture
of responses.

The response body is a text/xml or application/xml HTTP entity with a
multistatus root element. The XML body will list all individual response codes.

Status
------

207 Multi-Status

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/207
`
	statusAlreadyReported208 = `
Description
-----------

The HTTP 208 Already Reported response code is used in a 207 (207 Multi-Status)
response to save space and avoid conflicts. If the same resource is requested
several times (for example as part of a collection), with different paths, only
the first one is reported with 200. Responses for all other bindings will report
with this 208 status code, so no conflicts are created and the response stays
shorter.

Status
------

208 Already Reported

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/208
`
	statusIMUsed226 = `
Description
-----------

Note: Browsers don't support delta encoding with HTTP. This status code is sent
back by custom servers used by specific clients.

In the context of delta encodings, the HTTP 226 IM Used status code is set by
the server to indicate that it is returning a delta to the GET request that it
received.

With delta encoding a server responds to GET requests with differences (called
deltas) relative to a given base document (rather than the current document).
The client uses the A-IM: HTTP header to indicate which differencing algorithm
to use and the If-None-Match: header to hint the server about the last version
it got. The server generates a delta, sending it back in an HTTP response with
the 226 status code and containing the IM: (with the name of the algorithm
used) and Delta-Base: (with the ETag matching the base document associated to
the delta) HTTP headers.

IM stands for instance manipulations the term used to describe an algorithm
generating a delta.

Status
------

226 IM Used

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/226
`
	statusMultipleChoices300 = `
Description
-----------

The HTTP 300 Multiple Choices redirect status response code indicates that the
request has more than one possible response. The user-agent or the user should
choose one of them. As there is no standardized way of choosing one of the
responses, this response code is very rarely used.

If the server has a preferred choice, it should generate a Location header.

Status
------

300 Multiple Choices

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/300
`
	statusMovedPermanently301 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 301 Moved Permanently redirect status
response code indicates that the requested resource has been definitively moved
to the URL given by the Location headers. A browser redirects to the new URL
and search engines update their links to the resource.

Note: Although the specification requires the method and the body to remain
unchanged when the redirection is performed, not all user-agents meet this
requirement. Use the 301 code only as a response for GET or HEAD methods and use
the 308 Permanent Redirect for POST methods instead, as the method change is
explicitly prohibited with this status.

Status
------

301 Moved Permanently

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301
`
	statusFound302 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 302 Found redirect status response code
indicates that the resource requested has been temporarily moved to the URL
given by the Location header. A browser redirects to this page but search
engines don't update their links to the resource (in 'SEO-speak', it is said
that the 'link-juice' is not sent to the new URL).

Even if the specification requires the method (and the body) not to be altered
when the redirection is performed, not all user-agents conform here - you can
still find this type of bugged software out there. It is therefore recommended
to set the 302 code only as a response for GET or HEAD methods and to use 307
Temporary Redirect instead, as the method change is explicitly prohibited in
that case.

In the cases where you want the method used to be changed to GET, use 303 See
Other instead. This is useful when you want to give a response to a PUT method
that is not the uploaded resource but a confirmation message such as:
'you successfully uploaded XYZ'.

Status
------

302 Found

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302
`
	statusSeeOther303 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 303 See Other redirect status response
code indicates that the redirects don't link to the requested resource itself,
but to another page (such as a confirmation page, a representation of a
real-world object — see HTTP range-14 — or an upload-progress page). This
response code is often sent back as a result of PUT or POST. The method used to
display this redirected page is always GET.

Status
------

303 See Other

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/303
`
	statusNotModified304 = `
Description
-----------

The HTTP 304 Not Modified client redirection response code indicates that there
is no need to retransmit the requested resources. It is an implicit redirection
to a cached resource. This happens when the request method is a safe method,
such as GET or HEAD, or when the request is conditional and uses an
If-None-Match or an If-Modified-Since header.

The response must not contain a body and must include the headers that would
have been sent in an equivalent 200 OK response: Cache-Control,
Content-Location, Date, ETag, Expires, and Vary.

Status
------

304 Not Modified

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304
`
	statusUseProxy305          = ``
	_306                       = `-`
	statusTemporaryRedirect307 = `
Description
-----------

HTTP 307 Temporary Redirect redirect status response code indicates that the
resource requested has been temporarily moved to the URL given by the Location
headers.

The method and the body of the original request are reused to perform the
redirected request. In the cases where you want the method used to be changed to
GET, use 303 See Other instead. This is useful when you want to give an answer
to a PUT method that is not the uploaded resources, but a confirmation message
(like "You successfully uploaded XYZ").

The only difference between 307 and 302 is that 307 guarantees that the method
and the body will not be changed when the redirected request is made. With 302,
some old clients were incorrectly changing the method to GET: the behavior with
non-GET methods and 302 is then unpredictable on the Web, whereas the behavior
with 307 is predictable. For GET requests, their behavior is identical.

Status
------

307 Temporary Redirect

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307
`
	statusPermanentRedirect308 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 308 Permanent Redirect redirect status
response code indicates that the resource requested has been definitively moved
to the URL given by the Location headers. A browser redirects to this page and
search engines update their links to the resource (in 'SEO-speak', it is said
that the 'link-juice' is sent to the new URL).

The request method and the body will not be altered, whereas 301 may incorrectly
sometimes be changed to a GET method.

Note: Some Web applications may use the 308 Permanent Redirect in a non-standard
way and for other purposes. For example, Google Drive uses a 308 Resume
Incomplete response to indicate to the client when an incomplete upload stalled.

Status
------

308 Permanent Redirect

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308
`

	statusBadRequest400 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 400 Bad Request response status code
indicates that the server cannot or will not process the request due to
something that is perceived to be a client error (for example, malformed request
syntax, invalid request message framing, or deceptive request routing).

Warning: The client should not repeat this request without modification.

Status
------

400 Bad Request

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400
`
	statusUnauthorized401 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 401 Unauthorized response status code
indicates that the client request has not been completed because it lacks valid
authentication credentials for the requested resource.

This status code is sent with an HTTP WWW-Authenticate response header that
contains information on how the client can request for the resource again after
prompting the user for authentication credentials.

This status code is similar to the 403 Forbidden status code, except that in
situations resulting in this status code, user authentication can allow access
to the resource.

Status
------

401 Unauthorized

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
`
	statusPaymentRequired402 = `
Description
-----------

Experimental: This is an experimental technology
Check the Browser compatibility table carefully before using this in production.

The HTTP 402 Payment Required is a nonstandard response status code that is
reserved for future use. This status code was created to enable digital cash or
(micro) payment systems and would indicate that the requested content is not
available until the client makes a payment.

Sometimes, this status code indicates that the request cannot be processed until
the client makes a payment. However, no standard use convention exists and
different entities use it in different contexts.

Status
------

402 Payment Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402
`
	statusForbidden403 = `
Description
-----------

The HTTP 403 Forbidden response status code indicates that the server
understands the request but refuses to authorize it.

This status is similar to 401, but for the 403 Forbidden status code,
re-authenticating makes no difference. The access is tied to the application
logic, such as insufficient rights to a resource.

Status
------

403 Forbidden

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403
`
	statusNotFound404 = `
Description
-----------

The HTTP 404 Not Found response status code indicates that the server cannot
find the requested resource. Links that lead to a 404 page are often called
broken or dead links and can be subject to link rot.

A 404 status code only indicates that the resource is missing: not whether the
absence is temporary or permanent. If a resource is permanently removed, use
the 410 (Gone) status instead.

Status
------

404 Not Found

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
`
	statusMethodNotAllowed405 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 405 Method Not Allowed response status
code indicates that the server knows the request method, but the target resource
doesn't support this method.

The server must generate an Allow header field in a 405 status code response.
The field must contain a list of methods that the target resource currently
supports.

Status
------

405 Method Not Allowed

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
`
	statusNotAcceptable406 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 406 Not Acceptable client error response
code indicates that the server cannot produce a response matching the list of
acceptable values defined in the request's proactive content negotiation
headers, and that the server is unwilling to supply a default representation.

Proactive content negotiation headers include:

* Accept
* Accept-Encoding
* Accept-Language

In practice, this error is very rarely used. Instead of responding using this
error code, which would be cryptic for the end user and difficult to fix,
servers ignore the relevant header and serve an actual page to the user. It is
assumed that even if the user won't be completely happy, they will prefer this
to an error code.

If a server returns such an error status, the body of the message should contain
the list of the available representations of the resources, allowing the user to
choose among them.

Status
------

406 Not Acceptable

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406
`
	statusProxyAuthRequired407 = `
Description
-----------

The HTTP 407 Proxy Authentication Required client error status response code
indicates that the request has not been applied because it lacks valid
authentication credentials for a proxy server that is between the browser and
the server that can access the requested resource.

This status is sent with a Proxy-Authenticate header that contains information
on how to authorize correctly.

Status
------

407 Proxy Authentication Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407
`
	statusRequestTimeout408 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 408 Request Timeout response status code
means that the server would like to shut down this unused connection. It is sent
on an idle connection by some servers, even without any previous request by the
client.

A server should send the "close" Connection header field in the response, since
408 implies that the server has decided to close the connection rather than
continue waiting.

This response is used much more since some browsers, like Chrome, Firefox 27+,
and IE9, use HTTP pre-connection mechanisms to speed up surfing.

Status
------

408 Request Timeout

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
`
	statusConflict409 = `
Description
-----------

The HTTP 409 Conflict response status code indicates a request conflict with
the current state of the target resource.

Conflicts are most likely to occur in response to a PUT request. For example,
you may get a 409 response when uploading a file that is older than the existing
one on the server, resulting in a version control conflict.

Status
------

409 Conflict

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409
`
	statusGone410 = `
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
`
	statusLengthRequired411 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 411 Length Required client error response
code indicates that the server refuses to accept the request without a defined
Content-Length header.

Note: by specification, when sending data in a series of chunks, the
Content-Length header is omitted and at the beginning of each chunk you need to
add the length of the current chunk in hexadecimal format. See Transfer-Encoding
for more details.

Status
------

411 Length Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411
`
	statusPreconditionFailed412 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 412 Precondition Failed client error
response code indicates that access to the target resource has been denied. This
happens with conditional requests on methods other than GET or HEAD when the
condition defined by the If-Unmodified-Since or If-None-Match headers is not
fulfilled. In that case, the request, usually an upload or a modification of a
resource, cannot be made and this error response is sent back.

Status
------

412 Precondition Failed

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412
`
	statusRequestEntityTooLarge413 = `
The HTTP 413 Content Too Large response status code indicates that the request
entity is larger than limits defined by server; the server might close the
connection or return a Retry-After header field.

Prior to RFC 9110 the response phrase for the status was Payload Too Large. That
name is still widely used.

Status
------

413 Payload Too Large

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413
`
	statusRequestURITooLong414 = `
Description
-----------

The HTTP 414 URI Too Long response status code indicates that the URI requested
by the client is longer than the server is willing to interpret.

There are a few rare conditions when this might occur:

* when a client has improperly converted a POST request to a GET request with
  long query information,

* when the client has descended into a loop of redirection (for example, a
  redirected URI prefix that points to a suffix of itself),

* or when the server is under attack by a client attempting to exploit
  potential security holes.

Status
------

414 URI Too Long

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414
`
	statusUnsupportedMediaType415 = `
Description
-----------

The HTTP 415 Unsupported Media Type client error response code indicates that
the server refuses to accept the request because the payload format is in an
unsupported format.

The format problem might be due to the request's indicated Content-Type or
Content-Encoding, or as a result of inspecting the data directly.

Status
------

415 Unsupported Media Type

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415
`
	statusRequestedRangeNotSatisfiable416 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 416 Range Not Satisfiable error response
code indicates that a server cannot serve the requested ranges. The most likely
reason is that the document doesn't contain such ranges, or that the Range
header value, though syntactically correct, doesn't make sense.

The 416 response message contains a Content-Range indicating an unsatisfied
range (that is a '*') followed by a '/' and the current length of the resource.
E.g. Content-Range: bytes */12777

Faced with this error, browsers usually either abort the operation (for example,
a download will be considered as non-resumable) or ask for the whole document
again.

Status
------

416 Range Not Satisfiable

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416
`
	statusExpectationFailed417 = `
Description
-----------

The HTTP 417 Expectation Failed client error response code indicates that the
expectation given in the request's Expect header could not be met.

See the Expect header for more details.

Status
------

417 Expectation Failed

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417
`
	statusTeapot418 = `
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
`
	statusMisdirectedRequest421 = `
Description
-----------

The HTTP 421 Misdirected Request client error response code indicates that the
request was directed to a server that is not able to produce a response. This
might be possible if a connection is reused or if an alternative service is
selected.

Status
------

421 Misdirected Request

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/421
`
	statusUnprocessableEntity422 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 422 Unprocessable Content response status
code indicates that the server understands the content type of the request
entity, and the syntax of the request entity is correct, but it was unable to
process the contained instructions.

Warning: The client should not repeat this request without modification.

Status
------

422 Unprocessable Entity

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422
`
	statusLocked423 = `
Description
-----------

Note: The ability to lock a resource is specific to some WebDAV servers.
Browsers accessing web pages will never encounter this status code; in the
erroneous cases it happens, they will handle it as a generic 400 status code.

The HTTP 423 Locked error response code indicates that either the resources
tentatively targeted by is locked, meaning it can't be accessed. Its content
should contain some information in WebDAV's XML format.

Status
------

423 Locked

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/423
`
	statusFailedDependency424 = `
Description
-----------

The HTTP 424 Failed Dependency client error response code indicates that the
method could not be performed on the resource because the requested action
depended on another action, and that action failed.

Regular web servers will normally not return this status code. But some other
protocols, like WebDAV, can return it. For example, in WebDAV, if a PROPPATCH
request was issued, and one command fails then automatically every other command
will also fail with 424 Failed Dependency.

Status
------

424 Failed Dependency

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/424
`
	statusTooEarly425 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 425 Too Early response status code
indicates that the server is unwilling to risk processing a request that might
be replayed, which creates the potential for a replay attack.

Status
------

425 Too Early

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425
`
	statusUpgradeRequired426 = `
Description
-----------

The HTTP 426 Upgrade Required client error response code indicates that the
server refuses to perform the request using the current protocol but might be
willing to do so after the client upgrades to a different protocol.

The server sends an Upgrade header with this response to indicate the required
protocol(s).

Status
------

426 Upgrade Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426
`
	statusPreconditionRequired428 = `
Description
-----------

The HTTP 428 Precondition Required response status code indicates that the
server requires the request to be conditional.

Typically, this means that a required precondition header, such as If-Match, is
missing.

When a precondition header is not matching the server side state, the response
should be 412 Precondition Failed.

Status
------

428 Precondition Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428
`
	statusTooManyRequests429 = `
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
`
	statusRequestHeaderFieldsTooLarge431 = `
Description
-----------

The HTTP 431 Request Header Fields Too Large response status code indicates that
the server refuses to process the request because the request's HTTP headers are
too long. The request may be resubmitted after reducing the size of the request
headers.

431 can be used when the total size of request headers is too large, or when a
single header field is too large. To help those running into this error,
indicate which of the two is the problem in the response body — ideally, also
include which headers are too large. This lets users attempt to fix the problem,
such as by clearing their cookies.

Servers will often produce this status if:

* The Referer URL is too long
* There are too many Cookies sent in the request

Status
------

431 Request Header Fields Too Large

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431
`
	statusUnavailableForLegalReasons451 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 451 Unavailable For Legal Reasons client
error response code indicates that the user requested a resource that is not
available due to legal reasons, such as a web page for which a legal action has
been issued.

Status
------

451 Unavailable For Legal Reasons

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451
`

	statusInternalServerError500 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 500 Internal Server Error server error
response code indicates that the server encountered an unexpected condition that
prevented it from fulfilling the request.

This error response is a generic "catch-all" response. Usually, this indicates
the server cannot find a better 5xx error code to response. Sometimes, server
administrators log error responses like the 500 status code with more details
about the request to prevent the error from happening again in the future.

Status
------

500 Internal Server Error

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
`
	statusNotImplemented501 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 501 Not Implemented server error response
code means that the server does not support the functionality required to
fulfill the request.

This status can also send a Retry-After header, telling the requester when to
check back to see if the functionality is supported by then.

501 is the appropriate response when the server does not recognize the request
method and is incapable of supporting it for any resource. The only methods that
servers are required to support (and therefore that must not return 501) are GET
and HEAD.

If the server does recognize the method, but intentionally does not support it,
the appropriate response is 405 Method Not Allowed.

Status
------

501 Not Implemented

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501
`
	statusBadGateway502 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 502 Bad Gateway server error response
code indicates that the server, while acting as a gateway or proxy, received an
invalid response from the upstream server.

Note: A Gateway might refer to different things in networking and a 502 error is
usually not something you can fix, but requires a fix by the web server or the
proxies you are trying to get access through.

Status
------

502 Bad Gateway

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502
`
	statusServiceUnavailable503 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 503 Service Unavailable server error
response code indicates that the server is not ready to handle the request.

Common causes are a server that is down for maintenance or that is overloaded.
This response should be used for temporary conditions and the Retry-After HTTP
header should, if possible, contain the estimated time for the recovery of the
service.

Note: together with this response, a user-friendly page explaining the problem
should be sent.

Caching-related headers that are sent along with this response should be taken
care of, as a 503 status is often a temporary condition and responses shouldn't
usually be cached.

Status
------

503 Service Unavailable

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
`
	statusGatewayTimeout504 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 504 Gateway Timeout server error response
code indicates that the server, while acting as a gateway or proxy, did not get
a response in time from the upstream server that it needed in order to complete
the request.

Note: A Gateway might refer to different things in networking and a 504 error is
usually not something you can fix, but requires a fix by the web server or the
proxies you are trying to get access through.

Status
------

504 Gateway Timeout

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504
`
	statusHTTPVersionNotSupported505 = `
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
`
	statusVariantAlsoNegotiates506 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 506 Variant Also Negotiates response
status code may be given in the context of Transparent Content Negotiation
(see RFC 2295). This protocol enables a client to retrieve the best variant of
a given resource, where the server supports multiple variants.

The Variant Also Negotiates status code indicates an internal server
configuration error in which the chosen variant is itself configured to engage
in content negotiation, so is not a proper negotiation endpoint.

Status
------

506 Variant Also Negotiates

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506
`
	statusInsufficientStorage507 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 507 Insufficient Storage response status
code may be given in the context of the Web Distributed Authoring and Versioning
(WebDAV) protocol (see RFC 4918).

It indicates that a method could not be performed because the server cannot
store the representation needed to successfully complete the request.

Status
------

507 Insufficient Storage

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507
`
	statusLoopDetected508 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 508 Loop Detected response status code
may be given in the context of the Web Distributed Authoring and Versioning
(WebDAV) protocol.

It indicates that the server terminated an operation because it encountered an
infinite loop while processing a request with "Depth: infinity". This status
indicates that the entire operation failed.

Status
------

508 Loop Detected

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508
`
	statusNotExtended510 = `
Description
-----------

The HyperText Transfer Protocol (HTTP) 510 Not Extended response status code is
sent in the context of the HTTP Extension Framework, defined in RFC 2774.

In that specification a client may send a request that contains an extension
declaration, that describes the extension to be used. If the server receives
such a request, but any described extensions are not supported for the request,
then the server responds with the 510 status code.

Status
------

510 Not Extended

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510
`
	statusNetworkAuthenticationRequired511 = `
Description
-----------

The HTTP 511 Network Authentication Required response status code indicates that
the client needs to authenticate to gain network access.

This status is not generated by origin servers, but by intercepting proxies that
control access to the network.

Network operators sometimes require some authentication, acceptance of terms, or
other user interaction before granting access (for example in an internet café
or at an airport). They often identify clients who have not done so using
their Media Access Control (MAC) addresses.

Status
------

511 Network Authentication Required

Source
------

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511
`
)

var statusCodes = []string{
	"------------------ 1xx ------------------", // category header
	"100",
	"101",
	"102",
	"103",

	"------------------ 2xx ------------------",
	"200",
	"201",
	"202",
	"203",
	"204",
	"205",
	"206",
	"207",
	"208",
	"226",

	"------------------ 3xx ------------------",
	"300",
	"301",
	"302",
	"303",
	"304",
	"305",
	"306",
	"307",
	"308",

	"------------------ 4xx ------------------",
	"400",
	"401",
	"402",
	"403",
	"404",
	"405",
	"406",
	"407",
	"408",
	"409",
	"410",
	"411",
	"412",
	"413",
	"414",
	"415",
	"416",
	"417",
	"418",
	"421",
	"422",
	"423",
	"424",
	"425",
	"426",
	"428",
	"429",
	"431",
	"451",

	"------------------ 5xx ------------------",
	"500",
	"501",
	"502",
	"503",
	"504",
	"505",
	"506",
	"507",
	"508",
	"510",
	"511",
}

var statusCodeMap = map[string]string{
	// "------------------ 1xx ------------------"
	"100": statusContinue100,
	"101": statusSwitchingProtocols101,
	"102": statusProcessing102,
	"103": statusEarlyHints103,

	// "------------------ 2xx ------------------"
	"200": statusOK200,
	"201": statusCreated201,
	"202": statusAccepted202,
	"203": statusNonAuthoritativeInfo203,
	"204": statusNoContent204,
	"205": statusResetContent205,
	"206": statusPartialContent206,
	"207": statusMultiStatus207,
	"208": statusAlreadyReported208,
	"226": statusIMUsed226,

	// "------------------ 3xx ------------------"
	"300": statusMultipleChoices300,
	"301": statusMovedPermanently301,
	"302": statusFound302,
	"303": statusSeeOther303,
	"304": statusNotModified304,
	"305": statusUseProxy305,
	"306": _306,
	"307": statusTemporaryRedirect307,
	"308": statusPermanentRedirect308,

	// "------------------ 4xx ------------------"
	"400": statusBadRequest400,
	"401": statusUnauthorized401,
	"402": statusPaymentRequired402,
	"403": statusForbidden403,
	"404": statusNotFound404,
	"405": statusMethodNotAllowed405,
	"406": statusNotAcceptable406,
	"407": statusProxyAuthRequired407,
	"408": statusRequestTimeout408,
	"409": statusConflict409,
	"410": statusGone410,
	"411": statusLengthRequired411,
	"412": statusPreconditionFailed412,
	"413": statusRequestEntityTooLarge413,
	"414": statusRequestURITooLong414,
	"415": statusUnsupportedMediaType415,
	"416": statusRequestedRangeNotSatisfiable416,
	"417": statusExpectationFailed417,
	"418": statusTeapot418,
	"421": statusMisdirectedRequest421,
	"422": statusUnprocessableEntity422,
	"423": statusLocked423,
	"424": statusFailedDependency424,
	"425": statusTooEarly425,
	"426": statusUpgradeRequired426,
	"428": statusPreconditionRequired428,
	"429": statusTooManyRequests429,
	"431": statusRequestHeaderFieldsTooLarge431,
	"451": statusUnavailableForLegalReasons451,

	// "------------------ 5xx ------------------"
	"500": statusInternalServerError500,
	"501": statusNotImplemented501,
	"502": statusBadGateway502,
	"503": statusServiceUnavailable503,
	"504": statusGatewayTimeout504,
	"505": statusHTTPVersionNotSupported505,
	"506": statusVariantAlsoNegotiates506,
	"507": statusInsufficientStorage507,
	"508": statusLoopDetected508,
	"510": statusNotExtended510,
	"511": statusNetworkAuthenticationRequired511,
}
