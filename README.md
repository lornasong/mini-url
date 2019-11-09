# mini-url
A play on tiny url

### To start running
1. Set up a sql database for the application to connect with a table url with at least the following fields:
```
 +-------------+--------------+------+-----+
 | Field       | Type         | Null | Key |
 +-------------+--------------+------+-----+
 | id          | varchar(10)  | NO   | PRI |
 +-------------+--------------+------+-----+
 | originalUrl | text         | NO   |     |
 +-------------+--------------+------+-----+
```

2. Create a file for the following required environmental variables:
  * PORT - the port to serve the application
  * BASE_URL - the base url of your application e.g. http://localhost:8000
  * DATABASE_URL - the url of the database created in the first step

```
# Example
export PORT=8000
export BASE_URL=http://localhost:8000
export DATABASE_URL=postgres://<redacted>
```

3. Download dependencies. If go modules are on, run in command line `make deps`
4. Run in command line `make run`
5. In browser go to `http://localhost:<PORT>/minihome`
