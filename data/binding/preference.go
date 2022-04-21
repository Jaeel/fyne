// auto-generated
// **** THIS FILE IS AUTO-GENERATED, PLEASE DO NOT EDIT IT **** //

package binding

import (
	"sync"

	"github.com/Jaleel/fyne"
)

const keyTypeMismatchError = "A previous preference binding exists with different type for key: "

type prefBoundBool struct {
	base
	key string
	p   fyne.Preferences

	cacheLock sync.RWMutex
	cache     bool
}

// BindPreferenceBool returns a bindable bool value that is managed by the application preferences.
// Changes to this value will be saved to application storage and when the app starts the previous values will be read.
//
// Since: 2.0
func BindPreferenceBool(key string, p fyne.Preferences) Bool {
	binds := prefBinds.getBindings(p)
	if binds != nil {
		if listen := binds.getItem(key); listen != nil {
			if l, ok := listen.(Bool); ok {
				return l
			}
			fyne.LogError(keyTypeMismatchError+key, nil)
		}
	}

	listen := &prefBoundBool{key: key, p: p}
	binds = prefBinds.ensurePreferencesAttached(p)
	binds.setItem(key, listen)
	return listen
}

func (b *prefBoundBool) Get() (bool, error) {
	cache := b.p.Bool(b.key)
	b.cacheLock.Lock()
	b.cache = cache
	b.cacheLock.Unlock()
	return b.cache, nil
}

func (b *prefBoundBool) Set(v bool) error {
	b.p.SetBool(b.key, v)

	b.lock.RLock()
	defer b.lock.RUnlock()
	b.trigger()
	return nil
}

func (b *prefBoundBool) checkForChange() {
	b.cacheLock.RLock()
	cache := b.cache
	b.cacheLock.RUnlock()
	if b.p.Bool(b.key) == cache {
		return
	}

	b.trigger()
}

type prefBoundFloat struct {
	base
	key string
	p   fyne.Preferences

	cacheLock sync.RWMutex
	cache     float64
}

// BindPreferenceFloat returns a bindable float64 value that is managed by the application preferences.
// Changes to this value will be saved to application storage and when the app starts the previous values will be read.
//
// Since: 2.0
func BindPreferenceFloat(key string, p fyne.Preferences) Float {
	binds := prefBinds.getBindings(p)
	if binds != nil {
		if listen := binds.getItem(key); listen != nil {
			if l, ok := listen.(Float); ok {
				return l
			}
			fyne.LogError(keyTypeMismatchError+key, nil)
		}
	}

	listen := &prefBoundFloat{key: key, p: p}
	binds = prefBinds.ensurePreferencesAttached(p)
	binds.setItem(key, listen)
	return listen
}

func (b *prefBoundFloat) Get() (float64, error) {
	cache := b.p.Float(b.key)
	b.cacheLock.Lock()
	b.cache = cache
	b.cacheLock.Unlock()
	return b.cache, nil
}

func (b *prefBoundFloat) Set(v float64) error {
	b.p.SetFloat(b.key, v)

	b.lock.RLock()
	defer b.lock.RUnlock()
	b.trigger()
	return nil
}

func (b *prefBoundFloat) checkForChange() {
	b.cacheLock.RLock()
	cache := b.cache
	b.cacheLock.RUnlock()
	if b.p.Float(b.key) == cache {
		return
	}

	b.trigger()
}

type prefBoundInt struct {
	base
	key string
	p   fyne.Preferences

	cacheLock sync.RWMutex
	cache     int
}

// BindPreferenceInt returns a bindable int value that is managed by the application preferences.
// Changes to this value will be saved to application storage and when the app starts the previous values will be read.
//
// Since: 2.0
func BindPreferenceInt(key string, p fyne.Preferences) Int {
	binds := prefBinds.getBindings(p)
	if binds != nil {
		if listen := binds.getItem(key); listen != nil {
			if l, ok := listen.(Int); ok {
				return l
			}
			fyne.LogError(keyTypeMismatchError+key, nil)
		}
	}

	listen := &prefBoundInt{key: key, p: p}
	binds = prefBinds.ensurePreferencesAttached(p)
	binds.setItem(key, listen)
	return listen
}

func (b *prefBoundInt) Get() (int, error) {
	cache := b.p.Int(b.key)
	b.cacheLock.Lock()
	b.cache = cache
	b.cacheLock.Unlock()
	return b.cache, nil
}

func (b *prefBoundInt) Set(v int) error {
	b.p.SetInt(b.key, v)

	b.lock.RLock()
	defer b.lock.RUnlock()
	b.trigger()
	return nil
}

func (b *prefBoundInt) checkForChange() {
	b.cacheLock.RLock()
	cache := b.cache
	b.cacheLock.RUnlock()
	if b.p.Int(b.key) == cache {
		return
	}

	b.trigger()
}

type prefBoundString struct {
	base
	key string
	p   fyne.Preferences

	cacheLock sync.RWMutex
	cache     string
}

// BindPreferenceString returns a bindable string value that is managed by the application preferences.
// Changes to this value will be saved to application storage and when the app starts the previous values will be read.
//
// Since: 2.0
func BindPreferenceString(key string, p fyne.Preferences) String {
	binds := prefBinds.getBindings(p)
	if binds != nil {
		if listen := binds.getItem(key); listen != nil {
			if l, ok := listen.(String); ok {
				return l
			}
			fyne.LogError(keyTypeMismatchError+key, nil)
		}
	}

	listen := &prefBoundString{key: key, p: p}
	binds = prefBinds.ensurePreferencesAttached(p)
	binds.setItem(key, listen)
	return listen
}

func (b *prefBoundString) Get() (string, error) {
	cache := b.p.String(b.key)
	b.cacheLock.Lock()
	b.cache = cache
	b.cacheLock.Unlock()
	return b.cache, nil
}

func (b *prefBoundString) Set(v string) error {
	b.p.SetString(b.key, v)

	b.lock.RLock()
	defer b.lock.RUnlock()
	b.trigger()
	return nil
}

func (b *prefBoundString) checkForChange() {
	b.cacheLock.RLock()
	cache := b.cache
	b.cacheLock.RUnlock()
	if b.p.String(b.key) == cache {
		return
	}

	b.trigger()
}
