package util

// Config estructura utilizada para guardar los datos del archivo
// de configuración del servicio consumer-1.ini
type Config struct {
	Logs          Logs
}

type Logs struct {
	LogPath         string
	LogName         string
	LogLevel        int // LevelDebug = 1,	LevelInfo  = 2,	LevelWarn  = 3,	LevelError = 4 , LevelDebug = 5 o superior(desactiva el log)
}
