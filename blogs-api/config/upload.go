package config

type Upload struct {
	Root      string   `ini:"root"`
	Cdn       string   `ini:"cdn"`
	ImageExts []string `ini:"imageExts"`
	MaxImage  int64    `ini:"maxImage"`
	ImagePath string   `ini:"imagePath"`
	FileExts  []string `ini:"fileExts"`
	MaxFile   int64    `ini:"maxFile"`
	FilePath  string   `ini:"filePath"`
}
