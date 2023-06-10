package models

type PictureInfo struct {
	Pictureid int    `JSON:"pictureid"`
	Noteid    int    `JSON:"noteid"`
	PicUrl    string `JSON:"picUrl"` //图片路径
	PicTag    string `JSON:"picTag"` // 标签
}
