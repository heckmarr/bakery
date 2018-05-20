package taste

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
	"github.com/jawher/mow.cli"
	"github.com/xlab/pocketsphinx-go/sphinx"
)

const (
	samplesPerChannel = 512
	sampleRate        = 16000
	channels          = 1
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

	//	file, err := os.Create("taste/sound.wav")
	//	fileReader := bufio.NewWriter(file)

	in := make([]int16, 1024)
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(in), in)
	defer stream.Close()
	if err != nil {
		fmt.Println("Error opening default stream.")
	}

	stream.Start()

	fmt.Println("Processing")
	decoder.StartUtt()
	for i := 0; i < 1; i++ {
		stream.Read()
		decoder.EndUtt()
		stream.Stop()
		fmt.Println(decoder.UttDuration())
		decoder.ProcessRaw(in, true, true)
		fmt.Println(decoder.Hypothesis())

	}
	for {

	}

}
