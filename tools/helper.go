package tools

import "tools/utils"

func InitLog(path string) utils.Conf {
	return utils.Init(path)
}
