package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

// NewPostsFromFS returns a collection of blog posts from a file system.
// If it does not conform to the format then it'll return an error
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fs fs.FS, fileName string) (Post, error) {
	postFile, err := fs.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readLine()[len(titleSeparator):]
	descriptionLine := readLine()[len(descriptionSeparator):]

	return Post{Title: titleLine, Description: descriptionLine}, nil
}
