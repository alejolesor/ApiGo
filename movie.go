package main

//Movie ...
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

//Movies array
type Movies []Movie

//Errorctrl ...
type Errorctrl struct {
	IDError     int    `json:"IDError"`
	Description string `json:"Descripcion"`
}

//Errores ...
type Errores []Errorctrl

//ResponseGeneral ...
type ResponseGeneral struct{
	Code 	int     `json:"code"`
	Status 	string  `json:"status"`
	Data    Data    `json:"data"`            

}
//ResponseGen ...
type ResponseGen struct{
	ResponseGen []ResponseGeneral
}
//Data ...
type Data struct{
	_Id     int                   `json:"_id"`
	Total 	int              `json:"total"`
	Results []ResultsComics  `json:"results"`
}
//ResultsComics ...
type ResultsComics struct{
	IDCommic    int      `json:"id"`
	Tittle      string   `json:"title"`
	Description string   `json:"description"`
	Modified    string   `json:"modified"`
}

//Comicss ...
type Comicss struct {
	Comics []ResultsComics
}


