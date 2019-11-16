package registryutil

import (
	"internal/syscall/windows/registry"
	"log"
)

func SetregistryString(sRegKey, skey, sValue string) bool {

	key, exists, err := registry.CreateKey(registry.CURRENT_USER, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	err = key.SetStringValue(skey, sValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryString(sRegKey, skey string) (string, error) {

	key, err := registry.OpenKey(registry.CURRENT_USER, sRegKey, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetStringValue(skey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return val, nil
}

func SetregistryDWord(sRegKey string, skey string, nValue int) bool {

	key, exists, err := registry.CreateKey(registry.CURRENT_USER, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	dValue := uint32(nValue)
	err = key.SetDWordValue(skey, dValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryDWord(sRegKey string, skey string) (int, error) {

	key, err := registry.OpenKey(registry.CURRENT_USER, sRegKey, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue(skey)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(val), nil
}

func DeleteKeyregistry(sRegKey string) error {

	err := registry.DeleteKey(registry.CURRENT_USER, sRegKey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteValueregistry(sRegKey string, skey string) error {

	key, err := registry.OpenKey(registry.CURRENT_USER, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	err = key.DeleteValue(skey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func SetregistryStringWOW64(sRegKey, skey, sValue string) bool {

	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, sRegKey, registry.WOW64_64KEY|registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	err = key.SetStringValue(skey, sValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryStringWOW64(sRegKey, skey string) (string, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, sRegKey, registry.WOW64_64KEY|registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetStringValue(skey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return val, nil
}

func SetregistryDWordWOW64(sRegKey string, skey string, nValue int) bool {

	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, sRegKey, registry.WOW64_64KEY|registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	dValue := uint32(nValue)
	err = key.SetDWordValue(skey, dValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryDWordWOW64(sRegKey string, skey string) (int, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, sRegKey, registry.WOW64_64KEY|registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue(skey)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(val), nil
}

func SetregistryString32(sRegKey, skey, sValue string) bool {

	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	err = key.SetStringValue(skey, sValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryString32(sRegKey, skey string) (string, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, sRegKey, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetStringValue(skey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return val, nil
}

func SetregistryDWord32(sRegKey string, skey string, nValue int) bool {

	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		//fmt.Println(sRegKey + "exists")
	}

	dValue := uint32(nValue)
	err = key.SetDWordValue(skey, dValue)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetregistryDWord32(sRegKey string, skey string) (int, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, sRegKey, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue(skey)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(val), nil
}

func DeleteKeyregistry32(sRegKey string) error {

	err := registry.DeleteKey(registry.LOCAL_MACHINE, sRegKey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteValueregistry32(sRegKey string, skey string) error {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, sRegKey, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	err = key.DeleteValue(skey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
