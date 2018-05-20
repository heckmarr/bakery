package taste

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
	"github.com/jawher/mow.cli"
	"github.com/xlab/pocketsphinx-go/sphinx"
)

const (
	samplesPerChannel = 1024
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
	portaudio.Initialize()
	defer portaudio.Terminate()
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

	//file, err := os.Create("taste/sound.wav")
	//fileWriter := bufio.NewWriter(file)
	in := make([]int16, 10240)
	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(in), in)
	defer stream.Close()
	if err != nil {
		fmt.Println("Error opening default stream.")
	}

	stream.Start()
	defer stream.Stop()
	fmt.Println("Processing")
	for {
		decoder.StartUtt()
		stream.Read()

		decoder.ProcessRaw(in, false, false)
		if decoder.IsInSpeech() {
			fmt.Println("Listening...")
			decoder.ProcessRaw(in, false, true)
			decoder.EndUtt()

			fmt.Println(decoder.Hypothesis())
			fmt.Println("Done listening!")
		}
		decoder.EndUtt()
		//_, err := stream.AvailableToRead()
		//if err != nil {
		//	fmt.Println("Stream unable to read.")
		//}
		//fmt.Println(stream.Info())
		//fmt.Println(in)
		//fmt.Println(decoder.UttDuration())

		//		fmt.Println(stream.Time())
	}

	//stream.Stop()
	for {
		//fmt.Println("End listening.")
	}

}
