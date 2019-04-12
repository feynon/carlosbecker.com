---
title: "Prometheus authentication with oauth2_proxy"
date: 2018-05-28T21:24:04-03:00
city: Joinville
tags:
- prometheus
- monitoring
---

I wanted to set up a [prometheus][] machine for me to monitor random stuff,
but I was always postpone that because I didn't want to use SSH
port-forwarding, firewalls, create a VPC and/or setup an OpenVPN server or
anything like that.

<!--more-->

I just wanted something simple to maybe authenticate with github and go on.

Looking into some random GitLab wiki (I don't remember which one specifically),
I found about [oauth2_proxy][], and it seemed like a good idea.

Since this is a simple setup, I used [docker-compose][] and [rsync][] to set
up all the environment.

# Idea

The general idea is quite simple:

- all things but nginx listen on `127.0.0.1` only;
- nginx listens on `80` and `proxy_forward`s to [oauth2_proxy][] and the
  other services:
  - `/` forwards to [prometheus][];
  - `/grafana` forwards to [grafana][];
  - `/alertmanager` forwards to [alertmanager][];
  - all of the above authenticate using `proxy_forward` and [nginx][]'s
    `auth_request` directive.

So, let's get this thing started!

# Prometheus

First one is [prometheus][] itself:

```yaml
# docker-compose.yml
version: '3'
services:
  prometheus:
    image: prom/prometheus
    ports:
    - "127.0.0.1:9090:9090"
    command:
    - '--web.external-url=http://m.carlosbecker.com/'
    # content hidden for the sake of brevity
```

The important parts are:

- listen on `127.0.0.1:9090` - so it won't be exposed to the world;
- the `web.external-url` with the final URL: this is required for external
  links to work properly.

# AlertManager

We will also need [alertmanager][] to route the alerts, right? Let's do it:

```yaml
# docker-compose.yml
version: '3'
services:
  # content hidden for the sake of brevity
  alertmanager:
    image: prom/alertmanager
    ports:
    - "127.0.0.1:9093:9093"
    command:
    - '--web.external-url=http://m.carlosbecker.com/alertmanager/'
    # content hidden for the sake of brevity
```

The important parts are:

- listen on `127.0.0.1:9093` - so it won't be exposed to the world;
- the `web.external-url` with the final URL: this is required for external
  links to work properly.

We also need to add [alertmanager][] to `prometheus.yml`:

```yaml
# prometheus.yml
alerting:
  alertmanagers:
  - path_prefix: /alertmanager/
    static_configs:
    - targets:
      - alertmanager:9093
scrape_configs:
  - job_name: 'alertmanager'
    metrics_path: /alertmanager/metrics
    static_configs:
    - targets: ['alertmanager:9093']
```

Note that we change the `path_prefix` in the `alerting` section and also the
`metrics_path` in the `scrape_configs` section.

# Grafana

Everyone likes dashboards, and [grafana][] is DOPE for doing that. Let's
add it as well:

```yaml
# docker-compose.yml
version: '3'
services:
  # content hidden for the sake of brevity
  grafana:
    image: grafana/grafana:5.1.3
    restart: always
    user: "0"
    volumes:
    - /local/path/to/grafana.ini:/etc/grafana/grafana.ini:ro
    ports:
    - "127.0.0.1:3000:3000"
    # content hidden for the sake of brevity
```

Here things get a little more tricky: we also have a `granafa.ini` config
file. Here are its contents:

```ini
; grafana.ini
[server]
protocol = http
domain = m.carlosbecker.com
root_url = %(protocol)s://%(domain)s/grafana/

[users]
allow_sign_up = false
auto_assign_org = true
auto_assign_org_role = Admin

[auth.proxy]
enabled = true
header_name = X-Email
header_property = email
auto_sign_up = true
```

We have a couple of important things to look at here. On the
`docker-compose.yml` file:

- listen on `127.0.0.1:3000` - so it won't be exposed to the world;
- mount the `grafana.ini` config file;

And in the `grafana.ini` file:

#### In the `server` section:

The `root_url` defines the `/grafana/` suffix in the root.

This is needed because otherwise, even with `proxy_pass` on [nginx][],
[grafana][] keeps trying to redirect to `/`, as mentioned on the beggining,
[prometheus][] will leave on `/`. This config file fixes that;

#### In the `users` and `auth.proxy` sections:

Those sections tell [grafana][] to auto-create an user when someone authenticates
through the proxy (using the `X-Email` header) and to give this user the
`Admin` role and assing to the organization. It also disables the sign up form.

This way we don't need to login on both [oauth2_proxy][] **and** [grafana][]:
You log in within [oauth2_proxy][] and everything just works!

# oauth2_proxy

Finally, let's require authentication using [oauth2_proxy][]. I'm using
GitHub for that, so, the first thing I did was to
[create a new GitHub app][create-app]. You can find the details on
the [oauth2_proxy README][oauth2_proxy_readme].

Once I had the client ID and secret, it was pretty straightforward:

```yaml
# docker-compose.yml
version: '3'
services:
  # content hidden for the sake of brevity
  proxy:
    image: caarlos0/oauth2_proxy
    ports:
    - 127.0.0.1:4180:4180
    command:
    - '-client-id=123'
    - '-client-secret=456'
    - '-provider=github'
    - '-github-org=caarlos0-m'
    - '-email-domain=*'
    - '-cookie-secret=foo bar 1234'
    - '-cookie-secure=false'
    - '-upstream=http://nginx:80'
    - '-http-address=0.0.0.0:4180'
    - '-redirect-url=http://m.carlosbecker.com/oauth2/callback'
    - '-set-xauthrequest=true'
    # content hidden for the sake of brevity
```

The important things here are:

- listen on `127.0.0.1:4180` - so it won't be exposed to the world;
- `upstream` is set to the [nginx][] container;
- `http-address` is set to listen on `0.0.0.0` so we can expose the service
  to the host (`oauth2_proxy` listens on `127.0.0.1` by default);
- `redirect-url` must be the same as the one informed while creating the GitHub
  app;
- `client-id`, `client-secret` and `provider` are the GitHub oauth2 settings;
- `github-org` is the required org an user needs to be member of to be allowed
in;
- `email-domain` could be an additional email domain filter - for me the org
filter is enough;
- `set-xauthrequest` is set to true se we can pass through the user and
email headers - grafana uses this header to auto-create an user and log it in;
- `cookie-secure` is set ot false due the lack of https. I'll manage to add
  [let's encrypt][letsencrypt] anoother and will create a new post.

# nginx

**Finally**, the last part: [nginx][]!

First, let's create a nginx config file:

```nginx
# nginx.conf
server {
  listen 80;
  server_name m.carlosbecker.com;

  location /oauth2/ {
    proxy_pass       http://proxy:4180;
    proxy_set_header Host                    $host;
    proxy_set_header X-Real-IP               $remote_addr;
    proxy_set_header X-Scheme                $scheme;
    proxy_set_header X-Auth-Request-Redirect $request_uri;
  }
  location = /oauth2/auth {
    proxy_pass       http://proxy:4180;
    proxy_set_header Host             $host;
    proxy_set_header X-Real-IP        $remote_addr;
    proxy_set_header X-Scheme         $scheme;
    proxy_set_header Content-Length   "";
    proxy_pass_request_body           off;
  }

  location /alertmanager/ {
    auth_request /oauth2/auth;
    error_page 401 = /oauth2/sign_in;

    auth_request_set $user   $upstream_http_x_auth_request_user;
    auth_request_set $email  $upstream_http_x_auth_request_email;
    proxy_set_header X-User  $user;
    proxy_set_header X-Email $email;

    auth_request_set $auth_cookie $upstream_http_set_cookie;
    add_header Set-Cookie $auth_cookie;

    proxy_pass http://alertmanager:9093/alertmanager/;
  }

  location /grafana/ {
    auth_request /oauth2/auth;
    error_page 401 = /oauth2/sign_in;

    auth_request_set $user   $upstream_http_x_auth_request_user;
    auth_request_set $email  $upstream_http_x_auth_request_email;
    proxy_set_header X-User  $user;
    proxy_set_header X-Email $email;

    auth_request_set $auth_cookie $upstream_http_set_cookie;
    add_header Set-Cookie $auth_cookie;

    proxy_pass http://grafana:3000/;
  }

  location / {
    auth_request /oauth2/auth;
    error_page 401 = /oauth2/sign_in;

    auth_request_set $user   $upstream_http_x_auth_request_user;
    auth_request_set $email  $upstream_http_x_auth_request_email;
    proxy_set_header X-User  $user;
    proxy_set_header X-Email $email;

    auth_request_set $auth_cookie $upstream_http_set_cookie;
    add_header Set-Cookie $auth_cookie;

    proxy_pass http://prometheus:9090/;
  }
}
```

Most of it is based on the [oauth2_proxy README examples][oauth2_proxy_readme_nginx].

Now, let's add it to the `docker-compose.yml` file:

```yaml
# docker-compose.yml
version: '3'
services:
  # content hidden for the sake of brevity
  nginx:
    image: nginx
    ports:
    - 80:80
    volumes:
    - /local/path/to/nginx.conf:/etc/nginx/conf.d/m.conf:ro
    # content hidden for the sake of brevity
```

Yeah, this one is pretty simple! Now we finally expose one service - on the
port `80`, and we mount our config file to the `/etc/nginx/conf.d` folder.
[nginx][] will be the only thing facing the internet for real, and
it will route traffic to the right places.

# Closing

So that's it! I'm running this on http://m.carlosbecker.com - which, given
that everything is working as expected, you won't be able to access, and it
is working great.

<!-- You can see **all** the code (including all config files, the exporters I'm
using, the `Makefile` I use to sync things up and etc) on
[this github repository][repo]. -->

I hope this is usefull to you somehow and feel free to ask question in the
comments box down bellow.

Cheers!

[oauth2_proxy_readme_nginx]: https://github.com/bitly/oauth2_proxy#configuring-for-use-with-the-nginx-auth_request-directive
[create-app]: https://github.com/settings/developers
[oauth2_proxy_readme]: https://github.com/bitly/oauth2_proxy#github-auth-provider
[nginx]: https://www.nginx.com/
[letsencrypt]: https://letsencrypt.org/
[prometheus]: https://prometheus.io
[alertmanager]: https://prometheus.io/docs/alerting/alertmanager/
[grafana]: https://grafana.com/
[oauth2_proxy]: https://github.com/bitly/oauth2_proxy
<!-- [repo]: https://github.com/caarlos0/m.carlosbecker.com -->
[docker-compose]: https://docs.docker.com/compose/
[rsync]: https://linux.die.net/man/1/rsync
