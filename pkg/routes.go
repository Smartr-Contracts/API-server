package pkg

import (
	"log"
	"net/http"
    "os"
    "time"

     "github.com/gin-gonic/gin"
     "github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
    log.Println("POST /login has been called")
    requestBody := LoginBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("JSON was missing some required values")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

    //query contract for the month payment
    paid := true

    if paid == false {
        c.JSON(http.StatusPaymentRequired, gin.H{})
        return
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": requestBody.WalletAddress,
        "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
    })
    
    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

    if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
        return
 }

 // Respond
    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

    c.JSON(http.StatusOK, gin.H{})
}

func DebugContract(c *gin.Context) {
    log.Println("POST /debug-contract has been called")
    requestBody := DebugContractBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("JSON was missing some required values")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
    //debug contract
    c.JSON(http.StatusOK, gin.H{})
}

func GenerateContract(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context){
    user,_ := c.Get("user")

    // user.(models.User).Email    -->   to access specific data

    c.JSON(http.StatusOK, gin.H{
        "message": user,
    })
}
