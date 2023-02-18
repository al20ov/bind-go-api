[![Go](https://github.com/al20ov/bind-go-api/actions/workflows/go.yml/badge.svg)](https://github.com/al20ov/bind-go-api/actions/workflows/go.yml)

# bind-go-api

## Who is this project for

If you're running bind9 at home or in a small office and you're tired of having
to ssh into your DNS server and editing zone files by hand, this project is for
you.

## Example setup

### BIND9 config

One of the key features of this particular API over other BIND9 REST APIs is the
ability to list all zones on the server, but for that, you need to configure
your BIND server accordingly.

Here's what your `named.conf` should look like:
```conf
options {
    [...]
}

statistics-channels {
	inet 0.0.0.0 port 8080 allow { 192.168.1.0/24; };
};

zone "example.com" IN {
	type master;
	file "/etc/bind/com.example.zone";
	allow-update {
		key "axfr.";
	};
	zone-statistics yes;
};

[...]
```

The `statistics-channels` block as well as setting `zone-statistics` to "yes" in
all of the zones you want to be able to list are what's going to make this work.

Also make sure you generate a HMAC-SHA256 TSIG key using
`tsig-keygen <name of the key> > <name of the key>.key`
and copy/include it inside `named.conf`.

### bind-go-api config

All you have to do now is customize `config.json` to your own setup, and copy
the tsig key name and base64 encoded secret in that same file.

Use `config_example.json` as an example.
