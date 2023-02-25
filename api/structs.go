package api

type Perf struct {
	Games int
	Rating int
	Rd int
	Prog int
	Prov bool
}

type Perfs struct {
	Chess960 Perf
	Atomic Perf
	RacingKings Perf
	UltraBullet Perf
	Blitz Perf
	KingOfTheHill Perf
	Bullet Perf
	Correspondence Perf
	Horde Perf
	Puzzle Perf
	Classical Perf
	Rapid Perf
	Storm Perf
}

type Profile struct {
	Country string
	Location string
	Bio string
	FirstName string 
	LastName string
	FideRating int
	UscfRating int
	EcfRating int
	Links string
}

type Playtime struct {
	Total int
	Tv int
}

type Count struct {
	All int
	Rated int
	Ai int
	Draw int
	DrawH int
	Loss int
	LossH int
	Win int
	WinH int
	Bookmark int
	Playing int
	Import int
	Me int
}

type Account struct {
	Id string
	Username string
	Perfs Perfs
	CreatedAt int64
	Disabled bool
	TosViolation bool
	Profile Profile
	SeenAt int64
	Patron bool
	Verified bool
	PlayTime Playtime
	Title string
	Url string
	Playing string
	CompletionRate int
	Count Count
	Streaming bool
	Followable bool
	Following bool
	Blocking bool
	FollowsYou bool
}

type Email struct {
	Email string
}

type Prefs struct {
	Dark bool
	Transp bool
	BgImg string
	Is3d bool
	Theme string
	PieceSet string
	Theme3d string
	PieceSet3d string
	SoundSet string
	BlindFold int
	AutoQueen int
	AutoThreeFold int
	Takeback int
	Moretime int
	ClockTenths int
	ClockBar bool
	ClockSound bool
	Premove bool
	Animation int
	Captured bool
	Follow bool
	Highlight bool
	Destination bool
	Coords int
	Replay int
	Challenge int
	Message int
	CoordColor int
	SubmitMove int
	ConfirmResign int
	InsightShare int
	KeyboardMove  int
	Zen int
	MoveEvent int
	RookCastle int
}

type Preferences struct {
	Prefs Prefs
	Language string
}

type Kid struct {
	kid bool
}
