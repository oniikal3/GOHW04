package todo

func main() {}

type Todo struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Status string `json: "status"`
}
