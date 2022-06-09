# flvparser

```
A stupid ugly flv / ts file parser

Usage:
  flvparser [flags]

Flags:
      --a             show audio
  -f, --file string   flv / ts file
  -h, --help          help for flvparser
      --non-key       use with -v:  do not show keyframes
      --sei           show sei info
      --simple        only show nal type
      --v             show video (default true)
```

```
./flvparser -f  ~/visionular/avfiles/s.flv --non-key    
(*errors.errorString)(0xc000110060)(EOF)
Parsed done
+------+------+-------+--------+--------+---------+----------+-----------------+-------------+----------+-------+-----+--------+------+---+---+
|  NO  | TYPE |   I   | FLVTS  |   TS   | TS DIFF | DATASIZE | AVC PACKET TYPE | NALU FORMAT |   NUT    | BYTES | IDC |        |      |   |   |
+------+------+-------+--------+--------+---------+----------+-----------------+-------------+----------+-------+-----+--------+------+---+---+
|    1 | H264 | true  |      0 |      0 |       0 |       46 | SEQHDR          |             |
|    2 | H264 | true  |      0 |      0 |       0 |     5137 | NALU            | AVCC        | (6)SEI   |   688 |   0 | (5)IDR | 4436 | 0 | I |
|    3 | H264 | false |     40 |     40 |      40 |      132 | NALU            | AVCC        | (1)N-IDR |   123 |   0 | P      |
|    4 | H264 | false |     80 |     80 |      40 |      989 | NALU            | AVCC        | (1)N-IDR |   980 |   0 | P      |
|    5 | H264 | false |    120 |    120 |      40 |       55 | NALU            | AVCC        | (1)N-IDR |    46 |   0 | B      |
|    6 | H264 | false |    160 |    160 |      40 |       62 | NALU            | AVCC        | (1)N-IDR |    53 |   0 | B      |
|    7 | H264 | false |    200 |    200 |      40 |       25 | NALU            | AVCC        | (1)N-IDR |    16 |   0 | B      |
|    8 | H264 | false |    240 |    240 |      40 |      193 | NALU            | AVCC        | (1)N-IDR |   184 |   0 | P      |
|    9 | H264 | false |    280 |    280 |      40 |       40 | NALU            | AVCC        | (1)N-IDR |    31 |   0 | B      |
|   10 | H264 | false |    320 |    320 |      40 |       30 | NALU            | AVCC        | (1)N-IDR |    21 |   0 | B      |
|   11 | H264 | false |    360 |    360 |      40 |       24 | NALU            | AVCC        | (1)N-IDR |    15 |   0 | B      |
|   12 | H264 | false |    400 |    400 |      40 |       73 | NALU            | AVCC        | (1)N-IDR |    64 |   0 | P      |
|   13 | H264 | false |    440 |    440 |      40 |       26 | NALU            | AVCC        | (1)N-IDR |    17 |   0 | B      |
|   14 | H264 | false |    480 |    480 |      40 |       23 | NALU            | AVCC        | (1)N-IDR |    14 |   0 | B      |
|   15 | H264 | false |    520 |    520 |      40 |       25 | NALU            | AVCC        | (1)N-IDR |    16 |   0 | B      |
|   16 | H264 | false |    560 |    560 |      40 |      439 | NALU            | AVCC        | (1)N-IDR |   430 |   0 | P      |
|   17 | H264 | false |    600 |    600 |      40 |     3008 | NALU            | AVCC        | (1)N-IDR |  2999 |   0 | P      |
|   18 | H264 | false |    640 |    640 |      40 |      604 | NALU            | AVCC        | (1)N-IDR |   595 |   0 | P      |
|   19 | H264 | false |    680 |    680 |      40 |       28 | NALU            | AVCC        | (1)N-IDR |    19 |   0 | B      |
|   20 | H264 | false |    720 |    720 |      40 |       58 | NALU            | AVCC        | (1)N-IDR |    49 |   0 | B      |
|   21 | H264 | false |    760 |    760 |      40 |       25 | NALU            | AVCC        | (1)N-IDR |    16 |   0 | B      |
|   22 | H264 | false |    800 |    800 |      40 |     3429 | NALU            | AVCC        | (1)N-IDR |  3420 |   0 | P      |
|   23 | H264 | false |    840 |    840 |      40 |       64 | NALU            | AVCC        | (1)N-IDR |    55 |   0 | B      |

```
