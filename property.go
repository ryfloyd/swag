package swag

import (
	"fmt"
	"go/ast"
	"log"
)

// getPropertyName returns the string value for the given field if it exists, otherwise it panics.
// allowedValues: array, boolean, integer, null, number, object, string
func getPropertyName(field *ast.Field) (name string, fieldType string)  {
	if astTypeSelectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {

		// Support for time.Time as a structure field
		if "Time" == astTypeSelectorExpr.Sel.Name {
			return "string", "string"
		}

		// Support bson.ObjectId type
		if "ObjectId" == astTypeSelectorExpr.Sel.Name {
			return "string", "string"
		}

		panic("not supported 'astSelectorExpr' yet.")

	} else if astTypeIdent, ok := field.Type.(*ast.Ident); ok {
		name = astTypeIdent.Name

		// When its the int type will transfer to integer which is goswagger supported type
		schemeType:=TransToValidSchemeType(name)
		return schemeType,schemeType


	} else if ptr, ok := field.Type.(*ast.StarExpr); ok {
		// RF hack to derefence struct pointers and see if they are a simple astIdent
		if astTypeSelectorExpr, okk := ptr.X.(*ast.SelectorExpr); okk {

			// Support for time.Time as a structure field
			if "Time" == astTypeSelectorExpr.Sel.Name {
				return "string", "string"
			}

			// Support bson.ObjectId type
			if "ObjectId" == astTypeSelectorExpr.Sel.Name {
				return "string", "string"
			}

			panic("Can't dereference pointer, not supported 'astSelectorExpr' yet.")

		} else if astTypeIdent, okk := ptr.X.(*ast.Ident); okk {
			name = astTypeIdent.Name
			schemeType:=TransToValidSchemeType(name)
			return schemeType,schemeType
		} else {
			panic("Can't dereference pointer, not a simple ast.Ident type")
		}
	} else if _, ok := field.Type.(*ast.MapType); ok { // if map
		//TODO: support map
		return "object", "object"
	} else if astTypeArray, ok := field.Type.(*ast.ArrayType); ok { // if array
		str := fmt.Sprintf("%s", astTypeArray.Elt)
		return "array", str
	} else if _, ok := field.Type.(*ast.StructType); ok { // if struct
		return "object", "object"
	} else {
		log.Fatalf("Something goes wrong: %#v", field.Type)
	}

	return name, fieldType
}
