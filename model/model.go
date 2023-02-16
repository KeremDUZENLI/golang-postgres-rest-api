package model

import (
	"validation/configs"
)

type MyStruct struct {
	Id        int    `json:"id"`
	Nome      string `json:"name"`
	SobreNome string `json:"surname"`
	Ano       int    `json:"age"`
}

var newList = []MyStruct{
	{
		Id:        1,
		Nome:      "XXX",
		SobreNome: "YYY",
		Ano:       10,
	},
	{
		Id:        2,
		Nome:      "ZZZ",
		SobreNome: "TTT",
		Ano:       20,
	},
	{
		Id:        3,
		Nome:      "QQQ",
		SobreNome: "RRR",
		Ano:       30,
	},
}

func BeginDatabaseWithInfos() *[]MyStruct {
	configs.Database.Create(&newList)

	return &newList
}

func (my *MyStruct) CreateDatabaseInfo() *MyStruct {
	configs.Database.Create(&my)

	return my
}

func ReadDatabaseInfos() *[]MyStruct {
	var infos []MyStruct
	configs.Database.Find(&infos)

	return &infos
}

func (my *MyStruct) UpdateDatabaseInfo() *MyStruct {
	configs.Database.Save(&my)

	return my
}

func DeleteDatabaseInfo(id int) *MyStruct {
	var info MyStruct
	configs.Database.Where("id = ?", id).Delete(&info)

	return &info
}
