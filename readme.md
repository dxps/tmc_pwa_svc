# TM Community - PWA edition

This is version 2 of TM Community solution, written as a Progressive Web App (PWA) using `go-app` framework.

<br/>

## Usage

### Prerequisites

-   [entr](https://eradman.com/entrproject/) - used for triggering the recompile and restart of the app on code changes.

-   [node](https://nodejs.org/en/download/) - used for running the TailwindCSS compiler using `npm`.<br>
    It regenerates the `web/styles/main.css` file on detected TailwindCSS related rules.

-   [migrate](https://github.com/golang-migrate/migrate) - used for running database migrations.<br/>
    See [instalation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) instructions for details.

### Start

1. Run the PostgreSQL server using `cd ops && ./run_db.sh`.
1. Run the TailwindCSS compiler using `./run_css.sh`.
1. To just start the app, use `make dev`.

Besides the first step that starts the database server, the other steps are for running the app in "development mode" meaning that it gets restarted on code changes.

Go to http://localhost:8000 to access the Web UI.

<br/>
