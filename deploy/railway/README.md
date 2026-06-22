# Railway Deployment

This repository is ready to deploy on Railway with the root `Dockerfile` and
`railway.json`.

## Required Railway setup

1. Create a Railway project from this repository.
2. Railway will build with the root `Dockerfile`.
3. Add a Volume and mount it at `/etc/x-ui`.
   This keeps the default SQLite database (`/etc/x-ui/x-ui.db`) across
   redeploys.
4. Generate a public domain for the service.

The app reads `XUI_PORT`, while Railway injects `PORT`. The Docker entrypoint
maps `PORT` to `XUI_PORT` automatically when `XUI_PORT` is not set.

## Optional variables

Set these in Railway only when needed:

```env
XUI_INIT_WEB_BASE_PATH=/
XUI_LOG_LEVEL=info
XUI_DB_TYPE=sqlite
XUI_DB_FOLDER=/etc/x-ui
XUI_LOG_FOLDER=/etc/x-ui
```

Do not set `XUI_DB_FOLDER=x-ui` on Railway when using the `/etc/x-ui` volume;
that stores SQLite data in the app working directory instead of the persistent
volume.

To use Railway Postgres instead of SQLite, add a Railway Postgres service and
set:

```env
XUI_DB_TYPE=postgres
XUI_DB_DSN=${{Postgres.DATABASE_URL}}
```

## Railway limitations

Railway containers do not provide the `NET_ADMIN`/iptables capability used by
the bundled Fail2ban integration, so the entrypoint disables Fail2ban
automatically when it detects Railway.

The web panel can be exposed through Railway's public HTTP domain. Xray inbound
ports and UDP/TCP proxy traffic may require Railway TCP proxy support or another
host/VPS, depending on the protocol and port layout you configure.
