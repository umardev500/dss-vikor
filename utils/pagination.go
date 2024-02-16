package utils

import "github.com/umardev500/spk/domain/model"

func BuildPageInfo(info *model.PageInfo) {

	if info.PerPage == 0 {
		info.PerPage = 14
	}

	if info.Offset <= 0 {
		info.Offset = 0
	}

	if info.Offset > 0 {
		info.Offset = (info.Offset - 1) * info.PerPage
	}

	if info.SortInfo.SortStrategy == "" {
		info.SortInfo.SortStrategy = model.ASC
	}
}
