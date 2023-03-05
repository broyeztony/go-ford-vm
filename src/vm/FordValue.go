package main

type FordValue struct {
	dataType string
	data     interface{}
}

func (fv FordValue) FordValueType() string {
	return fv.dataType
}

func (fv FordValue) AsNumber() int {
	return fv.data.(int)
}

func (fv FordValue) AsString() string {
	return fv.data.(string)
}

func createFordValue(datatype string, data interface{}) FordValue {
	return FordValue{dataType: datatype, data: data}
}

func (fv *FordValue) IsNumber() bool {
	return fv.dataType == "number"
}

func (fv *FordValue) IsString() bool {
	return fv.dataType == "string"
}
