# shjp-server

shjp-server is the server program written in Go langugage to serve requests from shjp applications.
It exposes the GraphQL endpoint as outlined in the API section below.
For the database PostgreSQL is used.

## Installation

### Mac & Linux

1. Make sure you have Go language and Docker installed
2. Place this project directory inside the $GO_ROOT. If your $GO_ROOT is ~/go, it should be at ~/go/src/github.com/shjp/shjp-server
3. Fetch the postgres Docker image:
```
docker pull postgres
```
4. Install goose:
```
go get -u github.com/pressly/goose/cmd/goose
```
5. Start docker containers for the database:
```
make db_init
```
6. Migrate the database schema:
```
make db_up_dev
```
7. (Optional) If you want to import fixture data:
```
make db_fixtures
```

### Windows

1. Download TDM-GCC (https://sourceforge.net/projects/tdm-gcc/?source=typ_redirect)
2. Set the environment variable to use TDD-GCC:
  * Right-click on your "My Computer" icon and select "Properties"
  * Click on the "Advanced system settings", select "Advanced" tab, then "Environment Variables" button
  * You will see two boxes. The first box should say "User variables for <USERNAME>". Find "PATH" entry in the first box and click "Edit" button. DO NOT touch the bottom box. Click "New" and enter the bin path of the downloaded TDM-GCC. The default location would be C:\TDD-GCC-64\bin
3. Download Docker Toolbox and install it (https://docs.docker.com/toobox/overview/#ready-to-get-started)
4. Fetch the postgres docker image:
```
docker pull postgres
```
5. Install goose:
```
go get -u github.com/pressly/goose/cmd/goose
```
6. Start docker containers for the database:
```
mingw32-make.exe db_init
```
7. Migrate the database schema:
```
mingw32-make.exe db_up_dev_win
```
8. (Optional) If you want to import fixture data:
```
mingw32-make.exe db_fixtures_win
```

## Run the server

```
make run
```
On windows,
```
make run_win
```

Type localhost:8080/graphql into the browser address bar and verify that you can see the GraphQL interface.

## API

The best way to explore the API is through the GraphQL interface as explained in the previous section. For example, if you send a query like so
```
query {
  groups {
    id, name
  }
}
```
Then you will get back something like
```javascript
{
  "data": {
    "groups": [
      {
        "description": "Sing Sang Sung",
        "name": "Choir"
      },
      {
        "description": "booooo",
        "name": "Jollaboo"
      },
      ...
    ]
  }
}
```