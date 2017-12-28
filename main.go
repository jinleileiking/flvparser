package main

import (
	"fmt"
	"os"
	"strconv"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/davecgh/go-spew/spew"

	"github.com/jinleileiking/joy4/av"
	"github.com/jinleileiking/joy4/av/avutil"
	"github.com/jinleileiking/joy4/codec/h264parser"
	"github.com/jinleileiking/joy4/format"
	// "github.com/nareix/joy4/av"
	// "github.com/nareix/joy4/av/avutil"
	// "github.com/nareix/joy4/format"
	"github.com/olekukonko/tablewriter"
)

func init() {
	format.RegisterAll()
}

var (
	filename = kingpin.Arg("file", "flv file").Required().String()
	show_v   = kingpin.Flag("video", "Show video").Short('v').Bool()
	show_a   = kingpin.Flag("audio", "Show audio").Short('a').Bool()
	show_i   = kingpin.Flag("keyframe", "Show audio").Short('i').Bool()
)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()
	// fmt.Printf("file: %s \n", *filename)

	file, err := avutil.Open(*filename)

	if err != nil {
		fmt.Println("Open file failed, detail:", err.Error())
		os.Exit(0)
	}
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Type", "I", "DataSize", "AVC Packet Type", "NALU format", "NAL Ref Idc", "NAL_UNIT_TYPE"})
	// table.SetBorder(false)

	// for i := 0; i < 10000; i++ {
	for true {
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

		// if streams[pkt.Idx].Type().String() == "H264" && pkt.IsKeyFrame {

		if *show_a {
			if streams[pkt.Idx].Type().String() == "AAC" {
				table.Append([]string{strconv.Itoa(v_cnt), streams[pkt.Idx].Type().String(), strconv.FormatBool(pkt.IsKeyFrame), strconv.Itoa(len(pkt.Data) + 5), pkt.AVCPacketType})
				// fmt.Println("video tag", v_cnt, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)

			}
		}

		if *show_v {
			if streams[pkt.Idx].Type().String() == "H264" {

				if *show_i {
					if pkt.IsKeyFrame {

						line := []string{
							strconv.Itoa(v_cnt),
							streams[pkt.Idx].Type().String(),
							strconv.FormatBool(pkt.IsKeyFrame),
							strconv.Itoa(len(pkt.Data) + 5),
							pkt.AVCPacketType,
							pkt.NALUFormat,
						}

						for _, info := range pkt.NALUInfos {
							line = append(line, strconv.Itoa(info.RefIdc))
							line = append(line, info.UnitType)
						}

						table.Append(line)
					}

					if pkt.AVCPacketType == "SEQHDR" {
						seq_hdr := streams[pkt.Idx].(h264parser.CodecData)
						fmt.Println("SPS:")
						spew.Dump(seq_hdr.SPSInfo)
						// fmt.Printf("\tWidth : %d\n", seq_hdr.SPSInfo.Width)
						// fmt.Printf("\tHeight : %d\n", seq_hdr.SPSInfo.Height)
						// fmt.Printf("\tProfileIdc : %d\n", seq_hdr.SPSInfo.ProfileIdc)
						// fmt.Printf("\tLevelIdc : %d\n", seq_hdr.SPSInfo.LevelIdc)
						// fmt.Printf("\tMbWidth : %d\n", seq_hdr.SPSInfo.MbWidth)
						// fmt.Printf("\tMbHeight : %d\n", seq_hdr.SPSInfo.MbHeight)
						// fmt.Printf("\tCropLeft : %d\n", seq_hdr.SPSInfo.CropLeft)
						// fmt.Printf("\tCropLeft : %d\n", seq_hdr.SPSInfo.CropLeft)
					}

				} else {

					var is_I string

					if pkt.IsKeyFrame {
						is_I = "I"
					} else {
						is_I = "B/P"
					}
					line := []string{
						strconv.Itoa(v_cnt),
						streams[pkt.Idx].Type().String(),
						strconv.FormatBool(pkt.IsKeyFrame),
						is_I,
						strconv.Itoa(len(pkt.Data) + 5),
						pkt.AVCPacketType,
						pkt.NALUFormat,
					}

					for _, info := range pkt.NALUInfos {
						line = append(line, strconv.Itoa(info.RefIdc))
						line = append(line, info.UnitType)
					}

					table.Append(line)
				}
				// fmt.Println("video tag", v_cnt, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)

			}
		}
		// table.Render() // Send output
	}

	table.Render() // Send output
	file.Close()
}
