# HTTP What Every Web Developer Should Know About HTTP by K. Scott Allen

## Chapter 1: Introduction

> The path of a URL is case sensitive

However, many websites will try to make URLs behave as if they are
case-insensitive to avoid broken links and allows users to find
content even if they mistakenly type.

**Tips:** Some web servers, notably the web server on Microsoft Windows
Operating Systems, are case insensitive by default.

> URL Encoding

Unsafe characters: space, #, ^(it isn't always transmitted correctly through all network devices) ...

Fortunately, we can still transmit unsafe characters in a URL, **but
all unsafe characters must be percent encoded**. %20 is encoding for
a space character.

> Ports, Query Strings and Fragments

`<scheme>://<host>:<port>/<path>?<query>#<fragment>`

- Fragment

`http://server.com?recipe=broccoli#ingredients`

The part after number sign is known as *fragment*.

The fragment is not processed by the server. The fragment
is only used on the client and it identifies a particular section
of resource. Specifically, the fragment is typically used to identify a
specific HTML element in a page by the element's ID.

For example, the URL `http://odetocode.com/Blogs/scott/archieve/2011/11/29/programming-windows-8-the-sublime-to-the-strange.aspx#feedback`

Your web browser should scroll down the page to show the feedback
section of a blog post.

```html
<div id="post">

</div>

<div id="feedback"> <<< HERE

</div>

```

> Content Type

primaryMediaType/SubType

text/html
image/jpeg
image/png

Those content types are standard MIME types and are literally what will
appear in the HTTP response.

> File Extensions

Microsoft documentation for IE says IE will first look at the MIME type
specified by the host. If the host does not provide a MIME type, IE
will then scan the first 200 bytes of the response trying to guess the
content type. Finally, if IE doesn't find a type and can't guess the type,
it will fall back on the file extension used in the request for the resource.

This is one reason why the content type label is important, but it is far
from the only reason.

> Content Type Negotiation

A resource identified by a single URL can have multiple representation, e.g. in different languages, format

When a client makes an HTTP request to a URL, the client can specify the
**media types** it will accept. **It's up to the server to try to fulfill the request.**. This is why we call it content negotiation and not a content ultimatum.

## Chapter 2: Messages

> Full HTTP request message --- always in ASCII text

[method] [URL] [version]

[headers]

[body]

> Full HTTP response

[version] [status] [reason]

[headers]

[body]

> A Raw Request and Response

The message is a command in plain ASCII text and formatted per the HTTP specification.

> Using Telnet to make HTTP requests

- Normal Telnet session connects over port 23

```bash
$ telnet www.odetocode.com 80
...
GET / HTTP/1.1
host:www.odetocode.com
<Enter>
<Enter>
```

**Tips:** The host information is a **required** piece of information in an HTTP 1.1 request message.
This is to help servers that support multiple web sites to deliver the message to the correct
sites.

*Web servers and networking infrastructure can be expensive, and for small web sites,
we don't want to run a single website on a single server. Both <www.odetocode.com> and
<www.odetofood.com> might live on a same physical server.

> Redirect (301/302)

Response header: `Location: http://odetocode.com`

It's up to the client now to parse this response message and send a request
to odetocode.com instead of <www.odetocode.com>. Any web browser will go to the new location
automatically.

The redirect forces search engines to see one true URL for a given resource, what
we call the canonical URL, because having a canonical URL will improve search
result rankings for a given resource.

> HTTP Request Methods

| Method | Description |
| --- | --- |
| GET | Retrieve a resource |
| PUT | Store a resource |
| DELETE | Remove a resource |
| POST | Update a resource |
| HEAD | Retrieve the headers or metadata for a resource |

**GET and Safety**: Safe methods, as the name implies, do not do anything dangerous
like destroy a resource, submit a credit card transaction, or cancel an account.

A GET should only retrieve a resource and not alter the state of the resource.

A POST is not a safe method.

