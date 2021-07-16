package conf

import "strings"

const (
	// ShortVersion 短版本号
	ShortVersion = "droneDeploy 0.0.1"
)

// The value of variables come form `gb build -ldflags '-X "build.Build=xxxxx" -X "build.CommitID=xxxx"' `
var (
	// Build build time
	Build string
	// Branch current git branch
	Branch string
	// Commit git commit id
	Commit string
	// Message last git commit message
	Message  string
	Platform string
)

// BuildVersion 生成版本信息
func BuildVersion() string {
	var buf strings.Builder
	buf.WriteString(ShortVersion)

	if Build != "" {
		buf.WriteByte('\n')
		buf.WriteString("build: ")
		buf.WriteString(Build)
	}
	if Branch != "" {
		buf.WriteByte('\n')
		buf.WriteString("branch: ")
		buf.WriteString(Branch)
	}
	if Commit != "" {
		buf.WriteByte('\n')
		buf.WriteString("commit: ")
		buf.WriteString(Commit)
	}
	if Message != "" {
		buf.WriteByte('\n')
		buf.WriteString("message: ")
		buf.WriteString(Message)
	}
	if Platform != "" {
		buf.WriteByte('\n')
		buf.WriteString("Platform: ")
		buf.WriteString(Platform)
	}
	return buf.String()
}
