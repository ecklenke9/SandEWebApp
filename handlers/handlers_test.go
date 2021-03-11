package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/ecklenke9/SandEWebApp/internal/todo"
	"github.com/gin-gonic/gin"
)

func TestDeleteTodoHandler(t *testing.T) {
	type args struct {
		payload interface{}
		id      string
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successfully delete todo by id",
			args: args{
				payload: "",
				statusCode: 200,
			},
		},
		{
			name: "failed to delete todo by id",
			args: args{
				payload: "",
				id:      "bad test",
				statusCode: 404,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			con, _ := gin.CreateTestContext(rr)

			if tt.name == "successfully delete todo by id" {
				foundID := todo.Add("Yo! What up! ...")

				// attach path variable with a 'key' of 'id' and "value" of "foundID"
				con.Params = append(con.Params, gin.Param{Key: "id", Value: foundID})
				t.Log(con.Params)
				DeleteTodoHandler(con)

				// If response does not contain a status code of 200, test fails.
				if rr.Code != 200 {
					t.Fail()
				}

				// If unable to find by the foundID, the test is successful as the TODO has been
				// successfully deleted.
				// Else the test fails because the TODO has not been deleted.
				if _, err := todo.GetByID(foundID); err != nil {
					t.Log("Successfully deleted: ", foundID)
				} else {
					t.Fail()
				}
			}

			if tt.name == "failed to delete todo by id" {
				foundID := todo.Add("Test todo...")

				con.Params = append(con.Params, gin.Param{Key: "id", Value: tt.args.id})
				t.Log(con.Params)
				DeleteTodoHandler(con)

				// If response does not contain a status code of 404, test fails.
				t.Log("badRecorder.code: ", rr.Result().StatusCode)
				if rr.Code != 404 {
					t.Fail()
				}

				if _, err := todo.GetByID(foundID); err == nil {
					t.Log("Successful test")
				} else {
					t.Fail()
				}
			}
		})
	}
}