It's OK to refresh a web page retrieved by a GET request. **However, if the page
is the response of an HTTP POST request, the browser will warn us if we try to refresh the
page.**

Because of this, many web app always try to leave the client viewing the result of a
GET request. **After a user clicks a button to POST information to a server, the server will
process the information and response with an HTTP redirect**. Tells the browser to GET the result of the previous POST operation.

e.g. customer see "thank you" message, which can be refresh safely

It is POST/Redirect/GET pattern (PRG)

> Three Common HTTP Response

1. HTML

    - Viewing the result of POST request. A refresh might try to sign him up a second time!

2. Redirect message

    - force browser to issue a safe GET request for a page.

3. Error

> Request Headers

| Key | Value | Description|
| --- | --- | --- |
| Date | Fri, 9 Aug 2014 21:12:00 GMT | RFC822 format |
| Referrer | <http://www.google.com/url?&q=odetocode> | When the user clicks on a link, the client can send te URL of the referring page in this header |
| User-Agent | Mozilla/5.0 (Windows NT 6.1; WOW64) Chrome/16.0.912.75 Safari/535.7 | Information about the user agent (the software) making the request (Edge Versus Chrome versus Firefox) |
| Accept | texthtml,application/xhtml+xml,application/xml;q=0.9,*.*;q=0.8 | Describes the media types the user-agent is willing to accept. This header is used for content negotiation |
| Accept-Encoding | gzip,deflat,sdch | |
| Accept-Language | fr-FR,en;q=0.8 | client wants the see resource in French |
| Accept-Charset | ISO-8859-1,utf-8;q=0.7,*;q=0.3 | |
| Cookie | | Cookie information. Helps a server track or identify a user |
| If-Modified-Since | | Contain a date of when the user-agent last retrieved (and cached) the resource. The server only has to send back te entire resource if it's been modified since that time. |

**Tips:** "q" value is always a number between 0 and 1, it represents the **quality value** or "relative degree of preference" for a value. The default is 1.0

> Response Headers

| Key | Value | Description|
| --- | --- | --- |
| Cache-Control | private | |
| Content-Type | text/html; charset=utf-8 | character set used to encode the HTML is UTF-8 |
| Content-Length | 17151 | |
| Server | Microsoft-IIS/7.0 | |
| X-AspNet-Version | 2.0.50727 | |
| X-Powered-By | ASP.NET | |
| Date | Fri, 9 Aug 2014 21:12:00 GMT | RFC822 format |
| Connection | close | |
| Location | <http://odetocode.com> | |

> Response Status Code

| Range | Category |
| --- | --- |
| 100-199 | Informational |
| 200-299 | Successful |
| 300-399 | Redirection |
| 400-499 | Client Error |
| 500-599 | Server Error |

| Code | Reason | Description |
| --- | --- | --- |
| 200 | OK | The status code everyone wants to see. A 200 code in the response
means everything worked! |
| 301 | Move Permanently | The resource has moved to the URL specified in Location header |
| 302 | Move Temporarily | The resource has moved to the URL specified in Location header temporarily. Applications typically use this response code **after a successful POST operation** -- PRG |
| 304 | Not Modified | This is the server telling the the client resource hasn't changed since the last time the client retrieved the resource. The client can use a cached copy of the resousrce.
| 400 | Bad Request | Server could not understand the request, probably an incorrect syntax or sent inappropriate data |
| 403 | Forbidden | Server refused access to resource due to a user does not have permissions. Access Denied! |
| 404 | Not Found | Server could not locate the resource, e.g. a typo in URL |
| 500 | Internal Server Error | There is a programming error, database is offline. |
| 503 | Service Unavailable | Server will currently not service the request. This status code can appear when a server is throttling requests because the server is under heavy load. |

***Versus Your Application***

If it is a web application. When user forgot to fill out the field for his last name, you want
the application to return some content to the client with a 200(OK) status code, HTML content
tells the user they forgot to provide a last name. From an application perspective, the request
was a failure, but from an HTTP perspective, the HTTP transaction was a success.

