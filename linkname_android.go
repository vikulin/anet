//go:build android

package anet

import _ "unsafe"

//go:linkname zoneCache net.zoneCache
var zoneCache ipv6ZoneCache

//go:linkname zoneCacheX golang.org/x/net/internal/socket.zoneCache
var zoneCacheX ipv6ZoneCache
