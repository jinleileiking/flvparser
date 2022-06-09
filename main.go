package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/jinleileiking/joy4/av"
	"github.com/jinleileiking/joy4/av/avutil"
	"github.com/jinleileiking/joy4/codec/h264parser"
	"github.com/jinleileiking/joy4/format"
	"github.com/jinleileiking/joy4/format/ts"

	// "github.com/nareix/joy4/av"
	// "github.com/nareix/joy4/av/avutil"
	// "github.com/nareix/joy4/format"
	"github.com/olekukonko/tablewriter"
)

func init() {
	format.RegisterAll()
}

var rootCmd = &cobra.Command{
	Use:   "flvparser",
	Short: "A stupid ugly flv / ts file parser",
	Run:   cmdrun,
}

var last_ts int

func cmdrun(cmd *cobra.Command, args []string) {
	file, err := avutil.Open(filename)

	if err != nil {
		fmt.Println("Open file failed, detail:", err.Error())
		os.Exit(0)
	}

	// ext := path.Ext(*filename)

	// spew.Dump(ext)

	// _, ok := file.(*ts.Demuxer)

	// if ok {
	// 	fmt.Println("ts format")
	// }
	// file, _ := avutil.Open("live_140400463_4440104.flv")

	// spew.Dump(file, err)

	streams, err := file.Streams()

	if err != nil {
		spew.Dump(err)
		// return
	}

	// hd, ok := file.(*avutil.HandlerDemuxer)
	hd, _ := file.(*avutil.HandlerDemuxer)

	// if ok {
	// 	fmt.Println("ts format")
	// }

	td, ok := hd.Demuxer.(*ts.Demuxer)

	// spew.Dump(td)

	// os.Exit(0)

	var is_ts bool

	if ok {
		fmt.Println("ts format")
		is_ts = true
	}

	// spew.Dump(hd.Demuxer)
	// os.Exit(0)

	if is_ts {
		spew.Dump(td.Pat)
		spew.Dump(td.Pmt)
	}

	var a_cnt int
	var v_cnt int

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(80)
	// table.SetBorder(false)
	headers := []string{"PTS", "DTS"}
	defer func() {
		table.Render() // Send output
		file.Close()
	}()

	if is_ts {

		// header := []string{"T", "L", "IDC"}

		for _, payload := range td.PayloadInfos {
			nalues := h264parser.ParseNALUs(payload.PayloadInfo)
			// spew.Dump(nalues)

			line := []string{
				// strconv.Itoa(v_cnt),
				// streams[pkt.Idx].Type().String(),
				// strconv.FormatBool(pkt.IsKeyFrame),
				// strconv.Itoa(len(pkt.Data) + 5),
				// strconv.Itoa(int(pkt.Time) / 1000000),
				// strconv.Itoa(int(pkt.Time)/1000000 - last_ts),
				// pkt.AVCPacketType,
				// pkt.NALUFormat,
			}

			line = append(line, strconv.Itoa(payload.Pts/1000))
			line = append(line, strconv.Itoa(payload.Dts/1000))
			line = append(line, nalues.NALUFormat)
			for _, info := range nalues.Infos {
				line = append(line, info.UnitType)

				if !show_only_nalt {
					line = append(line, strconv.Itoa(info.NumBytes))
					line = append(line, strconv.Itoa(info.RefIdc))
					if info.UnitType == "N-IDR" ||
						info.UnitType == "SliceA" ||
						info.UnitType == "SliceB" ||
						info.UnitType == "SliceC" ||
						info.UnitType == "IDR" {
						line = append(line, info.SliceType)
					}
					if show_sei && info.UnitType == "SEI" {
						line = append(line, hex.Dump(info.Data))
					}
				}

			}
			// headers = append(headers, header...)
			table.Append(line)
		}
		// for _, payload := range td.Payloads {
		// 	nalues := h264parser.ParseNALUs(payload)
		// 	// spew.Dump(nalues)

		// 	line := []string{
		// 	// strconv.Itoa(v_cnt),
		// 	// streams[pkt.Idx].Type().String(),
		// 	// strconv.FormatBool(pkt.IsKeyFrame),
		// 	// strconv.Itoa(len(pkt.Data) + 5),
		// 	// strconv.Itoa(int(pkt.Time) / 1000000),
		// 	// strconv.Itoa(int(pkt.Time)/1000000 - last_ts),
		// 	// pkt.AVCPacketType,
		// 	// pkt.NALUFormat,
		// 	}

		// 	line = append(line, nalues.NALUFormat)
		// 	for _, info := range nalues.Infos {
		// 		line = append(line, info.UnitType)
		// 		line = append(line, strconv.Itoa(info.NumBytes))
		// 		line = append(line, strconv.Itoa(info.RefIdc))

		// 		if info.UnitType == "N-IDR" ||
		// 			info.UnitType == "SliceA" ||
		// 			info.UnitType == "SliceB" ||
		// 			info.UnitType == "SliceC" ||
		// 			info.UnitType == "IDR" {
		// 			line = append(line, info.SliceType)
		// 		}

		// 		if *show_sei && info.UnitType == "SEI" {
		// 			line = append(line, hex.Dump(info.Data))
		// 		}
		// 	}
		// 	// headers = append(headers, header...)
		// 	table.Append(line)
		// }

		table.SetHeader(headers)
		return
	}

	//flv
	// table.SetHeader([]string{"No", "Type", "I", "FLVTS", "TS", "TS Diff", "DataSize", "AVC Packet Type", "NALU format", "NAL_UNIT_TYPE", "Num bytes", "NAL Ref Idc"})
	table.SetHeader([]string{"No", "Type", "I", "FLVTS", "TS", "TS Diff", "DataSize", "AVC Packet Type", "NALU format",
		"NUT", "BYTES", "Idc"})
	for true {
		var pkt av.Packet
		var err error
		if pkt, err = file.ReadPacket(); err != nil {
			spew.Dump(err)
			fmt.Println("Parsed done")
			break
		}
		if streams[pkt.Idx].Type().String() == "H264" {
			v_cnt = v_cnt + 1
		}

		if streams[pkt.Idx].Type().String() == "AVC" {
			a_cnt = a_cnt + 1
		}

		if show_a {
			if streams[pkt.Idx].Type().String() == "AAC" {
				table.Append([]string{strconv.Itoa(v_cnt), streams[pkt.Idx].Type().String(), strconv.FormatBool(pkt.IsKeyFrame), strconv.Itoa(len(pkt.Data) + 5), pkt.AVCPacketType})
			}
		}

		if show_v {
			if streams[pkt.Idx].Type().String() == "H264" {

				if !no_show_i {
					if !pkt.IsKeyFrame {
						continue
					}
				}

				line := []string{
					strconv.Itoa(v_cnt),
					streams[pkt.Idx].Type().String(),
					strconv.FormatBool(pkt.IsKeyFrame),
					strconv.Itoa(int(pkt.Timestamp)),
					strconv.Itoa(int(pkt.Time) / 1000000),
					strconv.Itoa(int(pkt.Timestamp) - last_ts),
					strconv.Itoa(len(pkt.Data) + 5),
					pkt.AVCPacketType,
					pkt.NALUFormat,
				}
				last_ts = int(pkt.Timestamp)
				// spew.Dump(pkt.Timestamp, last_ts)

				for _, info := range pkt.NALUInfos {
					line = append(line, info.UnitType)

					if !show_only_nalt {
						line = append(line, strconv.Itoa(info.NumBytes))
						line = append(line, strconv.Itoa(info.RefIdc))

						if info.UnitType == "(1)N-IDR" ||
							info.UnitType == "(2)SliceA" ||
							info.UnitType == "(3)SliceB" ||
							info.UnitType == "(4)SliceC" ||
							info.UnitType == "(5)IDR" {
							line = append(line, info.SliceType)
						}

						if show_sei && info.UnitType == "SEI" {
							line = append(line, hex.Dump(info.Data))
						}
					}
				}

				table.Append(line)

			}
		}

		// last_ts = int(pkt.Time) / 1000000
	}

}

var show_sei bool
var show_v bool
var show_a bool
var no_show_i bool
var show_only_nalt bool
var filename string

func setup_cmd() {
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "flv / ts file")
	rootCmd.PersistentFlags().BoolVar(&show_sei, "sei", false, "show sei info")
	rootCmd.PersistentFlags().BoolVar(&show_only_nalt, "simple", false, "only show nal type")
	rootCmd.PersistentFlags().BoolVar(&show_a, "a", false, "show audio")
	rootCmd.PersistentFlags().BoolVar(&show_v, "v", true, "show video,  default:true")
	rootCmd.PersistentFlags().BoolVar(&no_show_i, "non-key", false, "use with -v:  show BP frames")
	rootCmd.MarkFlagRequired("file")
}

func main() {
	setup_cmd()
	rootCmd.Execute()
}
