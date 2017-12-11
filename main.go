package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
)

func init() {
	format.RegisterAll()
}

var (
	filename = kingpin.Arg("file", "flv file").Required().String()
)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("file: %s \n", *filename)

	file, _ := avutil.Open(*filename)
	// file, _ := avutil.Open("live_140400463_4440104.flv")

	// spew.Dump(file, err)

	streams, _ := file.Streams()
	// for _, stream := range streams {
	// 	if stream.Type().IsAudio() {
	// 		astream := stream.(av.AudioCodecData)
	// 		fmt.Println(astream.Type(), astream.SampleRate(), astream.SampleFormat(), astream.ChannelLayout())
	// 	} else if stream.Type().IsVideo() {
	// 		vstream := stream.(av.VideoCodecData)
	// 		fmt.Println(vstream.Type(), vstream.Width(), vstream.Height())
	// 	}
	// }

	var a_cnt int
	var v_cnt int

	for i := 0; i < 10000; i++ {
		var pkt av.Packet
		var err error
		if pkt, err = file.ReadPacket(); err != nil {
			spew.Dump(err)
			break
		}

		if streams[pkt.Idx].Type().String() == "H264" {
			v_cnt = v_cnt + 1
		}

		if streams[pkt.Idx].Type().String() == "AVC" {
			a_cnt = a_cnt + 1
		}

		if streams[pkt.Idx].Type().String() == "H264" && pkt.IsKeyFrame {
			fmt.Println("video tag", v_cnt, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
		}
	}

	file.Close()
}
