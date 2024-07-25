package definer

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	JSONSchemaValidationNumberCategory = "JSON Schema Validation Number"
	JSONSchemaValidationStringCategory = "JSON Schema Validation String"
	JSONSchemaValidationObjectCategory = "JSON Schema Validation Object"
	JSONSchemaValidationArrayCategory  = "JSON Schema Validation Array"
	JSONSchemaMetaCategory             = "JSON Schema Meta Category"
)

func (Maximum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: JSONSchemaValidationNumberCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "maximum integer",
			Details: "maximum integer which the specified has to be",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Minimum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: JSONSchemaValidationNumberCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "minimum integer",
			Details: "minimum integer which the specified has to be",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ExclusiveMaximum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: JSONSchemaValidationNumberCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "exclusive maximum",
			Details: "maximum integer but not inclusive maximum",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ExclusiveMinimum) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationNumberCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "exclusive minimum",
			Details: "minimum integer but not inclusive minimum",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MultipleOf) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationNumberCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "multiple of a number",
			Details: "multiple of a number",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MaxLength) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationStringCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "max length",
			Details: "max length of a string",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MinLength) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationStringCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "min length",
			Details: "min length of a string",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (Pattern) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationStringCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "pattern",
			Details: "pattern",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (ContentEncoding) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationStringCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "content-encoding",
			Details: "content-encoding",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (ContentMediatype) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationStringCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "content-mediatype",
			Details: "content-mediatype",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MaxItems) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationArrayCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "max items in array",
			Details: "max items in array",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MinItems) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationArrayCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "min items in array",
			Details: "min items in array",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}


func (UniqueItems) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationArrayCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "decides if the items should be unique",
			Details: "decides if the items should be unique",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MinProperties) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationObjectCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "min properties of the object",
			Details: "min properties of the object",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

func (MaxProperties) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: JSONSchemaValidationObjectCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "max properties of the object",
			Details: "max properties of the object",
		},
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}




func (ID) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: JSONSchemaMetaCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "id of the JSON Schema",
			Details: "id of the JSON Schema used for cross-referencing",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Draft) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: JSONSchemaMetaCategory,
		DetailedHelp: markers.DetailedHelp{
			Summary: "JSON Schema draft to use",
			Details: "specified JSON Schema draf for this type to use",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}
