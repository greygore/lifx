![License-MIT](https://img.shields.io/github/license/greygore/lifx.svg?style=flat)

# LIFX

## Overview

This is my attempt to create a complete Go library for [LIFX](http://www.lifx.com/) using the [LIFX LAN Protocol](https://lan.developer.lifx.com/). For now, there's a minimal implementation of the actual protocol in the `lan` package. Because LIFX was kind enough to provide [machine readable documentation](https://github.com/LIFX/public-protocol), most of the code is generated and is therefore complete (as of the last commit), and can be easily updated as changes are made to the protocol.

## To Do

Future work will be done to wrap this protocol in a more developer friendly API.

## License

[MIT License](https://github.com/greygore/lifx/blob/master/LICENSE).