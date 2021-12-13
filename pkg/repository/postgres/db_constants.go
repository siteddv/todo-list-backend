package repository

//Tables names
const (
	// UsersTable is a name of "users" table in db. Model: model.User
	UsersTable = "users"

	// TodoListTable is a name of "lists" table in db. Model: model.TodoList
	TodoListTable = "todo_lists"

	// UsersListsTable is a name of 3rd table linking "users" and "lists" in db. Model: model.UsersLists
	UsersListsTable = "users_lists"

	// TodoItemsTable is a name of "items" table in db. Model: model.TodoItem
	TodoItemsTable = "todo_items"

	// ListsItemsTable is a name of 3rd table linking "items" and "lists" in db. Model: model.ListsItems
	ListsItemsTable = "lists_items"
)

//Tables columns names
const (
	// Id is a name of "id" field of any model
	Id = "id"

	// Name is a name of "name" field of model.User
	Name = "name"

	// Username is a name of "username" field of model.User
	Username = "username"

	// PasswordHash is a name of "password" field of model.User
	PasswordHash = "password_hash"

	// Title is a name of "title" field of model.TodoItem, model.TodoList
	Title = "title"

	// Description is a name of "description" field of model.TodoItem, model.TodoList
	Description = "description"

	// Done is a name of "done" field of model.TodoItem
	Done = "done"

	// UserId is a name of "user_id" field of model.UsersLists
	UserId = "user_id"

	// ListId is a name of name "list_id" of 3rd table model.UsersLists, model.ListsItems
	ListId = "list_id"

	// ItemId is a name of name "item_id" of 3rd table model.ListsItems
	ItemId = "item_id"
)
