package util

type Pagination struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func GetPaginationData(pageSize *int32, pageID *int32) Pagination {
	defaultPageSize := int32(5)
	defaultPageID := int32(1)

	var limit, offset int32

	if pageSize == nil || *pageSize < 1 || *pageSize > 100 {
		limit = defaultPageSize
	} else {
		limit = *pageSize
	}

	if pageID == nil || *pageID < 1 {
		offset = (defaultPageID - 1) * limit
	} else {
		offset = (*pageID - 1) * limit
	}

	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
