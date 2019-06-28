package main

import (
	"context"
	"fmt"
	util "github.com/ipfs/go-ipfs-util"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/libp2p/go-libp2p-kbucket/keyspace"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	notif "github.com/libp2p/go-libp2p-routing/notifications"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
)

var bsaddrs = []string{
	"/ip4/3.121.100.130/tcp/20299/ipfs/12D3KooWByhRz8Quz98bVYWikKY4Wpx4sMCMfqYZQvQyruTSvgVp",
	"/ip4/3.121.100.130/tcp/20300/ipfs/12D3KooWSru22AzV79M5zxuKsfEx1d8S6Fd1m9AGH97a5aD4VZMn",
	"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	"/ip4/104.236.179.241/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",
	"/ip4/128.199.219.111/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",
	"/ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",
	"/ip4/178.62.158.247/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd",
}

func main() {

	var bspis []*peer.AddrInfo
	for _, a := range bsaddrs {
		maddr, err := ma.NewMultiaddr(a)
		if err != nil {
			panic(err)
		}
		ai, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			panic(err)
		}
		bspis = append(bspis, ai)
	}

	ctx := context.Background()

	ctx, events := notif.RegisterForQueryEvents(ctx)
	go func() {
		for e := range events {
			_ = e
			// if you want to see why things are broken
			fmt.Println("Event: ", e)
		}
	}()
	results := make([][]peer.ID, len(bspis))
	timing := make([]time.Duration, len(bspis))

	const k = "testkey"

	for i := 0; i < len(bspis); i++ {
		fmt.Println("Running query bench round ", i)
		start:=time.Now()
		peers, err := RunSingleCrawl(ctx, k, bspis[i:i+1])
		end := time.Now()
		if err != nil {
			fmt.Println("failed to run query: ", err)
			continue
		}

		results[i] = peers
		timing[i] = end.Sub(start)
	}

	fmt.Println("Results:")
	for i, r := range results {
		d := printDistances([]byte(k), results[i])
		fmt.Printf("%d: %d | time : %s | dist : %v \n", i, len(r), timing[i].String(), d)
	}
}

func printDistances(target []byte, peers []peer.ID) []int{
	t := keyspace.XORKeySpace.Key(target)
	d := make([]int, len(peers))
	for i, p := range peers {
		distb := xor(keyspace.XORKeySpace.Key([]byte(p)).Bytes, t.Bytes)
		dist := keyspace.ZeroPrefixLen(distb)
		d[i] = dist
		fmt.Printf("peer id=%s, distance=%d\n", p, dist)
	}
	return d
}

func xor(a, b kbucket.ID) kbucket.ID {
	return kbucket.ID(util.XOR(a, b))
}

func RunSingleCrawl(ctx context.Context, k string, bootstrap []*peer.AddrInfo) ([]peer.ID, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	h, err := libp2p.New(ctx)
	if err != nil {
		return nil, err
	}

	d, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}

	bspi := bootstrap[0]
	if err := h.Connect(ctx, *bspi); err != nil {
		return nil, err
	}

	peers, err := d.GetClosestPeersSingle(ctx, bspi.ID, k)
	if err != nil {
		return nil, err
	}

	var closest []peer.ID
	for _, p := range peers {
		closest = append(closest, p.ID)
	}
	return closest, nil
}
