package application

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-qiu/rrs-web-server/pkg/controllers"
	"github.com/go-qiu/rrs-web-server/pkg/http/handlers"
	"github.com/go-qiu/rrs-web-server/pkg/http/routers"

	"github.com/gorilla/mux"
)

// AppControllers struct to hold all the pointers to respective controllers.
type AppControllers struct {
	CRUD AppCRUDControllers
	Auth *controllers.AuthCtl
}

// AppCRUDControllers struct to hold all the pointers to the respective CRUD controllers.
type AppCRUDControllers struct {
	Users     *controllers.CRUD
	Vouchers  *controllers.CRUD
	Merchants *controllers.CRUD
}

// Application struct
type Application struct {
	Controllers AppControllers
	// Router      *mux.Router
}

// New will instantiate an application instance.
func New() *Application {
	return &Application{}
}

// Router will instantiate a router instance, with all the routes and controllers mapping.
func (a *Application) Router() *mux.Router {
	// r := routers.New(a.Controllers)

	// instantiate a gorilla/mux reouter.
	r := mux.NewRouter()
	PUBLIC := os.Getenv("PUBLIC")

	r.HandleFunc("/", handlers.ServeHtmlIndex)
	r.HandleFunc("/login", handlers.ServeHtmlLogin)
	r.HandleFunc("/logout", handlers.ServeHtmlLogout)

	// jwt auth api routes
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	routers.RegisterAuthRouter(apiRouter, a.Controllers.Auth)

	// users routes
	r.HandleFunc("/users", handlers.ServeHtmlIndexUsers)
	r.HandleFunc("/users/registration", handlers.ServeHtmlLogin)
	r.HandleFunc("/users/{id}", handlers.ServeHtmlUserProfile)
	r.HandleFunc("/users/{id}/transactions", handlers.ServeHtmlUserRecyclableTransactions)
	r.HandleFunc("/users/{id}/vouchers", handlers.ServeHtmlUserVouchers)
	r.HandleFunc("/users/{id}/points_to_vouchers/redepmtion", handlers.ServeHtmlUserPointsToVouchers)

	// vouchers routes
	r.HandleFunc("/vouchers", handlers.ServeHtmlIndexVouchers)

	// merchants routes
	r.HandleFunc("/merchants", handlers.ServeHtmlIndexMerchants)
	r.HandleFunc("/merchants/{id}/vouchers/aquire", handlers.ServeHtmlMerchantVouchersAquisition)
	r.HandleFunc("/merchants/{id}/vouchers/capture", handlers.ServeHtmlMerchantVoucherCapture)

	// static web pages or assets router
	fp := http.FileServer(http.Dir(PUBLIC))
	r.PathPrefix("/public").Handler(http.StripPrefix("/public/", fp))

	return r
}

// PullDataIntoDataStore will attempt to connect to the specific microservice to retrieve a set of data points and cache them into the in-memory data store for fast access.
func (a *Application) PullDataIntoDataStore() {

	// exception handling (exit fast)

	// ok.
	type dataPoint struct {
		UserID    float64
		Phone     string
		Name      string
		Password  string
		Points    float64
		LastLogin string
	}

	// respBody struct for unmarshalling response body json.
	type respBody struct {
		Ok   bool                   `json:"ok"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
	}

	// http client to connect to users microservice.
	// setup the client to bypass the ssl verification check so that a call to users microservice (via https, protected by self-signed ssl cert) can be done.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// get environment variables for connecting to user microservice.
	API_ROOT_URL := os.Getenv("API_URL_USERS")
	API_KEY := os.Getenv("API_KEY_USERS")
	API_USERNAME := os.Getenv("API_USERNAME_USERS")

	DATA_PTS_PER_PAGE, err := strconv.Atoi(os.Getenv("DATA_PTS_PER_PAGE"))
	if err != nil {
		log.Fatal(err)
	}

	// loop to request for all users data (by page index)
	pageIndex := 0
	dataPtsCount := 0
	isOk := false

	// slice to cache the all the users data points retrieved from users microservice.
	dataPoints := []dataPoint{}

	// !!!! code within this loop could potentially be made to retrieve all data points concurrently from the microservice.
	for (pageIndex == 0) || (isOk && dataPtsCount == DATA_PTS_PER_PAGE) {
		// execute while above condition is true

		// set the endpoint query string
		endpoint := fmt.Sprintf(`%s/getusers?page_index=%d&records_per_page=%d`, API_ROOT_URL, pageIndex, DATA_PTS_PER_PAGE)

		// prepare the GET request.
		apiReq, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			log.Fatal(err)
		}

		// set all the required header attributes of this GET request.
		apiReq.Header.Set("Content-Type", "application/json")
		apiReq.Header.Set("apiKey", API_KEY)
		apiReq.Header.Set("username", API_USERNAME)

		// send out the request.
		outcome, err := client.Do(apiReq)
		if err != nil {
			log.Fatal(err)
		}

		// handle the response body.
		defer outcome.Body.Close()
		body, err := ioutil.ReadAll(outcome.Body)
		if err != nil {
			log.Fatal(err)
		}

		// unmarshal response body json to struct.
		var rb respBody
		err = json.Unmarshal(body, &rb)
		if err != nil {
			log.Fatal(err)
		}
		isOk = rb.Ok

		// reset the data points count (for page index)
		dataPtsCount = 0

		// loop through the data maps returned.
		for _, data := range rb.Data {

			// break down the first level map.

			dp := dataPoint{}

			// loop through the second level map.
			// get the attribute and its value.
			// build the user data point struct.
			for k, v := range data.(map[string]interface{}) {

				switch k {
				case "UserID":
					dp.UserID = v.(float64)
				case "Phone":
					dp.Phone = v.(string)
				case "Name":
					dp.Name = v.(string)
				case "Password":
					dp.Password = v.(string)
				case "Points":
					dp.Points = v.(float64)
				case "LastLogin":
					dp.LastLogin = v.(string)
				}
				//
			}
			dataPoints = append(dataPoints, dp)

			// check the number of data points received in the current page index.
			dataPtsCount++
			//
		}

		// increment the page index count
		pageIndex++
	}

	// add into in-memory data store.
	fmt.Println(dataPoints)
	//
}
