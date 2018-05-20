package taste

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
	"github.com/xlab/pocketsphinx-go/sphinx"
)

const (
	samplesPerChannel = 1024
	sampleRate        = 16000
	channels          = 1
)

var (
	//	app     = cli.App("tasteTest", "This is a test system to see how well sphinx works with golang.")
	hmm     = "/usr/local/share/pocketsphinx/model/en-us/en-us"
	dict    = "/usr/local/share/pocketsphinx/model/en-us/cmudict-en-us.dict"
	lm      = "/usr/local/share/pocketsphinx/model/en-us/en-us.lm.bin"
	logfile = "/dev/null"
	//	logfile = app.StringOpt("log", "taste.log", "Log file to write log to.")
	stdout = false

//	outraw  = app.StringOpt("outraw", "", "Specify output dir for RAW recorded sound files (s16le). Directory must exist.")
)

func Listen() ([]int16, *portaudio.Stream, *sphinx.Decoder) {
	portaudio.Initialize()
	//	defer portaudio.Terminate()
	//defer listener.Close()
	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption(hmm),
		sphinx.DictFileOption(dict),
		sphinx.LMFileOption(lm),
		sphinx.SampleRateOption(sampleRate),
	)
	if stdout == false {
		sphinx.LogFileOption(logfile)(cfg)
	}
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
	//	defer stream.Stop()
	fmt.Println("Processing")
	return in, stream, decoder
}
func Plug(stream *portaudio.Stream, decoder *sphinx.Decoder) {
	stream.Stop()
	stream.Close()
	portaudio.Terminate()
}
func Interpret(in []int16, stream *portaudio.Stream, decoder *sphinx.Decoder) string {
	decoder.StartUtt()
	stream.Read()
	var word string
	decoder.ProcessRaw(in, false, false)
	if decoder.IsInSpeech() {
		//		fmt.Println("Listening...")
		decoder.ProcessRaw(in, false, true)
		decoder.EndUtt()

		word, _ = decoder.Hypothesis()
		//		fmt.Println("Done listening!")
	}
	decoder.EndUtt()
	return word
}
