package defaults

import "defaults/structs"

func StructOpt(tag ...string) *structs.StructDef {
	if len(tag) != 0 {
		return &structs.StructDef{DefaultTagName: tag[0]}
	}
	return &structs.StructDef{DefaultTagName: "default"}
}
