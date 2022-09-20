package config

type QiNiuYun struct {
	AccessKey string `ini:"accessKey"`
	SerectKey string `ini:"serectKey"`
	Bucket    string `ini:"bucket"`
	ImgUrl    string `ini:"imgUrl"`
}
