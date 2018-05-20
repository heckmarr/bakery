package taste

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/xlab/pocketsphinx-go/sphinx"
	"github.com/xlab/portaudio-go/portaudio"
)

const (
	samplesPerChannel = 512
	sampleRate        = 16000
	channels          = 1
	sampleFormat      = portaudio.PaInt16
)

var (
	app     = cli.App("tasteTest", "This is a test system to see how well sphinx works with golang.")
	hmm     = app.StringOpt("hmm", "/usr/local/share/pocketsphinx/model/en-us/en-us", "Sets directory containing acoustic model files.")
	dict    = app.StringOpt("dict", "/usr/local/share/pocketsphinx/model/en-us/cmudict-en-us.dict", "Sets main pronunciation dictionary (lexicon) input file..")
	lm      = app.StringOpt("lm", "/usr/local/share/pocketsphinx/model/en-us/en-us.lm.bin", "Sets word trigram language model input file.")
	logfile = app.StringOpt("log", "taste.log", "Log file to write log to.")
	stdout  = app.BoolOpt("stdout", false, "Disables log file and writes everything to stdout.")
	outraw  = app.StringOpt("outraw", "", "Specify output dir for RAW recorded sound files (s16le). Directory must exist.")
)

func Listen() {
	listener := portaudio.Initialize()
	fmt.Println(listener)
	//defer listener.Close()
	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption(*hmm),
		sphinx.DictFileOption(*dict),
		sphinx.LMFileOption(*lm),
		sphinx.SampleRateOption(sampleRate),
	)

	fmt.Println("Loading CMU sphinx...")
	decoder, err := sphinx.NewDecoder(cfg)
	if err != nil {
		fmt.Println("Error creating decoder!")
	}
	var stream *portaudio.Stream
	error := portaudio.OpenDefaultStream(&stream, channels, 0, sampleFormat, sampleRate,
		samplesPerChannel, nil, nil)

	fmt.Println(error)

	portaudio.StartStream(stream)
	decoder.StartUtt()
	for {

	}

}
