package blog

import (
	"errors"
	"sync"
)

type PostStore struct {
	sync.Mutex
	posts  map[int]Post
	nextID int
}

func NewPostStore() *PostStore {
	return &PostStore{
		posts:  make(map[int]Post),
		nextID: 1,
	}
}

func (s *PostStore) Create(post Post) Post {
	s.Lock()
	defer s.Unlock()

	if post.ID == 0 {
		post.ID = s.nextID
	}
	if post.ID >= s.nextID {
		s.nextID = post.ID + 1
	}

	s.posts[post.ID] = post
	return post
}

func (s *PostStore) GetAll() []Post {
	s.Lock()
	defer s.Unlock()
	result := make([]Post, 0, len(s.posts))
	for _, p := range s.posts {
		result = append(result, p)
	}
	return result
}

func (s *PostStore) Get(id int) (Post, error) {
	s.Lock()
	defer s.Unlock()
	post, ok := s.posts[id]
	if !ok {
		return Post{}, errors.New("post not found")
	}
	return post, nil
}

func (s *PostStore) Update(id int, post Post) (Post, error) {
	s.Lock()
	defer s.Unlock()
	_, exists := s.posts[id]
	if !exists {
		return Post{}, errors.New("post not found")
	}
	post.ID = id
	s.posts[id] = post
	return post, nil
}

func (s *PostStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()
	if _, exists := s.posts[id]; !exists {
		return errors.New("post not found")
	}
	delete(s.posts, id)
	return nil
}