If it is a web service, reply 400 and specify the missing field

> Caching and Performance Optimizations

- ETag, Expires, and Last-Modified all provide information about the cache-ability of a response

Etag -- is an identifier that will change when the underlying resource changes.
So comparing ETags is an efficient way for a client to know if a cached resource is stale.

Expires -- tells a client how long to cache a resource

## Chapter 3: Connections

### Application Layer Protocol

> HTTP is what we call an application layer protocol.

Because HTTP allows two applications to communicate over the network. (Often one is a web browser
, the other is a webserver like IIS or Apache)

### How a connection works

| From | Protocol | To |
| --- | --- | --- |
| Application | HTTP | Application |
| Transport | TCP | Transport |
| Network | IP | Network |
| Data Link | Ethernet | Data Link |
| | Media | |

Almost all HTTP traffic travels over TCP (Transmission Control protocol), although this is not required by HTTP.

When user types a URL into the browser

1. Browser extracts the host name from the URL (and port number, if any)

2. Opens a TCP socket by specifying the server address and port

3. TCP is responsible for error detection, flow control, and overall reliability.
TCP **guarantees the delivery of the message** to the other server.
It resend automatically that might get lost in transit. It is a **reliable protocol**.

4. TCP also provides flow control, this algorithm ensure the sender does not send data
too fast for the receiver to process the data. **Tips:** [you can observer TCP handshakes here](#wireshark)

5. IP (Internet Protocol) is responsible for taking pieces of information and moving them through the various switches, routers, gateways, repeaters, and other devices that move information
from one network to the next and all around the world. It does not guarantee delivery.

6. IP requires every device to have an address, an IP address (e.g. 208.192.32.40). It is also
responsible for breaking data into packets (often called datagrams), and sometimes fragmenting and reassembling these packets so they are optimized for a network.

7. Data link layer -- IP packets must travel over a piece of wire, a fire optic cable, a wireless network or a satellite.

### Connections

#### Parallel Connections

> Web browsers can open multiple, **parallel connections** to a server.

The number of parallel connections depends on the browser and the browser's configuration.

**Note:** IE was only obeying the rules spelled out in the HTTP 1.1 specification:
A single-user client SHOULD NOT maintain more than two connections with any server or proxy

Things are different today. Most browsers today will open at least six concurrent connections.

Too many connections can saturate and congest the network. A single server can only accept a
finite number of connections.

Fortunately, parallel connections are not the only performance optimization.

#### Persistent Connections

As the number of requests per page grew, so did the overhead generated by TCP handshakes and the
in-memory data structures required to establish each TCP socket.

To reduce this overhead and improve performance, HTTP 1.1 specification suggests that clients and
servers should implement persistent connections and make persistent connections the default type of connection.

A persistent connection stays open after the completion of one request-response transaction. It reduce memory usage, CPU usage, network congestion, latency, and generally improve the response time of a page.

However, a server can only support a finite number of incoming connections. Many servers have a default configuration to limit the maximum number of current connections. The configuration is a security measure to help prevent denial of service (DoS) attacks. A malicious person can initiate a DoS attack by creating programs that open thousands of persistent connections to a server, hence inhibit the server from responding to real customers.

Because a server only support a finite number of connections, most webservers will close a persistent
connection if it is idle for some period. The default Apache configuration says to close a persistent
connection if the connection is idle for 5s. (To see whether a forced close, use network analyzer [Wireshark](#wireshark))

**Because persistent connections are the default connection style with HTTP 1.1,
a server that does not allow persistent connections must include a Connection header in every
HTTP response.**

```bash
HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Connection: close
...
```

With this response, browser will not be able to make a second request on the same connection.

#### Pipelined Connections

Future specifications will allow for pipelined connections.

A browser can send multiple HTTP requests on a connection before waiting for the first response. It
allows for a more efficient packing of requests into packets and can reduce latency.

## Chapter 4: Web Architecture

### What we can NOT do with a URL

1. Cannot restrict a client or a server to a specific type of technology.

2. Cannot force the server to store a source using any specific technology.

3. Cannot specify the representation of a specific resource,

4. Cannot say what a user wants to do with a resource.

### HTTP adding value

The web server (e.g. Apache) will inspect information in a message, like the URL or the host
header, when deciding where to send a message. It can also perform additional actions with the message,
like logging the message to a local file. The web applications on the server don't need to worry
about logging because the serer is configured to log all messages.

Likewise, a server can know if a client supports gzip compression because a client can advertise this fact through an `accept-encoding` header in the HTTP request. Compression allows a server
to take a 100kb resource and turn it into 25kb resource for a faster transmission,

**Tip:** you can configure many web server to automatically use compression for certain content types.

### Proxies

> A proxy server sits between a client and server. Before forwarding messages, the proxy can inspect the
messages and potentially take some additional actions.

E.g. A company uses a proxy server to capture all HTTP traffic leaving the office.
They don't want employees and contractors spending any time on social media, so HTTP requests to those
servers will never reach their destination.

A proxy server could also inspect messages to remove confidential data, like the referrer
headers that point to internal resources on the company network.

An access control proxy can also log HTTP messages to create audit trails on all traffic. Many
access control proxies require user authentication to ensure only authorized users have external network access.

**Forward proxy**: closer to client than the server. It usually require some configuration in the client software or
web browser to work.

**Reverse proxy**: closer to the server than the client and is completely transparent to the client.

Some other popular proxy services includes:

- **Load balancing**: take a message and forward it to one of several web servers
on a round-robin basis, or other ways.

- **SSL acceleration**: encrypt and decrypt HTTP messages, taking the encryption load off a
web server.

- **Provide additional layer of Security**: filtering out potentially dangerous HTTP messages
such as XSS vulnerability or launch a SQL injection attack.

- **Caching proxies**: store copies of frequently accessed resources. The proxy can respond
to messages requesting a cached resource by returning the cached resource directly.

### Caching

> There are two types of caches: public cache and private cache

**Public cache**: is a cache shared among multiple users. A public cache generally resides
on a proxy server.

A public cache on a forward proxy is usually caching the resources that are popular in
a community of users, like the users of a specific company.

A public cache on a reverse proxy is commonly caching the resources that are popular on a
specific web site, like popular product images from Amazon.com

**Private cache**: is a cache for a single user. Web browsers always keep a private
cache of resources on your disk.

> In HTTP 1.1, a response message with a 200 (OK) status code for an HTTP
GET request is cacheable by default

An application can influence this default by using the proper headers
in an HTTP response: **Cache-control**

**Expires** is still around and widely supported despite the HTTP 1.1 specification
deprecating the header.

**Pragma** is another example of a header used to control caching behavior, but it
too is only around for backward compatibility.

#### Cache-Control

Have a value for

- public: means proxy servers can cache the response. e.g. company logo

- private: only user's browser can cache the response. e.g. responses customized to a specific user

- no-cache: no one should cache the response

- no-store: the message might contain sensitive info, and not only
should the message not be cached or saved, but the browser should
remove the message from memory as soon as possible.

Server can also specify a **max-age** value in the Cache-control header.
It is the number of seconds to cache the response.

```bash
HTTP/1.1 200 OK
Last-Modified: Wed, 25 Jan 2012 17:55:15 GMT
Expires: Sat, 22 Jan 2022 17:55:15 GMT
Cache-Control: max-age=3153600000,public
# it is 10 years
```

**Tip:** If a client is HTTP 1.1 compliant and understands Cache-control,
the client should use the value in max-age instead of Expires.

A **Last-Modified** header to indicate when the resource representation
last changed. Cache logic can use this value as a **validator**.
For example, if the browser decides it needs to check on the resource it can
issue the following request

```bash
GET ... HTTP/1.1
If-Modified-Since: Wed, 25 Jan 2012 17:55:15 GMT
```

It is telling the server the client only needs the full response if the resource has changed.
Else, **the server** can respond with a 304 - Not Modified Message. The client can use last
cached response. ![This is a good article (in Chinese)](https://developer.aliyun.com/article/919310)

```bash
HTTP/1.1 304 Not Modified
Expires: Sat, 22 Jan 2022 17:55:15 GMT
Cache-Control: max-age=3153600000,public
# telling the client "go ahead and use the bytes you already have cached."
```

#### ETag

Response:

```bash
HTTP1.1 200 OK
Server: Apache
Last-Modified: Fri, 06 Jan 2012 18:08:20 GMT
ETag: "8e5dcd-59f-4b5dfef104d00"
```

ETag is an opaque identifier, meaning the ETag does not have any inherent meaning.
It can be a timestamp, a GUID, or a value computed by running a hashing algorithm.

All the client knows is that if the resource ever changes,
the server will compute a **new ETag** value for the resource. Hence, the client
can validate if the resources are different with comparing ETag.

## Chapter 5: State and Security

> HTTP is a stateless protocol, meaning each request-response transaction is independent
of any previous or future transaction.

There are many options for storing state in a web application

1. Embed state in the resource transferred to client. This approach typically requires
some hidden input fields and works the best for short lived state e.g. for moving through
a three-page application process. This approach is a highly scalable approach, but it can
complicate the application programming.

2. Store the state on the server, this options is required for state that must be around
a long time. Many web development frameworks provide access to a **user session** which may live in memory or in a database. A developer can store information in the session and
retrieve the information on every subsequent request from the same user. Session storage has
an easy programming model but is only good for short-lived state. The server must assume the user
has left the site or closed his browser and the server will discard the session.

By using session storage, subsequent requests must go to the exact same server, there are some
load balancers help to support this scenario by implementing **sticky sessions**

### Session and Cookie

In early state: IP address of a request message. Nowadays, many user live behind
devices using Network Address Translation (multiple users on the same IP address)

Now: Identification and cookies

**Any response with Set-Cookie header should not be cached which interfere with user information and create security problems**

1. Server give a user cookie: **Set-Cookie** header

```bash
HTTP/1.1 200 OK
Set-Cookie: GUID=00a789bnvrew794vber;domain=.mywebsite.com;path=/
```

2. Browser will send the cookie to the server in **every subsequent HTTP request**

```bash
GET ... HTTP/1.1
Cookie: GUID=00a789bnvrew794vber;
```

3. When the ID arrives, the server software look up any associated user data from an in-memory
data structure, database or distributed cache. The session data stay alive based on default in different web application, e.g. 20min in ASP.NET. If there is no incoming cookie with a session ID, ASP.NET will create one with a Set-Cookie header.

**NOTE**: Cookies are defined by RFC6265

#### Cookies limitations

- Cookies does not necessarily authenticate users -- it just a unique identifier to differentiate one user from another.

- raise privacy concerns. Some users will disable cookies in their browsers, the browser will reject any cookies a server sends in a response. We can place the user identifier into URL.

#### Cookies Security Concern

- someone hijacking another user's session. e.g. if I use a tool like Fiddler to trace HTTP
traffic, I might see a Set-Cookie header and guess some other user already has a sessionID of 11 and create an HTTP request with that ID just to see if I can steal or view the HTML intended for some other user.

To combat this problem, most web application use large random numbers (e.g. ASP.NET uses 120bits of randomness)

- advertising sites use **third-party cookies** to track users across the internet. It get set from a different domain than the
domain in the browser's address bar. As an example, the home page at `server.com` can include a <script> tag
with a source set to bigadvertising.com. This allows bigadvertising.com to set a cookie when user is viewing content from
server.com

The cookie can only go back to bigadvertising.com, but if enough web sites use bigadvertising.com, then
Big Advertising can **start to profile individual users and the sites they visit**

Most web browsers will allow you to disable third-party cookies, but third-party cookies are on
by default.

- A cookie is visible as it travels across the network, and it sits on the file system (persistent cookie)

#### Types of Cookies

Above mentioned are **session cookies**.

**Persistent cookie** can outlive a single browsing session a browser will store the cookies
to disk. It implemented by an **Expires** attribute.

`Set-Cookie: name=value; expires=Monday, 09-July-2017 21:12:00 GMT`

#### Set-Cookie Attribute

> HttpOnly

Another security concern around cookies: vulnerable to XSS. ![Here is an excellent talk about XSS](https://www.youtube.com/watch?v=oEFPFc36weY)

1. A malicious user injects malevolent JavaScript code into someone else's website. This javascript was saved into the database.

2. If other web site sends the malicious script to their users, the malicious script
can modify, or inspect and steal cookie information

HttpOnly flag can tells the browser to prevent access to the cookie from the JavaScript.
The cookie exists only to travel out in the header of every HTTP request message. HttpOnly
will not allow JavaScript to read or write the cookie on the client.

> Discard

A session cookie can explicitly add a Discard attribute to a cookie, but
without an Expires value, the browser should discard the cookie in any case.

> Expires

Outlive a single browsing session

> Domain -- Control scope

`Set-Cookie: name=value; domain=.server.com; path=/stuff`

Above attribute allows a cookie to span sub-domains e.g <www.server.com>, image.server.com ...

**Important**: You can not use the domain attribute to span domains, so setting the
domain to .microsoft in a response to .server.com is not legal
and **the browser should reject the cookie**.

> Path -- Control scope

`Set-Cookie: name=value; domain=.server.com; path=/stuff`

Restrict a cookie to a specific resource path. In above example, the cookie
will only travel to a server.com site when the request URL is pointing to /stuff.
e.g /stuff/images.

### Authentication

#### Authentication protocols

`basic`, `digest`, `Windows`, `forms` and `OpenID`

**Tip**: only `basic` and `digest` authentication protocols are officially in HTTP specification.

##### Basic Authentication

if the user is not authenticated, the server will issue an **authentication challenge**

```bash
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Basic realm="localhost"
# WWW-Authenticate tells the client to collect the user credentials and try again.
# The realm attribute gives the client a string it can use as a description for the protected area.
```

Then most browsers can display a UI for user to enter credentials.

Then browser can send another request to the server

```bash
GET http://localhost/html5/ HTTP/1.1
Authorization: Basic basdfau89234yhuhaf
```

The value of the authorization header is the client's username
and password in a base64 encoding. **Basic authentication is insecure by default** because
anyone with a base 64 decoder can view the message and steal a user's password.

Therefore, basic authentication is rarely used without using secure HTTP.

##### Digest Authentication

Is an improvement over basic authentication because it does not transmit
user passwords using base 64 encoding. Instead, the client must send a
**digest** of the password. By using MD5 hashing algorithm with
a nonce the server provides during the authentication challenge.

**Tip**: A nonce is a cryptographic number used to help prevent
replay attacks.

The digest challenge response:

```bash
HTTP/1.0 401 Unauthorized
WWW-Authenticate: Digest realm="localhost",qop="auth,auth-int",nonce="dsad8dsd78asdd89gjrew",opaque="5ccc069c403ebaf8f9321"
```

Digest authentication is better than basic authentication when secure HTTP is not available, but it is
still far from perfect because the MD5 hashing algorithm is a weak form of encryption. Digest
authentication is also vulnerable to man in the middle attacks where someone is sniffing network traffic.

##### Windows Authentication

It is a standard authentication protocol among Microsoft products. However, many modern browsers and
not just IE support Windows Authentication. Windows Auth does not work well over the Internet.
You'll find Windows Auth is common on internal and intranet web sites where a Microsoft Active
Directory server exists.

```bash
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Negotiate
# NTLM or Negotiate(allows client to select Kerberos or NTLM)
```

Kerberos and NTLM is the underlying authentication protocols supported by Windows

Windows Auth has the advantage of being secure even without using secure HTTP and
of being unobtrusive for users of IE. IE will automatically authenticate a user
when a server challenges a request and will do so using the user's credentials
that they used to log into the Windows operating system.

##### Form-based Authentication

Is the most popular approach to user authentication over the Internet. It is not a standard
authentication protocol, and doesn't use WWW-Authenticate or Authorization headers.

With this approach, an application will respond to a request for a secure resource
by an anonymous user by redirecting the user to a login page. The redirect is HTTP 302 temporary
redirect.

**Tip**: Generally, the URL the user is requesting might be included into the query string of the redirect
location so that once the user has completed the login, the application can redirect
the user to the secure resource they were trying to reach.

```bash
HTTP/1.1 302 Found
Location: /Login.aspx?ReturnUrl=/Admin.aspx
```

Forms-based authentication is not secure unless you use secure HTTP.

##### OpenID and OAuth

Specifically, some applications don't want to manage and verify usernames and passwords, and
users don't want to have a different username and password for every website.
OpenID and OAuth are standards allowing the decentralized authentication.

With OpenID, a user registers with an OpenID identify provider (e.g Google, Yahoo),
and identify provider is the only site that needs to store and validate user credentials.

While OpenID and OAuth offer many benefits compared to forms authentication, they
do add complexity in implementing and debugging a web application.

### Secure HTTP -- HTTPS

It use https scheme in URL, default port is 443.

Secure HTTP encrypting messages before the messages start traveling across the network.
HTTPS add an security layer in the network protocol stack between HTTP and TCP layers,
and features the use of the Transport Layer Security protocol (TLS) or the
TLS predecessor known as Secure Sockets Layer (SSL).

HTTPS requires a server to have a cryptographic certificate. The server sends the
certificate to the client during setup of the HTTPS communication. The certificate
includes the server's host name. A browser use the certificate to validate
that the server is talking to is the correct server. The validation
happens using public key cryptography and certificate authorities, that will sign and
vouch for the integrity of a certificate. Server administrators must purchase
and install certificates from the certificate authorities.

**NOTE**: there is the possibility of using client-side certificate with HTTPS.
However, most sites on the internet won't use client-side certificates because
most users don't want the hassle and expense of purchasing and installing a personal
cetifiacte.

Corporations might require client certificates for employees to access
corporate servers. But in this case, the corporation can act as a
certificate authority and issue employees certificates they create
and manage.

HTTPS do:

- Encrypt all request and response traffic. Except hostname in URL, which means path
and query string, cookies are all encrypted. HTTPS prevents session hijacking because
no eavesdroppers can inspect a message and steal a cookie.

- The server certificate authenticates the server to the client. If you are talking
to abc.com over HTTPS, you can be sure your messages are really going to abc.com.
Evil person cannot modify request and response in a proxy server

- HTTPS does not authenticate the client.

HTTPS's downsides:

- Computationally expensive, large sites often use specialized hardware (SSL accelerators)
to take the cryptographic computation load off the web servers.

- Traffic is impossible to cache in a public cache.

- Expensive to set up and require an additional handshake between the client and server
to exchange cryptographic keys and ensure everyone is communicating with
proper secure protocol. Persistent connections can help to amortize this cost.

## Other

### Fiddler

> Fiddler allows you not only to see raw HTTP req and resp, but create and modify requests.

### Wireshark

> Wireshark allow you to have some visibility into TCP and IP.

You can observe TCP handshakes, which are the TCP messages required to establish a
connection between client and server before the actual HTTP messages start flowing.

###### TO EXPLORE

1. Proxy to capture all HTTP Traffic

2. Cache-Control header

3. "Last-Modified" With "If-Modified-Since"

4. "ETag" header With "If-None-Match"

5. span Domain cookie

6. XSS

    - Sanitizing (In both end)

    - Avoid using vulnerable third-party package

7. man in the middle attacks
