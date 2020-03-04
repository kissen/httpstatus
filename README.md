httpstatus
==========

This is a tiny Go package that provides nice human-readable descriptions of typical HTTP
errors codes. You may want to display them to a user instead of just the plain error code.

`httpstatus` is different from [`http.StatusText`](https://golang.org/pkg/net/http/#StatusText)
in that the provided descriptions are more elaborate and detailed.

Usage
-----

Call

	httpstatus.Describe(404)

to get the description. In this case, we get

	The requested URL was not found on the server. If you entered
	the URL manually please check your spelling and try again.

which is a lot nicer than just plain

	Not Found

Credit
------

The descriptions are taken from the [pallets/werkzeug](https://github.com/pallets/werkzeug)
project which is licensed under a BSD 3-Clause license, just like this package. See
`LICENSE.rst` for the specific terms.
