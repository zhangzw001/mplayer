package library

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	    Id        	string
	    Name    	string
	    Artist    	string
	    Genre    	string
	    Source    	string
	    Type    	string
}

type MusicManger struct {
	musics []MusicEntry
}


func NewMusicManager() * MusicManger {
	return &MusicManger{make([]MusicEntry,0)}
}

func (m *MusicManger) Len() int {
	return len(m.musics)
}

func (m *MusicManger) Get(index int) (music *MusicEntry,err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index],nil
}

func (m *MusicManger) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	fmt.Println(m)
	fmt.Println(name)
	for _, m := range m.musics {
		fmt.Println(m)
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManger) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music )
}

func (m *MusicManger) Remove(index int) * MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics) - 1 {
		m.musics = append(m.musics[:index-1],m.musics[index+1:]...)
	}else if index == 0 {
		m.musics = make([]MusicEntry,0)
	}else {
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManger) RemoveByName(name string) *MusicEntry {
	var removedMusic *MusicEntry = nil
	var iPos int = -1
	for i := 0; i < m.Len(); i++ {
		if m.musics[i].Name == name {
			iPos = i
			break
		}
	}
	if iPos < 0 {
		return nil
	}
	removedMusic = m.Remove(iPos)
	return removedMusic
}