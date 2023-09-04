package main

import (
	"E-Commerce_BE/api"
	"E-Commerce_BE/config/cors"
	DB "E-Commerce_BE/config/db"
	"E-Commerce_BE/model"
	"E-Commerce_BE/util"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	credentialsFile = "./driver.json"
	tokenFile       = "./token.json"
	driveFolder     = "ecom"
)

func main() {
	e := echo.New()

	// Validate input model
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	// ---------------------------------------

	// load file .env
	util.LoadEnv()
	// ---------------------------------------

	// connect db
	db, err := DB.ConnectDB()
	// ---------------------------------------

	// config cors
	cors.SetupCORS(e)
	// ---------------------------------------

	er := db.AutoMigrate(
		model.User{},
		//model.Role{},
		model.Product{},
		model.Category{},
		model.Order{},
	)

	if er != nil {
		return
	}

	if err != nil {
		fmt.Println("Error connect db: ", err)
		return
	} else {
		fmt.Println("connect db success!")
	}

	// init routes
	api.InitRoutes(e, db)
	// ---------------------------------------

	// ---------------------------------------
	// test upload
	//e.POST("/upload", uploadHandler)

	e.Logger.Fatal(e.Start(":5000"))
}

//
//func uploadHandler(c echo.Context) error {
//	// Parse form data
//	form, err := c.MultipartForm()
//
//	if err != nil {
//		return err
//		fmt.Println("err bind data")
//	}
//
//	// Configure OAuth2
//	config, err := getConfig()
//	if err != nil {
//		fmt.Println("err config: ", err)
//		return err
//
//	}
//
//	fmt.Println("err 123: ", form)
//
//	// Create Drive client
//	client, err := getClient(config)
//	if err != nil {
//		return err
//		fmt.Println("err getClinet")
//	}
//
//	// Create folder if not exists
//	folderID, err := createFolder(client, driveFolder)
//	if err != nil {
//		return err
//		fmt.Println("err create folder")
//	}
//
//	// Upload files to Drive
//	var links []string
//	files := form.File["files"]
//	for _, file := range files {
//		src, err := file.Open()
//		if err != nil {
//			return err
//			fmt.Println("err open")
//		}
//		defer src.Close()
//
//		filePath := filepath.Join(folderID, file.Filename)
//		dst, err := client.Files.Create(&drive.File{Name: filePath}).Media(src).Do()
//		if err != nil {
//			return err
//		}
//
//		link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", dst.Id)
//		links = append(links, link)
//	}
//
//	response := struct {
//		Message string   `json:"message"`
//		Links   []string `json:"links"`
//	}{
//		Message: "Files uploaded successfully",
//		Links:   links,
//	}
//
//	return c.JSON(http.StatusOK, response)
//}
//
//func getConfig() (*oauth2.Config, error) {
//	credentials, err := os.ReadFile(credentialsFile)
//	if err != nil {
//		return nil, err
//	}
//
//	config, err := google.ConfigFromJSON(credentials, drive.DriveScope)
//	if err != nil {
//		return nil, err
//	}
//
//	return config, nil
//}
//
//func getClient(config *oauth2.Config) (*drive.Service, error) {
//	tok, err := tokenFromFile(tokenFile)
//	if err != nil {
//		tok, err = getTokenFromWeb(config)
//		if err != nil {
//			return nil, err
//		}
//		saveToken(tokenFile, tok)
//	}
//	client := config.Client(context.Background(), tok)
//	srv, err := drive.New(client)
//	if err != nil {
//		return nil, err
//	}
//	return srv, nil
//}
//
//func tokenFromFile(file string) (*oauth2.Token, error) {
//	f, err := os.Open(file)
//	if err != nil {
//		return nil, err
//
//	}
//	defer f.Close()
//	tok := &oauth2.Token{}
//	err = json.NewDecoder(f).Decode(tok)
//
//	return tok, err
//
//}
//
////func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
////	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
////	fmt.Printf("Go to the following link in your browser then type the "+
////		"authorization code: \n%v\n", authURL)
////
////	var code string
////	if _, err := fmt.Scan(&code); err != nil {
////
////		return nil, err
////	}
////
////	tok, err := config.Exchange(context.TODO(), code)
////	if err != nil {
////		return nil, err
////
////	}
////	return tok, nil
////}
//
//func saveToken(file string, token *oauth2.Token) {
//	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
//	if err != nil {
//		log.Fatalf("Unable to cache oauth token: %v", err)
//	}
//	defer f.Close()
//	json.NewEncoder(f).Encode(token)
//}
//
//func createFolder(srv *drive.Service, folderName string) (string, error) {
//	folder := &drive.File{
//		Name:     folderName,
//		MimeType: "application/vnd.google-apps.folder",
//	}
//	file, err := srv.Files.Create(folder).Do()
//	if err != nil {
//		return "", err
//	}
//	return file.Id, nil
//}
