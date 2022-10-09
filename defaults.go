package defaults

func StructOpt(tag ...string) *StructDef {
	if len(tag) != 0 {
		return &StructDef{DefaultTagName: tag[0]}
	}
	return &StructDef{DefaultTagName: "default"}
}
