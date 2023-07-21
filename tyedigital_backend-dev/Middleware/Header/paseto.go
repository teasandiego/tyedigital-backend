package header

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"

	paseto "aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
)

func CreateToken(UUID, email string) string {
	token := paseto.NewToken()
	token.SetAudience(config.Audience)
	token.SetJti(config.Identifier)
	token.SetIssuer(config.Issuer)
	token.SetSubject(config.Subject)
	token.SetExpiration(time.Now().Add(time.Hour))
	token.SetNotBefore(time.Now())
	token.SetIssuedAt(time.Now())
	token.SetString("user_id", UUID)
	token.SetString("email", email)

	secretKey, _ := paseto.NewV4AsymmetricSecretKeyFromHex(config.SecretKeyHex)

	signed := token.V4Sign(secretKey, nil)
	return signed
}

func VerifyPasetoTokenWithOutExpiry(token string) (bool, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.ForAudience(config.Audience))
	parser.AddRule(paseto.IdentifiedBy(config.Identifier))
	parser.AddRule(paseto.IssuedBy(config.Issuer))
	parser.AddRule(paseto.Subject(config.Subject))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(config.PublicKeyHex) // this wil fail if given key in an invalid format
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	parserWithOutExpiry := paseto.NewParserWithoutExpiryCheck() // only used because this example token has expired,
	token1, parsedTokenWithoutExpiryErr := parserWithOutExpiry.ParseV4Public(publicKey, token, nil)
	if parsedTokenWithoutExpiryErr != nil {
		log.Println("parsedTokenWithoutExpiryErr" + parsedTokenWithoutExpiryErr.Error())
		return false, err
	}
	Jobject := string(token1.ClaimsJSON())
	var data model.PasetoClaimData
	err123 := json.Unmarshal([]byte(Jobject), &data)
	if err123 != nil {
		panic(err123)
	}
	return true, nil
}
func VarifyTokenPaseto(token string) (bool, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.ForAudience(config.Audience))
	parser.AddRule(paseto.IdentifiedBy(config.Identifier))
	parser.AddRule(paseto.IssuedBy(config.Issuer))
	parser.AddRule(paseto.Subject(config.Subject))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))
	// parser.AddRule(paseto.SetString(token))

	publicKey, publicKeyErr := paseto.NewV4AsymmetricPublicKeyFromHex(config.PublicKeyHex) // this wil fail if given key in an invalid format
	if publicKeyErr != nil {
		log.Println("publicKey" + publicKeyErr.Error())
		return false, publicKeyErr
	}
	token1, parserTokenWithEpiryErr := parser.ParseV4Public(publicKey, token, nil)
	if parserTokenWithEpiryErr != nil {
		log.Println("parserTokenWithEpiryErr " + parserTokenWithEpiryErr.Error())
		return false, parserTokenWithEpiryErr
	}
	log.Println(string(token1.ClaimsJSON()))
	return true, nil
}
func DecryptPasetoToken(c *gin.Context) (string, string, error) {
	authorizationToken := c.Request.Header.Get("Authorization")
	if len(authorizationToken) == 0 {
		c.JSON(config.FALIURE, gin.H{
			"status":  config.FALIURE,
			"message": "Your Authorization token is empty",
		})
	}
	parser := paseto.NewParser()
	parser.AddRule(paseto.ForAudience(config.Audience))
	parser.AddRule(paseto.IdentifiedBy(config.Identifier))
	parser.AddRule(paseto.IssuedBy(config.Issuer))
	parser.AddRule(paseto.Subject(config.Subject))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(config.PublicKeyHex) // this wil fail if given key in an invalid format
	if err != nil {
		log.Println(err.Error())
		return "", "", err
	}
	parserWithOutExpiry := paseto.NewParserWithoutExpiryCheck() // only used because this example token has expired,
	token1, parsedTokenWithoutExpiryErr := parserWithOutExpiry.ParseV4Public(publicKey, authorizationToken, nil)
	if parsedTokenWithoutExpiryErr != nil {
		log.Println("parsedTokenWithoutExpiryErr" + parsedTokenWithoutExpiryErr.Error())
		return "", "", err
	}
	Jobject := string(token1.ClaimsJSON())
	var data model.PasetoClaimData
	err123 := json.Unmarshal([]byte(Jobject), &data)
	if err123 != nil {
		log.Println(err123)
	}
	return data.Email, data.UserID, err

}
func PasetoEncrypt(email, UUID string) string {
	token := paseto.NewToken()
	token.SetAudience(config.Audience)
	token.SetJti(config.Identifier)
	token.SetIssuer(config.Issuer)
	token.SetSubject(config.Subject)
	token.SetExpiration(time.Now().Add(time.Hour))
	token.SetNotBefore(time.Now())
	token.SetIssuedAt(time.Now())
	token.SetString("user_id", UUID)
	token.SetString("email", email)
	key := paseto.NewV4SymmetricKey()

	encrypted := paseto.NewToken().V4Encrypt(key, nil)
	log.Println(encrypted)
	return encrypted
}

func TokenAuth() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")

		vTokenWithExpiry, vTokenWithExpiryErr := VarifyTokenPaseto(Authorization)
		vTokenWithOutExpiry, _ := VerifyPasetoTokenWithOutExpiry(Authorization)

		if len(Authorization) == 0 {
			c.JSON(config.UNAUTHORIZE, gin.H{
				"message": "Please give Authorization in headers",
				"status":  config.UNAUTHORIZE,
			})
			c.AbortWithStatus(http.StatusOK)
		} else if len(Authorization) != 0 {
			if vTokenWithExpiry && vTokenWithOutExpiry {
				c.JSON(config.AUTHORIZE, gin.H{
					"message": "Your are authorize persion",
					"status":  config.AUTHORIZE,
				})
			} else if vTokenWithOutExpiry && !vTokenWithExpiry {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"message": "Sorry your token is expired",
					"error":   vTokenWithExpiryErr,
					"status":  config.UNAUTHORIZE,
				})
			} else if !vTokenWithExpiry {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"message": "You are Unauthorized",
					"error":   vTokenWithExpiryErr,
					"status":  config.UNAUTHORIZE,
				})
			}
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
	return gin.HandlerFunc(fn)
}
