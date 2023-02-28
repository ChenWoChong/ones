package main

import (
	"flag"
	"github.com/ChenWoChong/ones/pkg/features"
	"github.com/ChenWoChong/ones/pkg/util/feature"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"math/rand"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

var (
	bindAddr  = flag.String("addr", ":10221", "addr")
	pprofAddr = flag.String("pprof-addr", ":10222", "pprof")
)

func main() {
	feature.DefaultMutableFeatureGate.AddFlag(pflag.CommandLine)
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	rand.Seed(time.Now().UnixNano())
	features.SetDefaultFeatureGates()
	ctrl.SetLogger(klogr.New())
}
