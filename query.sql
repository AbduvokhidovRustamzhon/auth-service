type User  struct {
		Id              int
		Name            string
		Login           string
		Password        string
		Role            bool
		Removed         bool
}






SELECT * FROM books WHERE removed = FALSE;

SELECT * FROM books where id = 1;


SELECT * FROM books where id = 4 and removed = false;