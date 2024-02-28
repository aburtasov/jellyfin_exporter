package exporter

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Name      = "jellyfin"
	Namespace = "jellyfin"
)

type Exporter struct {
	apiUrl string
	apiKey string

	activeUsers                  *prometheus.Desc
	activeStreamsDirectPlayCount *prometheus.Desc
	activeStreamsTranscodeCount  *prometheus.Desc

	movieCount      *prometheus.Desc
	seriesCount     *prometheus.Desc
	episodeCount    *prometheus.Desc
	artistCount     *prometheus.Desc
	programCount    *prometheus.Desc
	trailerCount    *prometheus.Desc
	songCount       *prometheus.Desc
	albumCount      *prometheus.Desc
	musicVideoCount *prometheus.Desc
	boxSetCount     *prometheus.Desc
	bookCount       *prometheus.Desc
	itemCount       *prometheus.Desc
}

func New(apiUrl string, apiKey string, timeout time.Duration) *Exporter {

	e := &Exporter{
		apiUrl: apiUrl,
		apiKey: apiKey,
		activeUsers: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "active_users_count"),
			"Number of current active users",
			nil,
			nil,
		),

		activeStreamsDirectPlayCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "active_streams_direct_play_count"),
			"Number of current active streams direct play",
			nil,
			nil,
		),
		activeStreamsTranscodeCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "active_streams_transcode_count"),
			"Number of current active streams transcode",
			nil,
			nil,
		),

		movieCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "movie_count"),
			"Total number of movies",
			nil,
			nil,
		),
		seriesCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "series_count"),
			"Total number of series",
			nil,
			nil,
		),
		episodeCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "episode_count"),
			"Total number of episodes",
			nil,
			nil,
		),
		artistCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "artist_count"),
			"Total number of artists",
			nil,
			nil,
		),
		programCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "program_count"),
			"Total number of programs",
			nil,
			nil,
		),
		trailerCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "trailer_count"),
			"Total number of trailers",
			nil,
			nil,
		),
		songCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "song_count"),
			"Total number of songs",
			nil,
			nil,
		),
		albumCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "album_count"),
			"Total number of albums",
			nil,
			nil,
		),
		musicVideoCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "music_video_count"),
			"Total number of music_videos",
			nil,
			nil,
		),
		boxSetCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "box_set_count"),
			"Total number of box_sets",
			nil,
			nil,
		),
		bookCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "book_count"),
			"Total number of books",
			nil,
			nil,
		),
		itemCount: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "item_count"),
			"Total number of items",
			nil,
			nil,
		),
	}

	return e
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.activeUsers
	ch <- e.activeStreamsDirectPlayCount
	ch <- e.activeStreamsTranscodeCount

	ch <- e.movieCount
	ch <- e.seriesCount
	ch <- e.episodeCount
	ch <- e.artistCount
	ch <- e.programCount
	ch <- e.trailerCount
	ch <- e.songCount
	ch <- e.albumCount
	ch <- e.musicVideoCount
	ch <- e.boxSetCount
	ch <- e.bookCount
	ch <- e.itemCount

}

func (e *Exporter) GetName() string {
	return Name
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	var (
		result                 Result
		statistics             Statistics
		sessionsCount          int = 0
		streamsCount           int = 0
		streamsDirectPlayCount int = 0
		streamsTranscodeCount  int = 0
	)

	urlSessions := e.apiUrl + "/sessions"
	urlItems := e.apiUrl + "/items/counts"

	///////////////////////////////////////// Starting request to /sessions ///////////////////////////////////////////////////////////////////////

	reqSessions, err := http.NewRequest("GET", urlSessions, nil)
	if err != nil {
		log.Fatal(err)
	}

	qSessions := reqSessions.URL.Query()
	qSessions.Add("ApiKey", e.apiKey)
	reqSessions.URL.RawQuery = qSessions.Encode()

	clientSessions := &http.Client{}
	respSessions, err := clientSessions.Do(reqSessions)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer respSessions.Body.Close()

	bodySessions, err := io.ReadAll(respSessions.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	if err := json.Unmarshal(bodySessions, &result); err != nil {
		log.Fatal(err)
	}

	for _, user := range result {
		if user.IsActive == true {
			sessionsCount += 1
		}
		if user.FullNowPlayingItem.Size != 0 {
			streamsCount += 1
		}
		if user.PlayState.PlayMethod == "" {
			streamsDirectPlayCount += 1
		}
		if user.PlayState.PlayMethod == "Transcode" {

			streamsTranscodeCount += 1
		}

	}
	///////////////////////////////////////// End of request to /sessions //////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////// Starting request to /items/counts ///////////////////////////////////////////////////////////////
	reqItems, err := http.NewRequest("GET", urlItems, nil)
	if err != nil {
		log.Fatal(err)
	}

	qItems := reqItems.URL.Query()
	qItems.Add("ApiKey", e.apiKey)
	reqItems.URL.RawQuery = qItems.Encode()

	clientItems := &http.Client{}
	respItems, err := clientItems.Do(reqItems)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer respItems.Body.Close()

	bodyItems, err := io.ReadAll(respItems.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	if err := json.Unmarshal(bodyItems, &statistics); err != nil {
		log.Fatal(err)
	}
	////////////////////////////////////////////// End of request to /items/counts ////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////// Send metrics ////////////////////////////////////////////////////////////////////////////////

	ch <- prometheus.MustNewConstMetric(
		e.activeUsers,
		prometheus.CounterValue,
		float64(sessionsCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.activeStreamsDirectPlayCount,
		prometheus.CounterValue,
		float64(streamsDirectPlayCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.activeStreamsTranscodeCount,
		prometheus.CounterValue,
		float64(streamsTranscodeCount),
	)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	ch <- prometheus.MustNewConstMetric(
		e.movieCount,
		prometheus.CounterValue,
		float64(statistics.MovieCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.seriesCount,
		prometheus.CounterValue,
		float64(statistics.SeriesCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.episodeCount,
		prometheus.CounterValue,
		float64(statistics.EpisodeCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.artistCount,
		prometheus.CounterValue,
		float64(statistics.ArtistCount),
	)
	ch <- prometheus.MustNewConstMetric(
		e.programCount,
		prometheus.CounterValue,
		float64(statistics.ProgramCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.trailerCount,
		prometheus.CounterValue,
		float64(statistics.TrailerCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.songCount,
		prometheus.CounterValue,
		float64(statistics.SongCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.albumCount,
		prometheus.CounterValue,
		float64(statistics.AlbumCount),
	)
	ch <- prometheus.MustNewConstMetric(
		e.musicVideoCount,
		prometheus.CounterValue,
		float64(statistics.MusicVideoCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.boxSetCount,
		prometheus.CounterValue,
		float64(statistics.BoxSetCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.bookCount,
		prometheus.CounterValue,
		float64(statistics.BookCount),
	)

	ch <- prometheus.MustNewConstMetric(
		e.itemCount,
		prometheus.CounterValue,
		float64(statistics.ItemCount),
	)

}
