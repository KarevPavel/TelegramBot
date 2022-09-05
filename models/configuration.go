package models

// Client Client
type Client struct {
	Token       string
	RequestURL  string
	RequestFile string
}

type Aria2C struct {
	DownloadDir             string
	SourcesDir              string
	LogDir                  string
	Secret                  string
	Port                    int
	MaxConnectionsPerServer int
	MaxConcurrentDownloads  int
	LogLevel                string
}

type DbConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// Conf Conf
type Conf struct {
	Title        string
	Client       Client
	Aria2C       Aria2C
	DbConnection DbConnection
}
