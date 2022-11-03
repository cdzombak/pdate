# `dateutil`

`dateutil` attempts to parse a given date/time stamp string, trying different formats. Once it parses the timestamp, it presents the parsed timestamp in UTC and your local timezone.

## Usage Examples

### Parse a specific date/time stamp

#### RFC 3339

```
$ dateutil 2022-10-06T17:49:19Z
 input:	2022-10-06T17:49:19Z
parsed:	2022-10-06 17:49:19 UTC
       	(verify this matches your input)

   UTC:	2022-10-06 5:49:19 PM
       	2022-10-06T17:49:19Z

 local:	2022-10-06 1:49:19 PM EDT
	(28 days ago)
```

[(sample screenshot)](https://github.com/cdzombak/dateutil/blob/main/screenshots/dateutil%20-%20rfc3339%20timestamp.png)

#### Unix timestamp

```
$ dateutil 1665001628419
 input:	1665001628419
parsed:	2022-10-05 16:27:08 EDT
       	(verify this matches your input)

   UTC:	2022-10-05 8:27:08 PM
       	2022-10-05T20:27:08Z

 local:	2022-10-05 4:27:08 PM EDT
	(29 days ago)
```

[(sample screenshot)](https://github.com/cdzombak/dateutil/blob/main/screenshots/dateutil%20-%20unix%20timestamp.png)

### Get the current date/time in UTC

```
$ dateutil
 input:	now

   UTC:	2022-11-03 3:08:53 PM
       	2022-11-03T15:08:53Z

 local:	2022-11-03 11:08:53 AM EDT
```

## Installation

Installation currently requires building from source; a working Go installation is required.

```
git clone https://github.com/cdzombak/dateutil.git
cd dateutil
make install
```

## See Also

- [clock.dzdz.cz](https://clock.dzdz.cz) displays the current time in UTC and your local timezone. [You can also install it using a trivial Electron wrapper.](https://github.com/cdzombak/clock/tree/master/app)
