package expenses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *handler) CreateExpensesHandler(c echo.Context) error {
	e := Expenses{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := h.DB.QueryRow(` INSERT INTO expenses (title,amount,note,tags) 
	values ($1,$2,$3,$4) RETURNING id `,
		e.Title, e.Amount, e.Note, pq.Array(e.Tags))

	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't create expenses please contact admin:" + err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}
