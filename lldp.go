package main

import (
	"fmt"
	"time"

	pb "github.com/nileshsimaria/jtisim/telemetry"
)

func streamLLDP(ch chan *pb.OpenConfigData, path *pb.Path) {
	pname := path.GetPath()
	freq := path.GetSampleFrequency()
	fmt.Println(pname, freq)

	seq := uint64(0)
	for {
		kv := []*pb.KeyValue{
			{Key: "__prefix__", Value: &pb.KeyValue_StrValue{StrValue: "/lldp"}},
			{Key: "state/foo", Value: &pb.KeyValue_UintValue{UintValue: 2222}},
		}

		d := &pb.OpenConfigData{
			SystemId:       "jtisim",
			ComponentId:    3,
			Timestamp:      uint64(MakeMSTimestamp()),
			SequenceNumber: seq,
			Kv:             kv,
		}
		ch <- d
		time.Sleep(time.Duration(freq) * time.Millisecond)
		seq++
	}
}
