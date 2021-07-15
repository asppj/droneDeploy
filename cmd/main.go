package main

import "github.com/kataras/iris/v12"

type (
	request struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	response struct {
		ID      uint64 `json:"id"`
		Message string `json:"message"`
	}
)

func main() {
	app := iris.New()
	app.Handle("GET", "/users/{id:uint64}", updateUser)
	app.Listen(":8080")
}

func updateUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint64("id")

	var req request
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.StopExecution()
		// ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp := response{
		ID:      id,
		Message: req.Firstname + " GET successfully",
	}
	ctx.JSON(resp)
}
