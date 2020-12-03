package nlib

import(
  "github.com/gin-gonic/gin"
	"net/http"
	//"strings"
	"path"
	//"fmt"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type EscFileSystem struct {
	Fs http.FileSystem
}

func (esc *EscFileSystem) Exists(urlPrefix string,fileUrl string)  bool {
	name:=urlPrefix+fileUrl
	//if name := strings.TrimPrefix(fileUrl, urlPrefix); len(name) < len(fileUrl) { //bizimle ilgili mi ?
	f, err := esc.Fs.Open(name);
	if err!=nil {
		return false
	}
	stats, err := f.Stat()
	if err != nil {
		return false
	}
	if stats.IsDir() {
		indexFile := path.Join(name, "index.html")
		_, err := esc.Fs.Open(indexFile);
		if err!=nil {
			return false
		}
	}
	return true
}

func (esc *EscFileSystem) Open(name string) (http.File, error) {
	return esc.Fs.Open(name)
}

/*type binaryFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func BinaryFile(root string, indexes bool) *binaryFileSystem {
	return &binaryFileSystem{
		FileSystem: Dir(root),
		root:       root,
		indexes:    indexes,
	}
}*/

/*
func (fs *binaryFileSystem) checkFileExists(urlPrefix string,fileUrl string) {

	if p := strings.TrimPrefix(fileUrl, urlPrefix); len(p) < len(fileUrl) { //bizimle ilgili mi ?
		f,err:=fs.Stat(name);

		if !fs.Exists(name) {
			return false
		}

		stats, err := fs.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !l.indexes {
				index := path.Join(name, INDEX)
				_, err := fs.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
} */

func StaticServeFromBinary(urlPrefix string,sfs ServeFileSystem) gin.HandlerFunc {
  // fmt.Println("Im a dummy!")
	// Middleware ilk olusturuldugunda calisacak bolge

	fileserver := http.FileServer(sfs)
	return func(c *gin.Context) {
		if sfs.Exists(urlPrefix, c.Request.URL.Path) {
			c.Request.URL.Path=urlPrefix+"/"+c.Request.URL.Path
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
