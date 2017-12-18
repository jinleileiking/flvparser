# flvparser

```
➜  flvparser git:(master) ✗ go run main.go ~/live_140400463_4440104.flv  --help
usage: main [<flags>] <file>

Flags:
      --help      Show context-sensitive help (also try --help-long and --help-man).
  -v, --video     Show video
  -a, --audio     Show audio
  -i, --keyframe  Show audio
      --version   Show application version.

Args:
```

```
➜  flvparser git:(master) ✗ go run main.go ~/live_140400463_4440104.flv  -i  -v
SPS:
(h264parser.SPSInfo) {
 ProfileIdc: (uint) 77,
 LevelIdc: (uint) 30,
 MbWidth: (uint) 30,
 MbHeight: (uint) 54,
 CropLeft: (uint) 0,
 CropRight: (uint) 0,
 CropTop: (uint) 0,
 CropBottom: (uint) 2,
 Width: (uint) 480,
 Height: (uint) 860
}
(*errors.errorString)(0xc42005e060)(EOF)
+-----+------+------+----------+-----------------+-------------+-------------+---------------+
| NO  | TYPE |  I   | DATASIZE | AVC PACKET TYPE | NALU FORMAT | NAL REF IDC | NAL UNIT TYPE |
+-----+------+------+----------+-----------------+-------------+-------------+---------------+
|   1 | H264 | true |       30 | SEQHDR          |             |
|   2 | H264 | true |    28815 | NALU            | AVCC        |           0 | IDR           |
|  52 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
|  53 | H264 | true |    22159 | NALU            | AVCC        |           0 | IDR           |
| 103 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 104 | H264 | true |    19598 | NALU            | AVCC        |           0 | IDR           |
| 154 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 155 | H264 | true |    25373 | NALU            | AVCC        |           0 | IDR           |
| 205 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 206 | H264 | true |    19279 | NALU            | AVCC        |           0 | IDR           |
| 256 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 257 | H264 | true |    15847 | NALU            | AVCC        |           0 | IDR           |
| 307 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 308 | H264 | true |    20169 | NALU            | AVCC        |           0 | IDR           |
| 358 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 359 | H264 | true |    26105 | NALU            | AVCC        |           0 | IDR           |
| 409 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 410 | H264 | true |    20325 | NALU            | AVCC        |           0 | IDR           |
| 460 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 461 | H264 | true |    20002 | NALU            | AVCC        |           0 | IDR           |
| 511 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 512 | H264 | true |    26779 | NALU            | AVCC        |           0 | IDR           |
| 562 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 563 | H264 | true |    15725 | NALU            | AVCC        |           0 | IDR           |
| 613 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 614 | H264 | true |    20948 | NALU            | AVCC        |           0 | IDR           |
| 664 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 665 | H264 | true |    23362 | NALU            | AVCC        |           0 | IDR           |
| 715 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 716 | H264 | true |    23604 | NALU            | AVCC        |           0 | IDR           |
| 766 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 767 | H264 | true |    27053 | NALU            | AVCC        |           0 | IDR           |
| 817 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 818 | H264 | true |    26764 | NALU            | AVCC        |           0 | IDR           |
| 868 | H264 | true |       40 | NALU            | AVCC        |           0 | SEI           |
| 869 | H264 | true |    23499 | NALU            | AVCC        |           0 | IDR           |
+-----+------+------+----------+-----------------+-------------+-------------+---------------+
```
