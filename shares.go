package main 

import (
  "fmt"
  "flag"
  "log"
  "os"
	"crypto/rand"
  mathRand "math/rand"
  "encoding/base64"
	"github.com/hashicorp/vault/shamir"
)


func main() {
    if err := root(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func root() error {
    // Args
    total_keys_flag := flag.Int("k", 0, "# of keys")
    threshold_keys_flag := flag.Int("t", 0, "Threshold # of keys to recreate")
    flag_key_length := flag.Int("l", 64, "Length of generated key")
    flag.Parse()

    total_keys := *total_keys_flag	
    threshold_keys := *threshold_keys_flag

    if total_keys < threshold_keys {
        log.Fatal("Total keys less than rebuild keys")
    }

    // Gen key + shared
    rand_bytes, err := generateKey(*flag_key_length)
    if err != nil { return err }
    base64_key := base64.StdEncoding.EncodeToString(rand_bytes)
    fmt.Printf("random key[%v]\n%s\n", *flag_key_length, base64_key)
    fmt.Printf("generating keys[%v] threshold[%v]\n", total_keys, threshold_keys)
    keys, err := splitBytes(rand_bytes, total_keys, threshold_keys)
    if err != nil { return err }

    for index, key := range keys {
        base64_key := base64.StdEncoding.EncodeToString(key)
        fmt.Printf("key[%v]: %s\n", index, base64_key)
    }

    // Check what was generate
    for i := threshold_keys; i <= total_keys; i++ {
        rebuildKey(keys, i) 
    }
    return nil
}

func generateKey(key_size int) ([]byte, error) {
    rand_bytes := make([]byte, key_size)
    _, err := rand.Read(rand_bytes)
    return rand_bytes, err
}

func splitBytes(input []byte, number_of_keys int, threshold_number_of_keys int) ([][]byte, error) {
    return shamir.Split(input, number_of_keys, threshold_number_of_keys)
}

func rebuildBytes(inputs [][]byte) ([]byte, error){
    return shamir.Combine(inputs)
}

func rebuildKey(inputs [][]byte, count int) ([]byte, error){
    keys_to_rebuild := [][]byte{}
    perm := mathRand.Perm(len(inputs))
    for _, v := range perm[:count] {
        keys_to_rebuild = append(keys_to_rebuild, inputs[v])
    }
    rebuilt_key, err := rebuildBytes(keys_to_rebuild) 
    if err != nil { return []byte{}, err }
    rebuilt_key_base64 := base64.StdEncoding.EncodeToString(rebuilt_key)
    fmt.Printf("Using [%v] keys %v\n%s\n", len(keys_to_rebuild), perm[:count], rebuilt_key_base64)
    return rebuilt_key, nil
}

