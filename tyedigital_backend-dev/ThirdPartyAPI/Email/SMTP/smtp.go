package email

// import (
// 	"bytes"
// 	"crypto/tls"
// 	"errors"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net"
// 	"net/smtp"
// 	"strconv"

// 	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
// 	header "github.com/Kushmanda-Tech/tyedigital_backend/Middleware/Header"
// 	model "github.com/Kushmanda-Tech/tyedigital_backend/Model"
// 	util "github.com/Kushmanda-Tech/tyedigital_backend/Util"

// 	"github.com/gin-gonic/gin"
// )

// type loginAuth struct {
// 	username, password string
// }

// func LoginAuth(username, password string) smtp.Auth {
// 	return &loginAuth{username, password}
// }

// func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
// 	return "LOGIN", []byte(a.username), nil
// }

// func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
// 	if more {
// 		switch string(fromServer) {
// 		case "Username:":
// 			return []byte(a.username), nil
// 		case "Password:":
// 			return []byte(a.password), nil
// 		default:
// 			return nil, errors.New("Unknown from server")
// 		}
// 	}
// 	return nil, nil
// }

// func SendEmailForVerifyingEmail(users []model.StudentDetails, ginC *gin.Context) {
// 	if users[0].Verify_Email == "Unverified" {
// 		// Sender data.
// 		from := config.EMAIL_SEND_FROM
// 		password := config.NO_REPLY_EMAIL_PASSWORS

// 		// Receiver email address.
// 		to := []string{
// 			users[0].Email,
// 		}

// 		// smtp server configuration.
// 		smtpHost := config.SMTP_SERVER_ADDRESS
// 		smtpPort := config.SMTP_PORT

// 		conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
// 		if err != nil {
// 			log.Println("Failed to connect to the SMTP server:", err)
// 			ginC.JSON(200, gin.H{
// 				"status":  451,
// 				"message": "Failed to connect to the SMTP server",
// 				"error":   err.Error(),
// 			})
// 			return
// 		}

// 		c, err := smtp.NewClient(conn, smtpHost)
// 		if err != nil {
// 			log.Println("Failed to create SMTP client:", err)
// 			ginC.JSON(200, gin.H{
// 				"status":  452,
// 				"message": "Failed to create SMTP client",
// 				"error":   err.Error(),
// 			})
// 			return
// 		}

// 		tlsconfig := &tls.Config{
// 			ServerName: smtpHost,
// 		}

// 		if err = c.StartTLS(tlsconfig); err != nil {
// 			log.Println("Failed to start TLS connection:", err)
// 			ginC.JSON(200, gin.H{
// 				"status":  453,
// 				"message": "Failed to start TLS connection",
// 				"error":   err.Error(),
// 			})
// 			return
// 		}

// 		auth := LoginAuth(from, password)

// 		if err = c.Auth(auth); err != nil {
// 			log.Println("Failed to authenticate with SMTP server:", err)
// 			ginC.JSON(200, gin.H{
// 				"status":  455,
// 				"message": "Failed to authenticate with SMTP server",
// 				"error":   err.Error(),
// 			})
// 			return
// 		}

// 		t, _ := template.ParseFiles("./ThirdPartyAPI/Email/templates/resend_email.html")

// 		var body bytes.Buffer

// 		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 		body.Write([]byte(fmt.Sprintf("Subject: Verify Email \n%s\n\n", mimeHeaders)))

// 		t.Execute(&body, struct {
// 			Name string
// 			Url  string
// 		}{
// 			Name: users[0].Name,
// 			Url:  config.VERIFY_EMAIL_API + users[0].Email,
// 		})

// 		// Sending email.
// 		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// 		if err != nil {
// 			log.Println("Failed to send email:", err)
// 			ginC.JSON(200, gin.H{
// 				"status":  501,
// 				"message": "Failed to send email",
// 				"error":   err.Error(),
// 			})
// 			return
// 		}
// 		log.Println("Email Sent Sucessfully!")
// 		ginC.JSON(200, gin.H{
// 			"status":  200,
// 			"message": "Email Sent Sucessfully!",
// 		})
// 	} else {
// 		ginC.JSON(200, gin.H{
// 			"status":  452,
// 			"message": "Your email already verified",
// 		})
// 	}
// }

// func SendEmailForOTP(users []model.GetUsers, ginC *gin.Context) {

// 	// Sender data.
// 	from := config.EMAIL_SEND_FROM
// 	password := config.NO_REPLY_EMAIL_PASSWORS

// 	// Receiver email address.
// 	to := []string{
// 		users[0].Email,
// 	}

// 	// smtp server configuration.
// 	smtpHost := config.SMTP_SERVER_ADDRESS
// 	smtpPort := config.SMTP_PORT

// 	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
// 	if err != nil {
// 		log.Println("Failed to connect to the SMTP server:", err)
// 		ginC.JSON(200, gin.H{
// 			"status":  451,
// 			"message": "Failed to connect to the SMTP server",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c, err := smtp.NewClient(conn, smtpHost)
// 	if err != nil {
// 		log.Println("Failed to create SMTP client:", err)
// 		ginC.JSON(200, gin.H{
// 			"status":  452,
// 			"message": "Failed to create SMTP client",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	tlsconfig := &tls.Config{
// 		ServerName: smtpHost,
// 	}

// 	if err = c.StartTLS(tlsconfig); err != nil {
// 		log.Println("Failed to start TLS connection:", err)
// 		ginC.JSON(200, gin.H{
// 			"status":  453,
// 			"message": "Failed to start TLS connection",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	auth := LoginAuth(from, password)

// 	if err = c.Auth(auth); err != nil {
// 		log.Println("Failed to authenticate with SMTP server:", err)
// 		ginC.JSON(200, gin.H{
// 			"status":  455,
// 			"message": "Failed to authenticate with SMTP server",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	t, _ := template.ParseFiles("./ThirdPartyAPI/Email/templates/send_otp_to_email.html")

// 	var body bytes.Buffer

// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body.Write([]byte(fmt.Sprintf("Subject: OTP \n%s\n\n", mimeHeaders)))

// 	t.Execute(&body, struct {
// 		Name string
// 		Otp  int
// 	}{
// 		Name: users[0].Name,
// 		Otp:  util.GenerateSixDigitesRandomNumber(),
// 	})

// 	// Sending email.
// 	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// 	if err != nil {
// 		log.Println("Failed to send email:", err)
// 		ginC.JSON(200, gin.H{
// 			"status":  501,
// 			"message": "Failed to send OTP via email",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// 	log.Println("Email Sent Sucessfully!")
// 	ginC.JSON(200, gin.H{
// 		"status":  200,
// 		"data":    users,
// 		"message": "Otp sent successfully",
// 		"token":   header.CreateToken(strconv.Itoa(users[0].User_Id), config.Input.Email),
// 	})
// 	return
// }
