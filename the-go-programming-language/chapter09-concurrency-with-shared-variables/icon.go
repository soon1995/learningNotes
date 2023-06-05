package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// NOTE: not concurrency-safe!
// func Icon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons() // one-time initialization
// 	}
// 	return icons[name]
// }

func loadIcon(name string) image.Image {
	return nil
}

// var mu sync.Mutex

// Concurrency -safe
// func Icon(name string) image.Image {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if icons == nil {
// 		loadIcons() // one-time initialization
// 	}
// 	return icons[name]
// }

var muIcon sync.RWMutex

// Give greater concurrency but is complex and thus error-prone.
// func Icon(name string) image.Image {
// 	muIcon.RLock()
// 	if icons != nil {
// 		icon := icons[name]
// 		muIcon.Unlock()
// 		return icon
// 	}
// 	defer muIcon.RUnlock()

// 	// acquire an exclusive lock
// 	muIcon.Lock()
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	icon := icons[name]
// 	muIcon.Unlock()
// 	return icon
// }

var loadIconsOnce sync.Once

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
