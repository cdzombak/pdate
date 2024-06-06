# `pdate`

`pdate` attempts to parse a given date/timestamp string, trying different formats. Once it parses the timestamp, it presents the parsed timestamp in UTC and your local timezone.

## Usage Examples

### Parse a specific date/time stamp

#### RFC 3339

```text
$ pdate 2022-10-06T17:49:19Z
 input:	2022-10-06T17:49:19Z
parsed:	2022-10-06 17:49:19 UTC
       	(verify this matches your input)

   UTC:	2022-10-06 5:49:19 PM
       	2022-10-06T17:49:19Z

 local:	2022-10-06 1:49:19 PM EDT
	(28 days ago)
```

[(sample screenshot)](https://github.com/cdzombak/pdate/blob/main/screenshots/dateutil%20-%20rfc3339%20timestamp.png)

#### Unix timestamp

```text
$ pdate 1665001628419
 input:	1665001628419
parsed:	2022-10-05 16:27:08 EDT
       	(verify this matches your input)

   UTC:	2022-10-05 8:27:08 PM
       	2022-10-05T20:27:08Z

 local:	2022-10-05 4:27:08 PM EDT
	(29 days ago)
```

[(sample screenshot)](https://github.com/cdzombak/pdate/blob/main/screenshots/dateutil%20-%20unix%20timestamp.png)

#### ULID

```text
$ pdate "01D78XZ44G0000000000000000"
 input: 01D78XZ44G0000000000000000
parsed: 2019-03-30 23:51:23 EDT
        (verify this matches your input/expectations)

   UTC: 2019-03-31 3:51:23 AM
        2019-03-31T03:51:23Z

 local: 2019-03-30 11:51:23 PM EDT
        (2019-03-30 23:51:23 EDT)
```

### Get the current date/time

```text
$ pdate
 input:	now

   UTC:	2022-11-03 3:08:53 PM
       	2022-11-03T15:08:53Z

 local:	2022-11-03 11:08:53 AM EDT
```

## Installation

### macOS via Homebrew

```shell
brew install cdzombak/oss/pdate
```

### Debian/Ubuntu and derivatives, via apt repository

Install my Debian repository if you haven't already:

```shell
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://dist.cdzombak.net/deb.key | sudo gpg --dearmor -o /etc/apt/keyrings/dist-cdzombak-net.gpg
sudo chmod 0644 /etc/apt/keyrings/dist-cdzombak-net.gpg
echo -e "deb [signed-by=/etc/apt/keyrings/dist-cdzombak-net.gpg] https://dist.cdzombak.net/deb/oss any oss\n" | sudo tee -a /etc/apt/sources.list.d/dist-cdzombak-net.list > /dev/null
sudo apt-get update
```

Then install `pdate` via `apt-get`:

```shell
sudo apt-get install pdate
```

### From source

A working Go installation is required.

```shell
git clone https://github.com/cdzombak/pdate.git
cd pdate
go build -ldflags="-X main.version=$(./.version.sh)" -o /usr/local/bin/pdate .
```

## See Also

- [clock.dzdz.cz](https://clock.dzdz.cz) displays the current time in UTC and your local timezone.

## License

Apache 2.0; see LICENSE in this repo.

## Author

Chris Dzombak
- [dzombak.com](https://www.dzombak.com)
- [github.com/cdzombak](https://github.com/cdzombak)
