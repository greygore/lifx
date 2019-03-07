/*
Package lan implements the LIFX LAN Protocol documented on the LIFX dev site:

https://lan.developer.lifx.com/

This is a minimal package meant only to create message packets and to send
them over the network. It is meant to be used by another library to create
a more developer friendly API.

The bulk of the code to support this package is generated automatically from
a machine readable representation of the API as described here:

https://github.com/LIFX/public-protocol

To regenerate these files, `go generate` will use the `genlifx` command to
pull the YAML file and process it. Because the documentation is not included,
you will need to use the API documentation to determine what messages do,
which fields are required, and what the values mean.
*/
package lan

//go:generate genlifx
