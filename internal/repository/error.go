package repository

type errKind int

const (
	CtxAddProduct            = "Repository#AddProduct"
	CtxGetProducts           = "Repository#GetProducts"
	CtxGetProductById        = "Repository#GetProductById"
	CtxUpdateProduct         = "Repository#UpdateProduct"
	CtxUpdateProductQuantity = "Repository#UpdateProductQuantity"
	CtxDeleteProduct         = "Repository#DeleteProduct"
)

var (
	ErrDataNotFound       = Error{kind: dataNotFound}
	ErrTableDoesNotExists = Error{kind: tableDoesNotExists}
	ErrUnknown            = Error{kind: unknown}
)

const (
	_ errKind = iota
	dataNotFound
	tableDoesNotExists
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case dataNotFound:
		return "ERROR: [Repository] Data Not Found"
	case tableDoesNotExists:
		return "ERROR: [Repository] Table Does Not Exists"
	default:
		return "ERROR: [Repository] Unknown Error"
	}
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(err error) bool {
	target, ok := err.(*Error)
	if !ok {
		return false
	}
	return target.kind == e.kind
}

func (e *Error) SetError(err error) *Error {
	e.err = err
	return e
}
