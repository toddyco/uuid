# uuid
Wrapper around Google's `uuid.UUID` type

## Basic Info

In this package is a struct called `UUID`, which has two fields:
1. `UUID`: A [Google uuid.UUID](https://github.com/google/uuid) (UUIDv4, specifically)
2. `Str`: A string form of the Google UUID

I created this because I need to be able to see the string representation of a UUID when 
running a Delve debug session. The Google `type UUID [16]byte` does not get picked up by
Delve in a way where its `String()` function gets called, and instead the debugger just
shows human-unreadable byte code for a UUID when hitting a breakpoint.

That's pretty much it.

To make this _somewhat_ a drop-in replacement for existing usages of Google's `uuid.UUID`,
my `UUID` struct has these **receiver** functions:
* `String() string`
* `IsNil() bool`
* `URN() string`
* `Scan(src any) error`
* `Value() (driver.Value, error)`
* `MarshalText() ([]byte, error)`
* `UnmarshalText(data []byte) error`
* `MarshalBinary() ([]byte, error)`
* `UnmarshalBinary(data []byte) error`
* `MarshalJSON() ([]byte, error)`
* `UnmarshalJSON(data []byte) error`

These **non-receiver** helper functions are also provided:
* `Parse(s string) (UUID, error)`
* `New() UUID`

I'll keep the release tag versions here in line with the ones in Google's
project (which implies I'll do my best to keep this updated each time the
Google repo has a new release).

------

## Tag Versions

The release tag is generated each time based on this format (standard Go pseudo version format):

```
VERSION={Google repo release number}
echo "v$VERSION-$(git show -s --format=%cd --date=format:%Y%m%d%H%M%S)-$(git rev-parse --short=12 HEAD)"
```

...where `{Google repo release number}` is replaced with a SemVer value.

For example:

```
VERSION=1.6.0
echo "v$VERSION-$(git show -s --format=%cd --date=format:%Y%m%d%H%M%S)-$(git rev-parse --short=12 HEAD)"
```
