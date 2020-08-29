package storage

type option struct{}

var Opt option

type storageFile struct {
	path     string
	size     int
	line     int
	day      int
	total    int
	compress bool
}

func (o *option) WithFile(path string) *storageFile {
	var opt = new(storageFile)
	opt.path = path
	return opt
}

// MaxSize Mb
func (o *storageFile) MaxSize(M int) *storageFile {
	o.size = M * (1 << 20)
	return o
}

// Maximum file line
func (o *storageFile) MaxLine(line int) *storageFile {
	o.line = line
	return o
}

// SaveTime day
func (o *storageFile) SaveTime(day int) *storageFile {
	o.day = day
	return o
}

// 开启压缩
func (o *storageFile) Compress() *storageFile {
	o.compress = true
	return o
}

// 最多保存文件数量
func (o *storageFile) Backups(total int) *storageFile {
	o.total = total
	return o
}

func (o *storageFile) Done() *rotate {
	return newFileWrite(o.path, o.size, o.line, o.day, o.total, o.compress)
}
