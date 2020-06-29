
package library

import (
	"testing"
)

func TestOps( t * testing.T ) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager faild.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager faild , not empty")
	}

	m0 := &MusicEntry{"1","My Heart Will Go On","Celion Dion", "Pop", "http://qbox.me/24501234", "MP3"}
	mm.Add(m0)

	if mm.Len() !=1 {
		t.Error("NewMusicManager.Add faild.")
	}
	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find faild")
	}

	if  m.Id != m0.Id ||
		m.Name != m0.Name ||
		m.Artist != m0.Artist ||
		m.Genre != m0.Genre ||
		m.Source != m0.Source ||
		m.Type != m0.Type {
		t.Error("MusicManager.Find() faild. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() faild.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() faild.", err)
	}
}