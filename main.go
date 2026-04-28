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

	http.HandleFunc("/loketcabangjawatengah", utils.AuthMiddleware(utils.CacheMiddleware(handlers.LoketCabangHandler)))      //Oke
	http.HandleFunc("/samsatkendal", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKendalHandler)))              //Oke
	http.HandleFunc("/samsatdemak", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatDemakHandler)))                //Oke
	http.HandleFunc("/samsatpurwodadi", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPurwodadiHandler)))        //Oke
	http.HandleFunc("/samsatungaran", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatUngaranHandler)))            //Oke
	http.HandleFunc("/samsatsalatiga", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSalatigaHandler)))          //Oke
	http.HandleFunc("/samsatlokperwsra", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPerwSRAHandler)))      //Oke
	http.HandleFunc("/samsatsurakarta", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSurakartaHandler)))        //Oke
	http.HandleFunc("/samsatklaten", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKlatenHandler)))              //Oke
	http.HandleFunc("/samsatsragen", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSragenHandler)))              //Oke
	http.HandleFunc("/samsatboyolali", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBoyolaliHandler)))          //Oke
	http.HandleFunc("/samsatprambanan", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPrambananHandler)))        //Oke
	http.HandleFunc("/samsatdelanggu", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatDelangguHandler)))          //Oke
	http.HandleFunc("/samsatlokpwkmgl", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPwkMGLHandler)))        //Oke
	http.HandleFunc("/samsatmagelang", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatMagelangHandler)))          //Oke
	http.HandleFunc("/samsatpurworejo", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPurworejoHandler)))        //Oke
	http.HandleFunc("/samsatkebumen", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKebumenHandler)))            //oke
	http.HandleFunc("/samsattemanggung", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatTemanggungHandler)))      //Oke
	http.HandleFunc("/samsatwonosobo", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatWonosoboHandler)))          //oke
	http.HandleFunc("/samsatmungkid", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatMungkidHandler)))            //Oke
	http.HandleFunc("/samsatbagelen", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBagelenHandler)))            //Oke
	http.HandleFunc("/samsatlokprwpwt", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPrwPWTHandler)))        //Oke
	http.HandleFunc("/samsat/purwokerto", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPurwokertoHandler)))     //Oke
	http.HandleFunc("/samsat/purbalingga", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPurbalinggaHandler)))   //oke
	http.HandleFunc("/samsat/banjarnegara", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBanjarnegaraHandler))) //Oke
	http.HandleFunc("/samsat/majenang", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatMajenangHandler)))         //Oke
	http.HandleFunc("/samsat/cilacap", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatCilacapHandler)))           //Oke
	http.HandleFunc("/samsat/wangon", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatWangonHandler)))             //Oke
	http.HandleFunc("/samsat/lokprwpkl", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPrwPKLHandler)))       //oke
	http.HandleFunc("/samsat/pekalongan", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPekalonganHandler)))     //Oke
	http.HandleFunc("/samsat/pemalang", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPemalangHandler)))         //Oke
	http.HandleFunc("/samsat/tegal", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatTegalHandler)))               //Oke
	http.HandleFunc("/samsat/brebes", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBrebesHandler)))             //Oke
	http.HandleFunc("/samsat/batang", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBatangHandler)))             //Oke
	http.HandleFunc("/samsat/kajen", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKajenHandler)))               //Oke
	http.HandleFunc("/samsat/slawi", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSlawiHandler)))               //Oke
	http.HandleFunc("/samsat/bumiayu", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBumiayuHandler)))           //Oke
	http.HandleFunc("/samsat/tanjung", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatTanjungHandler)))           //Oke
	http.HandleFunc("/samsat/lokprwpti", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPrwPTIHandler)))       //Oke
	http.HandleFunc("/samsat/pati", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPatiHandler)))                 //Oke
	http.HandleFunc("/samsat/kudus", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKudusHandler)))               //Oke
	http.HandleFunc("/samsat/jepara", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatJeparaHandler)))             //Oke
	http.HandleFunc("/samsat/rembang", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatRembangHandler)))           //Oke
	http.HandleFunc("/samsat/blora", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBloraHandler)))               //Oke
	http.HandleFunc("/samsat/cepu", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatCepuHandler)))                 //Oke
	http.HandleFunc("/samsat/lokprwsmg", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPrwSMGHandler)))       //Oke
	http.HandleFunc("/samsat/semarang1", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSemarang1Handler)))       //Oke
	http.HandleFunc("/samsat/semarang2", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSemarang2Handler)))       //Oke
	http.HandleFunc("/samsat/semarang3", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSemarang3Handler)))       //Oke
	http.HandleFunc("/samsat/lokprwskh", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatLokPrwSKHHandler)))       //Oke
	http.HandleFunc("/samsat/sukoharjo", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatSukoharjoHandler)))       //Oke
	http.HandleFunc("/samsat/karanganyar", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatKaranganyarHandler)))   //Oke
	http.HandleFunc("/samsat/wonogiri", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatWonogiriHandler)))         //Oke
	http.HandleFunc("/samsat/purwantoro", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatPurwantoroHandler)))     //Oke
	http.HandleFunc("/samsat/baturetno", utils.AuthMiddleware(utils.CacheMiddleware(handlers.SamsatBaturetnoHandler)))       //Oke

	// Data Rekap CICO 3M per Cabang

	endpoints := []string{
		"/loketcabangjawatengah", "/samsatkendal", "/samsatdemak", "/samsatpurwodadi",
		"/samsatungaran", "/samsatsalatiga", "/samsatlokperwsra", "/samsatsurakarta",
		"/samsatklaten", "/samsatsragen", "/samsatboyolali", "/samsatprambanan",
		"/samsatdelanggu", "/samsatlokpwkmgl", "/samsatmagelang", "/samsatpurworejo",
		"/samsatkebumen", "/samsattemanggung", "/samsatwonosobo", "/samsatmungkid",
		"/samsatbagelen", "/samsatlokprwpwt", "/samsat/purwokerto", "/samsat/purbalingga",
		"/samsat/banjarnegara", "/samsat/majenang", "/samsat/cilacap", "/samsat/wangon",
		"/samsat/lokprwpkl", "/samsat/pekalongan", "/samsat/pemalang", "/samsat/tegal",
		"/samsat/brebes", "/samsat/batang", "/samsat/kajen", "/samsat/slawi",
		"/samsat/bumiayu", "/samsat/tanjung", "/samsat/lokprwpti", "/samsat/pati",
		"/samsat/kudus", "/samsat/jepara", "/samsat/rembang", "/samsat/blora",
		"/samsat/cepu", "/samsat/lokprwsmg", "/samsat/semarang1", "/samsat/semarang2",
		"/samsat/semarang3", "/samsat/lokprwskh", "/samsat/sukoharjo", "/samsat/karanganyar",
		"/samsat/wonogiri", "/samsat/purwantoro", "/samsat/baturetno",
	}

	utils.StartPrewarmer(endpoints)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
