namespace go yima.news

struct BaseResp {
    1: string code,
    2: string message,
}

struct SciNews {
	1: string id,
	2: string title,
	3: string type,
	4: string userId,
	5: string link,
	6: string cover,
	7: i64 downloads
}

struct GetNewsByTitleRequest {
    1: required string title,
}

struct GetNewsByTitleResponse {
    1: optional SciNews news,

    255: BaseResp baseResp,
}

struct AddNewsRequest {
    1: required SciNews news,
}

struct AddNewsResponse {
    1: BaseResp baseResp,
}

struct GetAllNewsByUserIdRequest {
    1: string userId
}

struct GetAllNewsByUserIdResponse {
    1: list<SciNews> newsList,

    255: BaseResp baseResp,
}

service YimaNews {
    GetNewsByTitleResponse GetNewsByTitle(GetNewsByTitleRequest request)
    AddNewsResponse AddNews(AddNewsRequest request)
    GetAllNewsByUserIdResponse GetAllNewsByUserId(GetAllNewsByUserIdRequest request)
}
