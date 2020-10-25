package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	//"data-manager/crypto"
	"encoding/pem"
)

func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	//publicKeyStr = bufferPublic.String()
	//log.Debug("-------------公钥----------------")
	//log.Debug("\r", publicKeyStr)
	//log.Debug("--------------私钥---------------")
	//log.Debug("\r", privateKeyStr)
	return privateKeyStr, publicKeyStr, nil

}

//func Genarate() (address, public crypto2.PublicKey, private ecdsa.PrivateKey) {
//
//	var privateKey *ecdsa.PrivateKey
//	var err error
//	// generate random.
//	privateKey1, err := crypto.GenerateKey()
//	if err != nil {
//		logrus.Debug("Failed to generate random private key: %v", err)
//	}
//
//	// Output some information.
//	address = privateKey1.PublicKey
//	publicKey := privateKey1.PublicKey
//	privateKey = privateKey1
//	return address, publicKey, *privateKey
//		//Address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
//		//PublicKey :=  hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])
//		//PrivateKey := hex.EncodeToString(crypto.FromECDSA(privateKey))
//
//	//if ctx.Bool(jsonFlag.Name) {
//	//	mustPrintJSON(out)
//	//} else {
//	//	fmt.Println("Address   : ", out.Address)
//	//	fmt.Println("PrivateKey: ", out.PrivateKey)
//	//	fmt.Println("PublicKey : ", out.PublicKey)
//	//}
//}
//func (a Address) Hex() string {
//	unchecksummed := hex.EncodeToString(a[:])
//	sha := sha3.NewKeccak256()
//	sha.Write([]byte(unchecksummed))
//	hash := sha.Sum(nil)
//
//	result := []byte(unchecksummed)
//	for i := 0; i < len(result); i++ {
//		hashByte := hash[i/2]
//		if i%2 == 0 {
//			hashByte = hashByte >> 4
//		} else {
//			hashByte &= 0xf
//		}
//		if result[i] > '9' && hashByte > 7 {
//			result[i] -= 32
//		}
//	}
//	return "0x" + string(result)
//}
