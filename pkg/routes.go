package pkg

import (
	"log"
	"net/http"
    "os"
    "time"
    "strings"

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
    
    if (!CheckMonthlyPayment(requestBody.WalletAddress, requestBody.Network)) {
        c.JSON(http.StatusPaymentRequired, gin.H{})
    }
    
    contractText := requestBody.Contract
    tempFileName := requestBody.WalletAddress + "-" + time.Now().Unix() + ".sol"
    //debug contract
    err := WriteStringToFile(tempFileName, contractText)
    if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
    }
    cmd := 
    `
    ./smartbugs -t all -f ` + tempFileName + `
./reparse results
./results2csv -p results > results.csv
    `

    output, err := ExecuteCommand(cmd)
    if err != nil {
	    c.AbortWithError(http.StatusInternalServerError, err)
    }
    
    resp := DebugContractRes{Bugs: strings.Split(output, ",")}
    c.JSON(http.StatusOK, resp)
}

func GenerateContract(c *gin.Context) {
    log.Println("POST /generate-contract has been called")
    requestBody := GenerateContractBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("JSON was missing some required values")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
    
    if (!CheckMonthlyPayment(requestBody.WalletAddress, requestBody.Network)) {
        c.JSON(http.StatusPaymentRequired, gin.H{})
    }
    
    prompt= requestBody.Prompt
    cmd := "python run_model.py " + prompt

    output, err := ExecuteCommand(cmd)
    if err != nil {
	    c.AbortWithError(http.StatusInternalServerError, err)
    }
    
    resp := GenerateContractRes{Contract: output}
    c.JSON(http.StatusOK, resp)
}
