package some

import (
	"fmt"
	"github.com/newtools/zsocket"
	"github.com/newtools/zsocket/nettypes"
	"testing"
)

//zero copy 在go中的实现
func TestZSocket(t *testing.T) {
	// args: interfaceIndex, options, maxFrameSize, and maxTotalFrames

	// inerfaceIndex: the index of the net device you want to open a raw socket to
	// options: RX and TX, or just one or the other?
	// maxFrameSize: must be a power of 2, bigger than zsocket.MinimumFrameSize,
	// 	and smaller than maximum frame size
	// maxTotalFrames: must be at least 16, and be a multiple of 8.
	zs, err := zsocket.NewZSocket(14, zsocket.EnableRX, 2048, 64, nettypes.All)
	// the above will result in a ring buffer of 64 frames at
	// 	(2048 - zsocket.PacketOffset()) *writeable* bytes each
	// 	for a total of 2048*64 bytes of *unswappable* system memory consumed.
	if err != nil {
		panic(err)
	}
	zs.Listen(func(f *nettypes.Frame, frameLen, capturedLen uint16) {
		fmt.Printf(f.String(capturedLen, 0))
	})
}
