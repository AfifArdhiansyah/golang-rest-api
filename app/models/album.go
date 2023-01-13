package app

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  int64
}

func AlbumModel() []Album {
	var album []Album
	return album
}
