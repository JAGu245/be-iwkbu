package main

import (
	"fmt"
	"jm-CICO/config"
	"jm-CICO/handlers"
	"jm-CICO/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	config.InitDB()
	handlers.SeedUser()

	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/user/update", utils.AuthMiddleware(handlers.UpdateUserHandler))
	http.HandleFunc("/api/user/register", utils.AuthMiddleware(handlers.RegisterHandler))
	http.HandleFunc("/api/users", utils.AuthMiddleware(handlers.GetUsersHandler))
	http.HandleFunc("/api/user/delete", utils.AuthMiddleware(handlers.DeleteUserHandler))

	http.HandleFunc("/loketcabangjawatengah", utils.AuthMiddleware(handlers.LoketCabangHandler))      //Oke
	http.HandleFunc("/samsatkendal", utils.AuthMiddleware(handlers.SamsatKendalHandler))              //Oke
	http.HandleFunc("/samsatdemak", utils.AuthMiddleware(handlers.SamsatDemakHandler))                //Oke
	http.HandleFunc("/samsatpurwodadi", utils.AuthMiddleware(handlers.SamsatPurwodadiHandler))        //Oke
	http.HandleFunc("/samsatungaran", utils.AuthMiddleware(handlers.SamsatUngaranHandler))            //Oke
	http.HandleFunc("/samsatsalatiga", utils.AuthMiddleware(handlers.SamsatSalatigaHandler))          //Oke
	http.HandleFunc("/samsatlokperwsra", utils.AuthMiddleware(handlers.SamsatLokPerwSRAHandler))      //Oke
	http.HandleFunc("/samsatsurakarta", utils.AuthMiddleware(handlers.SamsatSurakartaHandler))        //Oke
	http.HandleFunc("/samsatklaten", utils.AuthMiddleware(handlers.SamsatKlatenHandler))              //Oke
	http.HandleFunc("/samsatsragen", utils.AuthMiddleware(handlers.SamsatSragenHandler))              //Oke
	http.HandleFunc("/samsatboyolali", utils.AuthMiddleware(handlers.SamsatBoyolaliHandler))          //Oke
	http.HandleFunc("/samsatprambanan", utils.AuthMiddleware(handlers.SamsatPrambananHandler))        //Oke
	http.HandleFunc("/samsatdelanggu", utils.AuthMiddleware(handlers.SamsatDelangguHandler))          //Oke
	http.HandleFunc("/samsatlokpwkmgl", utils.AuthMiddleware(handlers.SamsatLokPwkMGLHandler))        //Oke
	http.HandleFunc("/samsatmagelang", utils.AuthMiddleware(handlers.SamsatMagelangHandler))          //Oke
	http.HandleFunc("/samsatpurworejo", utils.AuthMiddleware(handlers.SamsatPurworejoHandler))        //Oke
	http.HandleFunc("/samsatkebumen", utils.AuthMiddleware(handlers.SamsatKebumenHandler))            //oke
	http.HandleFunc("/samsattemanggung", utils.AuthMiddleware(handlers.SamsatTemanggungHandler))      //Oke
	http.HandleFunc("/samsatwonosobo", utils.AuthMiddleware(handlers.SamsatWonosoboHandler))          //oke
	http.HandleFunc("/samsatmungkid", utils.AuthMiddleware(handlers.SamsatMungkidHandler))            //Oke
	http.HandleFunc("/samsatbagelen", utils.AuthMiddleware(handlers.SamsatBagelenHandler))            //Oke
	http.HandleFunc("/samsatlokprwpwt", utils.AuthMiddleware(handlers.SamsatLokPrwPWTHandler))        //Oke
	http.HandleFunc("/samsat/purwokerto", utils.AuthMiddleware(handlers.SamsatPurwokertoHandler))     //Oke
	http.HandleFunc("/samsat/purbalingga", utils.AuthMiddleware(handlers.SamsatPurbalinggaHandler))   //oke
	http.HandleFunc("/samsat/banjarnegara", utils.AuthMiddleware(handlers.SamsatBanjarnegaraHandler)) //Oke
	http.HandleFunc("/samsat/majenang", utils.AuthMiddleware(handlers.SamsatMajenangHandler))         //Oke
	http.HandleFunc("/samsat/cilacap", utils.AuthMiddleware(handlers.SamsatCilacapHandler))           //Oke
	http.HandleFunc("/samsat/wangon", utils.AuthMiddleware(handlers.SamsatWangonHandler))             //Oke
	http.HandleFunc("/samsat/lokprwpkl", utils.AuthMiddleware(handlers.SamsatLokPrwPKLHandler))       //oke
	http.HandleFunc("/samsat/pekalongan", utils.AuthMiddleware(handlers.SamsatPekalonganHandler))     //Oke
	http.HandleFunc("/samsat/pemalang", utils.AuthMiddleware(handlers.SamsatPemalangHandler))         //Oke
	http.HandleFunc("/samsat/tegal", utils.AuthMiddleware(handlers.SamsatTegalHandler))               //Oke
	http.HandleFunc("/samsat/brebes", utils.AuthMiddleware(handlers.SamsatBrebesHandler))             //Oke
	http.HandleFunc("/samsat/batang", utils.AuthMiddleware(handlers.SamsatBatangHandler))             //Oke
	http.HandleFunc("/samsat/kajen", utils.AuthMiddleware(handlers.SamsatKajenHandler))               //Oke
	http.HandleFunc("/samsat/slawi", utils.AuthMiddleware(handlers.SamsatSlawiHandler))               //Oke
	http.HandleFunc("/samsat/bumiayu", utils.AuthMiddleware(handlers.SamsatBumiayuHandler))           //Oke
	http.HandleFunc("/samsat/tanjung", utils.AuthMiddleware(handlers.SamsatTanjungHandler))           //Oke
	http.HandleFunc("/samsat/lokprwpti", utils.AuthMiddleware(handlers.SamsatLokPrwPTIHandler))       //Oke
	http.HandleFunc("/samsat/pati", utils.AuthMiddleware(handlers.SamsatPatiHandler))                 //Oke
	http.HandleFunc("/samsat/kudus", utils.AuthMiddleware(handlers.SamsatKudusHandler))               //Oke
	http.HandleFunc("/samsat/jepara", utils.AuthMiddleware(handlers.SamsatJeparaHandler))             //Oke
	http.HandleFunc("/samsat/rembang", utils.AuthMiddleware(handlers.SamsatRembangHandler))           //Oke
	http.HandleFunc("/samsat/blora", utils.AuthMiddleware(handlers.SamsatBloraHandler))               //Oke
	http.HandleFunc("/samsat/cepu", utils.AuthMiddleware(handlers.SamsatCepuHandler))                 //Oke
	http.HandleFunc("/samsat/lokprwsmg", utils.AuthMiddleware(handlers.SamsatLokPrwSMGHandler))       //Oke
	http.HandleFunc("/samsat/semarang1", utils.AuthMiddleware(handlers.SamsatSemarang1Handler))       //Oke
	http.HandleFunc("/samsat/semarang2", utils.AuthMiddleware(handlers.SamsatSemarang2Handler))       //Oke
	http.HandleFunc("/samsat/semarang3", utils.AuthMiddleware(handlers.SamsatSemarang3Handler))       //Oke
	http.HandleFunc("/samsat/lokprwskh", utils.AuthMiddleware(handlers.SamsatLokPrwSKHHandler))       //Oke
	http.HandleFunc("/samsat/sukoharjo", utils.AuthMiddleware(handlers.SamsatSukoharjoHandler))       //Oke
	http.HandleFunc("/samsat/karanganyar", utils.AuthMiddleware(handlers.SamsatKaranganyarHandler))   //Oke
	http.HandleFunc("/samsat/wonogiri", utils.AuthMiddleware(handlers.SamsatWonogiriHandler))         //Oke
	http.HandleFunc("/samsat/purwantoro", utils.AuthMiddleware(handlers.SamsatPurwantoroHandler))     //Oke
	http.HandleFunc("/samsat/baturetno", utils.AuthMiddleware(handlers.SamsatBaturetnoHandler))       //Oke

	// Data Rekap CICO 3M per Cabang

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
