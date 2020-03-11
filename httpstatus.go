// This package provides nice human-readable descriptions of typical HTTP
// errors codes. You may want to display them to a user instead of just the
// plain error code.
package httpstatus

var descriptions = map[int]string{
	400: "The browser (or proxy) sent a request that this server could not understand.",
	401: "The server could not verify that you are authorized to access" +
		" the URL requested. You either supplied the wrong credentials" +
		" (e.g. a bad password), or your browser doesn't understand" +
		" how to supply the credentials required.",
	403: "You don't have the permission to access the requested" +
		" resource. It is either read-protected or not readable by the" +
		" server.",
	404: "The requested URL was not found on the server. If you entered" +
		" the URL manually please check your spelling and try again.",
	405: "The method is not allowed for the requested URL.",
	406: "The resource identified by the request is only capable of" +
		" generating response entities which have content" +
		" characteristics not acceptable according to the accept" +
		" headers sent in the request.",
	408: "The server closed the network connection because the browser" +
		" didn't finish the request within the specified time.",
	409: "A conflict happened while processing the request. The" +
		" resource might have been modified while the request was being" +
		" processed.",
	410: "The requested URL is no longer available on this server and" +
		" there is no forwarding address. If you followed a link from a" +
		" foreign page, please contact the author of this page.",
	411: "A request with this method requires a valid <code>Content-" +
		"Length</code> header.",
	412: "The precondition on the request for the URL failed positive evaluation.",
	413: "The data value transmitted exceeds the capacity limit.",
	414: "The length of the requested URL exceeds the capacity limit for" +
		" this server. The request cannot be processed.",
	415: "The server does not support the media type transmitted in the request.",
	416: "The server cannot provide the requested range.",
	417: "The server could not meet the requirements of the Expect header",
	418: "This server is a teapot, not a coffee machine",
	422: "The request was well-formed but was unable to be followed due" +
		" to semantic errors.",
	423: "The resource that is being accessed is locked.",
	424: "The method could not be performed on the resource because the" +
		" requested action depended on another action and that action" +
		" failed.",
	428: "This request is required to be conditional; try using" +
		"\"If-Match\" or \"If-Unmodified-Since\".",
	429: "This user has exceeded an allotted request count. Try again later.",
	431: "One or more header fields exceeds the maximum size.",
	451: "Unavailable for legal reasons.",
	500: "The server encountered an internal error and was unable to" +
		" complete your request. Either the server is overloaded or" +
		" there is an error in the application.",
	501: "The server does not support the action requested by the browser.",
	502: "The proxy server received an invalid response from an upstream server.",
	503: "The server is temporarily unable to service your request due" +
		" to maintenance downtime or capacity problems. Please try" +
		" again later.",
	504: "The connection to an upstream server timed out.",
	505: "The server does not support the HTTP protocol version used in the request.",
}

// Given an HTTP status code, return a nice human-readable description of
// what this error is about and how it may be fixed. If no description is
// available for status, an empty string is returned instead.
//
// This function is different from http.StatusText in that the provided
// descriptions are more elaborate and detailed.
func Describe(status int) (description string) {
	return descriptions[status]
}
