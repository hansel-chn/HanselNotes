package elasticsearch

import (
	elastic "github.com/olivere/elastic/v7"
	"math"
	"time"
)

type SensitiveDictService struct {
}

func (item SensitiveDictService) SearchLatestDictAccessTime() (map[int]time.Time, error) {
	result := make(map[int]time.Time, 0)
	// 字典规则最近访问时间
	hitDictIdAgg := elastic.NewTermsAggregation().Field("hitDict.id").Size(math.MaxInt32)
	maxLogTime := elastic.NewMaxAggregation().Field("logTime")
	hitDictIdAgg = hitDictIdAgg.SubAggregation("maxLogTime", maxLogTime)

	indexName := es.CalculateIndex(enum.SENSITIVEFILEDATA, time.Now().AddDate(0, 0, -7), time.Now())
	if len(indexName) == 0 {
		return result, nil
	}
	searchResult, err := es.ElasticClient.Search().
		Index(indexName...).
		Aggregation("hitDictIdAgg", hitDictIdAgg).
		Pretty(true).Size(0).
		Do(getCtx())
	if err != nil {
		log.Zlog.Error(err)
		return result, err
	}

	agg, found := searchResult.Aggregations.Terms("hitDictIdAgg")
	if !found {
		log.Zlog.Error("we should have a terms aggregation called %q", "hitDictIdAgg")
	}

	for _, dictIdBucket := range agg.Buckets {
		dictId := dictIdBucket.Key
		//maxLogTime, found := dictIdBucket.Max("maxLogTime")
		maxLogTime, found := dictIdBucket.MaxBucket("maxLogTime")
		if found {
			parseTime, err := time.Parse(time.RFC3339, maxLogTime.ValueAsString)
			if err != nil {
				log.Zlog.Error(err)
				return nil, err
			}
			result[int(dictId.(float64))] = parseTime
		}
	}
	return result, nil
	//var a struct {
	//	Buckets []struct {
	//		Key        int `json:"key"`
	//		MaxlogTime struct {
	//			Value         float64 `json:"value"`
	//			ValueAsString string  `json:"value_as_string"`
	//		} `json:"maxLogTime"`
	//	} `json:"buckets"`
	//}
	//agg := searchResult.Aggregations["hitDictIdAgg"]
	//err = json.Unmarshal(agg, &a)
	//
	//return result, nil
}
