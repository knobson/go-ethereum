package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/eth-go/ethlog"
	"os"
	"os/user"
	"path"
)

var Identifier string
var StartRpc bool
var RpcPort int
var UseUPnP bool
var OutboundPort string
var ShowGenesis bool
var AddPeer string
var MaxPeer int
var GenAddr bool
var UseSeed bool
var ImportKey string
var ExportKey bool
var NonInteractive bool
var Datadir string
var LogFile string
var ConfigFile string
var DebugFile string
var LogLevel int

// flags specific to cli client
var StartMining bool
var StartJsConsole bool
var InputFile string

func defaultDataDir() string {
	usr, _ := user.Current()
	return path.Join(usr.HomeDir, ".ethereum")
}

var defaultConfigFile = path.Join(defaultDataDir(), "conf.ini")

func Init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s [options] [filename]:\noptions precedence: default < config file < environment variables < command line\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&Identifier, "id", "", "Custom client identifier")
	flag.StringVar(&OutboundPort, "port", "30303", "listening port")
	flag.BoolVar(&UseUPnP, "upnp", false, "enable UPnP support")
	flag.IntVar(&MaxPeer, "maxpeer", 10, "maximum desired peers")
	flag.IntVar(&RpcPort, "rpcport", 8080, "port to start json-rpc server on")
	flag.BoolVar(&StartRpc, "rpc", false, "start rpc server")
	flag.BoolVar(&NonInteractive, "y", false, "non-interactive mode (say yes to confirmations)")
	flag.BoolVar(&UseSeed, "seed", true, "seed peers")
	flag.BoolVar(&GenAddr, "genaddr", false, "create a new priv/pub key")
	flag.BoolVar(&ExportKey, "export", false, "export private key")
	flag.StringVar(&LogFile, "logfile", "", "log file (defaults to standard output)")
	flag.StringVar(&ImportKey, "import", "", "imports the given private key (hex)")
	flag.StringVar(&Datadir, "datadir", defaultDataDir(), "specifies the datadir to use")
	flag.StringVar(&ConfigFile, "conf", defaultConfigFile, "config file")
	flag.StringVar(&DebugFile, "debug", "", "debug file (no debugging if not set)")
	flag.IntVar(&LogLevel, "loglevel", int(ethlog.InfoLevel), "loglevel: 0-5: silent,error,warn,info,debug,debug detail)")

	flag.BoolVar(&StartMining, "mine", false, "start dagger mining")
	flag.BoolVar(&StartJsConsole, "js", false, "launches javascript console")

	flag.Parse()

	InputFile = flag.Arg(0)
}