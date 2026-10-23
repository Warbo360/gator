# Gator (The Aggregator!)

## Intro

- CLI command utilizing a PostgreSQL database to add, follow, scrape, and save 
  RSS feeds! Built following the lovely guided project of boot.dev! Support for 
  multiple users and persistent state due to the use of the database.

## Requirements

- Install PostgreSQL, Go, and Goose on your system:
- Additionally SQLC is used for creating Go code from SQL query files. To run 
  the program as is, this is not required but if you'd like to make alterations 
  to the program with your own queries, it is a useful CLI tool.
  - For installing PostgreSQL on Linux/WSL (Debian):

```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

- Verify installation, should be version 15+:

```
psql --version
```

- Update the postgres password to password of preference:

```
sudo passwd postgres
```

- Start the postgres server in the background:

```
sudo service postgresql start
```

- Enter the `psql` shell:

```
sudo -u psql
```

- Create a new database, you can name it whatever:

```
CREATE DATABASE gator;
```

- Connect to the new database:

```
\c gator
```

- Set the user password (Linux/WSL only), note this is the database password, 
  where as above it was the system password:

```
ALTER USER postgres PASSWORD 'postgres';
```

- Query the database:

```
SELECT version();
```

- Exit the database:

```
exit
```

- Goose is a Go CLI migration tool that makes up and down migrations relatively 
  easy and reproducible, install as follows:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Run `goose -version` to confirm installation

- Next you'll need to get your connection string which has the following format:

```
'protocol://username:password@host:port/database'
```

- If following the instructions as is up to this point your connection string 
  will probably be:

```
'postgres://postgres:postgres@localhost:5432/gator'
```

- You can test if it is by running:

```
psql 'postgres://postgres:postgres@localhost:5432/gator'
```

- If the connection works save the connection string for usage with goose like 
  so in the `sql/schema` directory in `gator/`:

```
goose postgres 'postgres://postgres:postgres@localhost:5432/gator' up
```

- This will create the needed tables in the database for the program to execute 
  properly

- You can also add the connection string to the `.gatorconfig.json` file seen in 
  the installation setup below for the `"db_url":` field with the addition of to 
  the:

  ```
  `postgres://postgres:postgres@localhost:5432/gator?sslmode=disable`
  ```

- Optionally for SQLC:

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

- Then you'll see in the root of `gator` a `sqlc.yaml` file with the following:

```
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

- This just lets SQLC know that your schema files live in `sql/schema` and your 
  queries in `sql/queries`. `internal/database` is where the Go generated code 
  will be populated from your queries written in `sql/queries`. To use SQLC to 
  generate your own Go code from queries just run `sqlc generate` from the root 
  of `gator/`.

## Installation

1. Once both PostgreSQL and Go are installed on your system you can either:
  a. Remote install via `go install github.com/Warbo360/gator`
  b. Build from source:
    - `git clone git@github.com:Warbo360/gator.git` in your favorite place
    - `cd gator`
    - `go install .`

2. Create the config file is your home directory called `.gatorconfig.json` and 
copy the following snippet into it, this will hold who is currently logged in as 
well as the connection credentials for the database. You can replace the example 
insert for `"db_url"` field with your connection string that you established in 
the requirements section when configuring `goose`.

```
{
  "db_url": "postgres://example",
}
```

## Usage

- General format of using `gator`:

```
./gator <command> <args>
```

### Commands

#### Register

- Registers a new user, takes a single argument in the form of a username which 
  are unique to the database, will login the newly registered user:

```
./gator register username
```

#### Login

- Will login the first passed user to the command if they exist in the database:

```
./gator login username
```

#### Users

- Will list all registered users from in the database and indicate the currently 
  logged in user:

```
./gator users
```

#### Addfeed

- Takes two string arguments, a feed name, and a URL of the feed, adds the feed 
  to the feeds table in the database and as added by the currently logged in 
  user:

```
./gator addfeed "feed name" "url"
```

#### Feeds

- Lists all feeds added by all users in the database, lists their feed name, 
  URL, and what user added it:

```
./gator feeds
```

#### Follow

- Follows a feed that was added by another user to the currently logged in user, 
  adding a feed automatically follows it as well so following a feed added by 
  the current user does nothing:

```
./gator follow "url"
```

#### Unfollow

- Unfollows a currently followed feed by its passed URL:

```
./gator unfollow "url"
```

#### Following

- Lists all feeds currently being followed by the current user

```
./gator following
```

#### Agg

- Aggregrates all followed feeds by current user into posts that are saved to 
  the database, takes a single argument in the form of a time interval such as 
  `5s`. Will scrape the oldest last fetched feed every time interval passed:

```
./gator agg "time inverval"
```

#### Browse

- Will print the most recent added posts from the `agg` command for a user. 
  Takes a single argument in the form of a limit of number of posts to print, 
  will default to 2 posts if not specified:

```
./gator browse limit
```

#### Reset

- Mostly intended for end-to-end testing of the database, will `goose` migration 
  down and delete all tables in the database, not intended for normal operation, 
  use with caution:

```
./gator reset
```

## Feedback

- If you care to leave feedback on any of my implementation details or even the 
  format of this README file all feed back is welcome as I am a hobbyist at best 
  right now :)
