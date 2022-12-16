# Ambi

Ambi is an Golang-based web service that presents a frontend to display real time ambient room conditions
like temperature, humidity, pressure, air quality, dust concentration, etc. It uses the Bulma CSS framework
for some attractive base UI components.
## Install Postgresql

For macOS the [Postgreql.app](https://postgresapp.com/) is the easiest and best option for your local development machine. Version 12
is the most tested version of the DB at the time of writing.

For Linux, use the system package manager to install Postgresql@14

Keep the local db instance running on port 5432 and make sure to set up a postgres user/password:

`$ sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"`

To allow Ambi to access your Postgresql cluster via localhost, make sure your `/etc/postgresql/14/main/pg_hba.conf` looks like the following:

```
# Database administrative login by Unix domain socket
local   all             postgres                                peer

# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     peer
host    all             postgres        localhost               trust
# IPv4 local connections:
host    all             all             127.0.0.1/32            scram-sha-256
# IPv6 local connections:
host    all             all             ::1/128                 scram-sha-256
# Allow replication connections from localhost, by a user with the
# replication privilege.
local   replication     all                                     peer
host    replication     all             127.0.0.1/32            scram-sha-256
host    replication     all             ::1/128                 scram-sha-256
```
## Starting Ambi

To start your server:
  * Create database `ambi_go_dev`
  * Run `make migrate-up` to create required tables
  * Start server with `make run`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.
## License

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).

Any submissions to this project (e.g. as Pull Requests) must be made available under these terms.