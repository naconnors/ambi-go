# Ambi

Ambi is an Elixir-based web service that presents a Phoenix LiveView frontend to display real time ambient room conditions
like temperature, humidity, pressure, air quality, dust concentration, etc. It uses the Bulma CSS framework
for some attractive base UI components and Phoenix LiveView to push updates to the client with no page
refresh needed.

What Ambi's web-based UI looks like as of April, 2022

<img width="1482" alt="Screen Shot 2022-04-23 at 19 23 51" src="https://user-images.githubusercontent.com/3219120/164950502-6e58e0ed-5f58-4018-ae67-908960bb94a3.png">

## Set Up Git Hooks

The Ambi repository makes use of several Git hooks to ensure that code quality standards are met and consistent. To automatically configure these hooks for your local workspace, you can run the following:
```bash
./scripts/create-git-hooks
```

This will create symlinks to the Git hooks, preserving any hooks that you may have already configured.

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

## Install Node

Make sure to install [Node.js](https://nodejs.org/en/download/package-manager/), version 14 is the most tested version at the time of writing.
If you have multiple versions of Node installed, make sure that version 14 is run when running `node -v` from a command line.

For example, at the time of writing, for Ubuntu 22.04:
```
$ node -v
v14.19.1
```

## Starting the Ambi backend (Elixir)

To start your Phoenix server:

  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.setup`
  * Install Node.js dependencies with `npm install` inside the `assets` directory
  * Start Phoenix endpoint with `mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.

## Running as a systemd service on Linux (Ubuntu)

To automatically run your Ambi backend as a service that starts up when your Linux system boots, create the file `/etc/systemd/system/ambi.service`:

```
[Unit]
Description=Ambi Backend
After=display-manager.service
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=on-failure
RestartSec=5
User=jhodapp
WorkingDirectory=/home/<user>/Projects/ambi
ExecStart=/usr/bin/mix phx.server

[Install]
WantedBy=multi-user.target
```

Note: replace `<user>` with your username that you've cloned the Ambi repo as

Next, enable your new service to start on bootup: `sudo systemctl enable ambi`
Finally, start your new service: `sudo systemctl start ambi`

Note: if it does not start for you, check to make sure you have a valid service to start after (i.e. `After=display-manager.service`. Check in `/etc/systemd/system/` for available services that your system has.

## Running in a Docker container

To run ambi in a Docker container along with another one for Postgresql 11:

 * Build the web/DB containers: `docker-compose build`
 * Create the DB in the DB container: `docker-compose run web mix ecto.create`
 * Run the Ecto DB migrations inside the web container: `docker-compose run web mix ecto.migrate`
 * Run the application in the container: `docker-compose up`

 Note: this basic Docker setup was done following this [guide](https://dev.to/hlappa/development-environment-for-elixir-phoenix-with-docker-and-docker-compose-2g17)

Ready to run in production? Please [check our deployment guides](https://hexdocs.pm/phoenix/deployment.html).

## Learn more

  * Official website: https://www.phoenixframework.org/
  * Guides: https://hexdocs.pm/phoenix/overview.html
  * Docs: https://hexdocs.pm/phoenix
  * Forum: https://elixirforum.com/c/phoenix-forum
  * Source: https://github.com/phoenixframework/phoenix

## License

This project is licensed under the [BSD + Patent license](https://opensource.org/licenses/BSDplusPatent).

Any submissions to this project (e.g. as Pull Requests) must be made available under these terms.