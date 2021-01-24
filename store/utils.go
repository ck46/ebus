package store

type ItemNotExistError struct{}
type StoreNotExistError struct{}
type ProductNotExistError struct{}
type CategoryNotExistError struct{}
type ImageNotExistError struct{}
type DepartmentNotExistError struct{}

func (*ItemNotExistError) Error() string {
	return "Item does not exist."
}

func (*StoreNotExistError) Error() string {
	return "Store does not exist."
}

func (*ProductNotExistError) Error() string {
	return "Product does not exist."
}

func (*CategoryNotExistError) Error() string {
	return "Category does not exist."
}

func (*ImageNotExistError) Error() string {
	return "Image does not exist."
}

func (*DepartmentNotExistError) Error() string {
	return "Department does not exist."
}
