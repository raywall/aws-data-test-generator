package motor

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var datafile string = "data/data-library.dat"

type RandomUserResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Genero string `json:"gender"`
	Name   struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Location struct {
		Street struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"street"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Postcode    int64  `json:"postcode"`
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinates"`
		Timezone struct {
			Offset      string `json:"offset"`
			Description string `json:"description"`
		}
	} `json:"location"`
	Email string `json:"email"`
	Login struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
		MD5      string `json:"md5"`
		SHA1     string `json:"sha1"`
		SHA256   string `json:"SHA256"`
	} `json:"login"`
	DOB struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	} `json:"dob"`
	Registered struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	} `json:"registered"`
	Phone string `json:"phone"`
	Cell  string `json:"cell"`
	ID    struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"id"`
	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	NAT string `json:"nat"`
}

// Gera dados ficticios de usuário
func FetchRandomUserData(size int) (RandomUserResponse, error) {
	response, err := http.Get(fmt.Sprint("https://randomuser.me/api/?results=", size, "&nat=br"))
	if err != nil {
		return RandomUserResponse{}, err
	}
	defer response.Body.Close()

	var data RandomUserResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return RandomUserResponse{}, err
	}

	return data, nil
}

func (data *RandomUserResponse) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(*data); err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func (data *RandomUserResponse) Deserialize(file io.Reader) {
	decoder := gob.NewDecoder(file)

	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
}

func (data *RandomUserResponse) Save() {
	// Crie um arquivo binário para escrita
	file, err := os.Create(datafile)
	if err != nil {
		log.Fatalln("Falhou ao criar arquivo binário:", err)
	}

	defer file.Close()

	// Escreva os dados da struct no arquivo binário
	err = binary.Write(file, binary.LittleEndian, data.Serialize())
	if err != nil {
		log.Fatalln("Falhou ao criar arquivo com os dados de usuário:", err)
	}
}

func (data *RandomUserResponse) Load() {
	file, err := os.Open(datafile)
	if err != nil {
		log.Fatalln("Falhou ao abrir o arquivo com os dados:", err)
	}

	defer file.Close()

	// Leia o conteúdo do arquivo na struct
	(*data).Deserialize(file)
}

func (data *RandomUserResponse) Get() Result {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	pos := len((*data).Results)
	return (*data).Results[randGenerator.Intn(pos)]
}

func (result *Result) UF() string {
	stateMap := map[string]string{
		"Acre":                "AC",
		"Alagoas":             "AL",
		"Amapá":               "AP",
		"Amazonas":            "AM",
		"Bahia":               "BA",
		"Ceará":               "CE",
		"Distrito Federal":    "DF",
		"Espírito Santo":      "ES",
		"Goiás":               "GO",
		"Maranhão":            "MA",
		"Mato Grosso":         "MT",
		"Mato Grosso do Sul":  "MS",
		"Minas Gerais":        "MG",
		"Pará":                "PA",
		"Paraíba":             "PB",
		"Paraná":              "PR",
		"Pernambuco":          "PE",
		"Piauí":               "PI",
		"Rio de Janeiro":      "RJ",
		"Rio Grande do Norte": "RN",
		"Rio Grande do Sul":   "RS",
		"Rondônia":            "RO",
		"Roraima":             "RR",
		"Santa Catarina":      "SC",
		"São Paulo":           "SP",
		"Sergipe":             "SE",
		"Tocantins":           "TO",
	}

	// Consulte o mapa para obter a sigla do estado
	uf, found := stateMap[(*result).Location.State]
	if !found {
		log.Fatalln("Estado não encontrado:", (*result).Location.State)
	}

	return uf
}
