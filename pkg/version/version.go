package version

// AppVersion is the app-global version string, which should be substituted with a
// real value during build.
var (
	AppVersion = "UNKNOWN"
	BuildDate  string // date -u
)
