module github.com/libp2p/dht-utils

go 1.12

require (
	github.com/gogo/protobuf v1.2.1
	github.com/ipfs/go-cid v0.0.2
	github.com/ipfs/go-datastore v0.0.5
	github.com/ipfs/go-ipfs v0.4.21
	github.com/ipfs/go-ipfs-util v0.0.1
	github.com/ipfs/go-log v0.0.1
	github.com/libp2p/go-libp2p v0.2.0-0.20190628095754-ccf9943938b9
	github.com/libp2p/go-libp2p-core v0.0.7-0.20190626134135-aca080dccfc2
	github.com/libp2p/go-libp2p-crypto v0.1.0
	github.com/libp2p/go-libp2p-host v0.1.0
	github.com/libp2p/go-libp2p-kad-dht v0.1.2-0.20190628100158-d8d74a239cb8
	github.com/libp2p/go-libp2p-kbucket v0.2.0
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.1.2-0.20190628102222-693780b745ad
	github.com/libp2p/go-libp2p-record v0.1.0
	github.com/libp2p/go-libp2p-routing v0.1.0
	github.com/libp2p/go-libp2p-transport v0.0.5 // indirect
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/urfave/cli v1.20.0
	google.golang.org/appengine v1.4.0 // indirect
)

replace github.com/libp2p/go-libp2p-kad-dht => github.com/aschmahmann/go-libp2p-kad-dht v0.1.1-0.20190628175927-c7e75e511982

replace github.com/dgraph-io/badger => github.com/dgraph-io/badger v1.6.0-rc1
