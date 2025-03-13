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

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	post := Post{Title: titleLine[7:], Description: descriptionLine[13:]}
	return post, nil
}
