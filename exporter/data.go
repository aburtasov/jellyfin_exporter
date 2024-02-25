package exporter

type TranscodingInfo struct {
	AudioCodec           string   `json:"AudioCodec"`
	VideoCodec           string   `json:"VideoCodec"`
	Container            string   `json:"Container"`
	IsVideoDirect        bool     `json:"IsVideoDirect"`
	IsAudioDirect        bool     `json:"IsAudioDirect"`
	Bitrate              int      `json:"Bitrate"`
	Framerate            int      `json:"Framerate"`
	CompletionPercentage float64  `json:"CompletionPercentage"`
	Width                int      `json:"Width"`
	Height               int      `json:"Height"`
	AudioChannels        int      `json:"AudioChannels"`
	TranscodeReasons     []string `json:"TranscodeReasons"`
}

type PlayState struct {
	CanSeek             bool   `json:"CanSeek"`
	IsPaused            bool   `json:"IsPaused"`
	IsMuted             bool   `json:"IsMuted"`
	RepeatMode          string `json:"RepeatMode"`
	PositionTicks       int    `json:"PositionTicks"`
	VolumeLevel         int    `json:"VolumeLevel"`
	AudioStreamIndex    int    `json:"AudioStreamIndex"`
	SubtitleStreamIndex int    `json:"SubtitleStreamIndex"`
	MediaSourceId       string `json:"MediaSourceId"`
	PlayMethod          string `json:"PlayMethod"`
}

type NowPlayingQueue struct {
	ID             string `json:"ID"`
	PlaylistItemId string `json:"PlaylistItemId"`
}

type Capabilities struct {
	PlayableMediaTypes           []string `json:"PlayableMediaTypes"`
	SupportedCommands            []string `json:"SupportedCommands"`
	SupportsMediaControl         bool     `json:"SupportsMediaControl"`
	SupportsContentUploading     bool     `json:"SupportsContentUploading"`
	SupportsPersistentIdentifier bool     `json:"SupportsPersistentIdentifier"`
	SupportsSync                 bool     `json:"SupportsSync"`
}

type FullNowPlayingItem struct {
	Size                     int      `json:"Size"`
	Container                string   `json:"Container"`
	IsHD                     bool     `json:"IsHD"`
	IsShortcut               bool     `json:"IsShortcut"`
	Width                    int      `json:"Width"`
	Height                   int      `json:"Height"`
	ExtraIds                 []int    `json:"ExtraIds"`
	DateLastSaved            string   `json:"DateLastSaved"`
	RemoteTrailers           []string `json:"RemoteTrailers"`
	SupportsExternalTransfer bool     `json:"SupportsExternalTransfer"`
}
type Session struct {
	PlayState             PlayState          `json:"PlayState"`
	AdditionalUsers       []string           `json:"AdditionalUsers"`
	Capabilities          Capabilities       `json:"Capabilities"`
	RemoteEndPoint        string             `json:"RemoteEndPoint"`
	PlayableMediaTypes    []string           `json:"PlayableMediaTypes"`
	ID                    string             `json:"ID"`
	UserID                string             `json:"UserID"`
	UserName              string             `json:"UserName"`
	Client                string             `json:"Client"`
	LastActivityDate      string             `json:"LastActivityDate"`
	LastPlaybackCheckIn   string             `json:"LastPlaybackCheckIn"`
	DeviceName            string             `json:"DeviceName"`
	FullNowPlayingItem    FullNowPlayingItem `json:"FullNowPlayingItem"`
	DeviceId              string             `json:"DeviceId"`
	ApplicationVersion    string             `json:"ApplicationVersion"`
	IsActive              bool               `json:"IsActive"`
	SupportsMediaControl  bool               `json:"SupportsMediaControl"`
	SupportsRemoteControl bool               `json:"SupportsRemoteControl"`
	NowPlayingQueue       []NowPlayingQueue  `json:"NowPlayingQueue"`
	HasCustomDeviceName   bool               `json:"HasCustomDeviceName"`
	ServerId              string             `json:"ServerId"`
	SupportedCommands     []string           `json:"SupportedCommands"`
	TranscodingInfo       TranscodingInfo    `json:"TranscodingInfo"`
}

type Result []Session

type Statistics struct {
	MovieCount      int `json:"MovieCount"`
	SeriesCount     int `json:"SeriesCount"`
	EpisodeCount    int `json:"EpisodeCount"`
	ArtistCount     int `json:"ArtistCount"`
	ProgramCount    int `json:"ProgramCount"`
	TrailerCount    int `json:"TrailerCount"`
	SongCount       int `json:"SongCount"`
	AlbumCount      int `json:"AlbumCount"`
	MusicVideoCount int `json:"MusicVideoCount"`
	BoxSetCount     int `json:"BoxSetCount"`
	BookCount       int `json:"BookCount"`
	ItemCount       int `json:"ItemCount"`
}
