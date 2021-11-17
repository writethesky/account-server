package api

import (
	"account-server/internal"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// SetSecureKey godoc
// @Summary Set SecureKey
// @Schemes
// @Tags secures
// @Accept json
// @Produce json
// @Success 200 {object} string "session id"
// @Failure 500 {object} string "error"
// @Router /secures [post]
func SetSecureKey(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	block, _ := pem.Decode([]byte(internal.Config.Server.PrivateKey))
	if block == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Provided data is not private key")
		return
	}
	if len(block.Bytes) < 100 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Provided data is not private key")
		return
	}

	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	privateKey := privateKeyInterface.(*rsa.PrivateKey)

	transferKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, bodyBytes, nil)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	sessionID := uuid.NewString()
	internal.RDB.Set(c, getRedisSessionKey(sessionID), transferKey, time.Hour*24)
	c.JSON(http.StatusOK, sessionID)
}

func Secure(c *gin.Context) {
	sessionID := c.Request.Header.Get("Security")
	if sessionID == "" {
		return
	}
	nonce := c.Request.Header.Get("Nonce")
	iv, err := hex.DecodeString(nonce)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(iv) != 12 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "nonce bytes length must be 12")
		return
	}

	transferKey, err := internal.RDB.Get(c, getRedisSessionKey(sessionID)).Bytes()
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	internal.RDB.Set(c, getRedisSessionKey(sessionID), transferKey, time.Hour*24)

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(bodyBytes) != 0 {
		body, err := aesGCMDecrypt(transferKey, iv, bodyBytes)
		if nil != err {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
	}

	tmpWriter := &bodyWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = tmpWriter

	c.Next()

	bodyBytes, err = aesGCMEncrypt(transferKey, iv, tmpWriter.body.Bytes())
	if nil != err {
		tmpWriter.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = tmpWriter.ResponseWriter.WriteString(err.Error())
		return
	}

	_, _ = tmpWriter.ResponseWriter.Write(bodyBytes)

}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return len(b), nil
}

func getRedisSessionKey(sessionID string) string {
	return fmt.Sprintf("session.%s", sessionID)
}

func aesGCMDecrypt(key, iv, cypherBody []byte) (body []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	body, err = aesGCM.Open(nil, iv, cypherBody, nil)
	return
}

func aesGCMEncrypt(key, iv, body []byte) (cypherBody []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	cypherBody = aesGCM.Seal(nil, iv, body, nil)
	return
}
