### Product Review

A review site is a website on which reviews can be posted about people, businesses, products, or services. These sites may use Web 2.0 techniques to gather reviews from site users or may employ professional writers to author reviews on the topic of concern for the site. Early review sites included first ConsumerDemocracy.com which introduced the helpfulness ratings, Hithen later Epinions.com and Amazon.com.

— [More at WikiPedia, Review site](https://en.wikipedia.org/wiki/Review_site)

![screenshot](screenshot.png)

### Installation

- Make sure you have [Go](https://golang.org/) installed,
- Download the project into your computer,
- `go get -u github.com/cixtor/middleware`
- `go build -o server -- src/*.go && ./server`
- [Open in your web browser](http://127.0.0.1:8080/)
- _(Optional)_ `curl -s "http://127.0.0.1:8080/reviews/list?uid=F79MEIM7" | python -m json.tool`

### Assumptions (Database)

The specification of the project didn't mention anything about the data storage format nor the lifetime of the data itself. I went ahead and assumed that the project manager wanted to keep the reviews in a long-lived state file (aka. database) and so I decided to choose one of the popular database engines, in this case SQLite.

**Why SQLite?** SQLite is a relational database management system contained in a C programming library. In contrast to many other database management systems, SQLite is not a client–server database engine. Rather, it is embedded into the end program.

There are advantages of using this over other popular SQL engines like MySQL, Postgresql, Cassandra, MongoDB, etc; one of the benefits is that the database can be easily moved around without the need of a reconfiguration, we can attached the file into a Zip archive and even push to a CVS like GitHub and the data will be safe.

```
id        - Primary key to keep track of the counter.
uid       - Unique identifier of the product that is being reviewed.
name      - Arbitrary name of the reviewer.
email     - E-mail address of the reviewer.
rating    - Integer in the range 0-10.
comment   - Arbitrary text with the review.
approved  - Either True or False (format to be defined).
timestamp - Number of seconds that have elapsed since Jan 01, 1970.
```

### Assumptions (Router)

There are not details about the URL routes necessary to expose the API endpoints. I went ahead and chose the ones that make more sense considering the data that is being handled and the way this data is being managed.

```
GET  /             - Entry point.
GET  /admin        - Administration interface.
GET  /reviews/list - JSON-encoded object with the reviews.
POST /reviews/save - Inserts a new review into the database.
```

### Assumptions (HTTP API - )

The details about the HTTP API are scarce, there is some information in the document but there are some things missing, for example, what is the format of the product identifier? I assumed it is a SKU (alpha-numeric code) and so I went ahead and hardcoded one in the Vue application for simplicity.

`GET /reviews/list` expects an `uid=[SKU]` parameter where `SKU` is the unique identifier of the product that is going to be reviewed. The API will retrieve all reviews associated to this product ID that were already approved by an administrator, these reviews will be ordered in descending way according to the time they were created _(latest review will be on top)_.

`POST /reviews/save` expects the following HTTP request:

```
Content-Type: application/x-www-form-urlencoded

rating  - Number from zero to ten (0-10)
          Numbers lower   than  0 turn into  0
          Numbers greater than 10 turn into 10
uid     - Hopefully a SKU [0-9A-Z] (but flexible)
name    - Arbitrary string representing the name of the reviewer.
email   - Arbitrary e-mail address identifying the reviewer.
comment - Arbitrary text with the product review.
```

### Caveats (Vue Router)

At first I wanted to have access to the parameters that are coming via GET to the application, but later I realized that I don't have enough information about the other project to know how the UID is going to be sent to this interface. Because of this, I have decided to get rid of "Vue Router" and keep the code lightweight, the URL query accessor has been implemented using vanilla JavaScript.

### Caveats (Hardcoded UID)

To simplify the development and make the demo usable, I have decided to hardcode the ID of the product that is being reviewed, and with this all the information associated to such product including the featured image, title, and description.

### Caveats (Leaking Emails)

In the Go source code you will notice that there are two custom structures, one to hold all the data associated to each review and another one to hold the data that can be shared to the public via GET. If you are thinking that these two structures could be merged, you are wrong.

Because I decided to display the avatar of the reviewer using [Gravatar.com](https://en.gravatar.com/) I wanted to send the email of each user from the database to the interface, however, because the API is public this means anyone can query the database looking for email addresses to spam. To prevent this, I have decided to generate the hash that Gravatar needs in the server and send the URL to the avatar completely built.

**WARNING!!** One could still obtain the original email address brute-forcing the hash which is simply a MD5 _(don't ask me why, that's what Gravatar uses)_ but the point of this challenge is to demonstrate my Go skills and not my cryptographic skills.
