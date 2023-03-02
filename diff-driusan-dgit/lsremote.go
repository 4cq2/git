package git

type LsRemoteOptions struct {
	Heads, Tags   bool
	RefsOnly      bool
	Quiet         bool
	UploadPack    string
	ExitCode      bool
	GetURL        bool
	SymRef        bool
	Sort          string
	ServerOptions []string
}
