package dto

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SKBuiltinStringT struct {
	String *string `json:"string,omitempty"`
}

type SKBuiltinBufferT struct {
	ILen   *uint32 `protobuf:"varint,1,opt,name=iLen" json:"iLen,omitempty"`
	Buffer []byte  `protobuf:"bytes,2,opt,name=buffer" json:"buffer,omitempty"`
}

type SKBuiltinString_S struct {
	ILen   *uint32 `protobuf:"varint,1,opt,name=iLen" json:"iLen,omitempty"`
	Buffer *string `protobuf:"bytes,2,opt,name=buffer" json:"buffer,omitempty"`
}
