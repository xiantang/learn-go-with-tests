package dictionary

//使错误更具可重用性和不可变性。
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exists")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// map 是引用类型 不需要作为引用就可以完成
func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {

	// map 有一个内置函数 delete。它需要两个参数。第一个是这个 map，第二个是要删除的键。
	delete(d, key)

}
