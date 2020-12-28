package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/util/m"
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/util/must"
)

const date = "20201223"
const orderBy = "desc"

//const api = "/app/common/app-init"
const api = "/app/post/info"

const timestamp = 1608689786420
const size = 1000

func main() {
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://10.70.11.217:9200"),
		elastic.SetBasicAuth("sp", "gaeamobile"))
	must.Must(err)
	q := m.M{
		"_source": m.M{
			"includes": m.L{"client_addr", "log_id", "request_time", "status", "uri", "created_at"},
		},
		"query": m.M{
			"bool": m.M{
				"must": m.M{"match_all": m.M{}},
				"filter": m.L{
					m.M{
						"term": m.M{
							"item_id": "yingdi_api",
						},
					},
					m.M{
						"term": m.M{
							"status": 400,
						},
					},
					m.M{
						"range": m.M{
							"created_at": m.M{"gte": timestamp},
						},
					},
				},
			},
		},
		"sort": m.L{m.M{"created_at": orderBy}},
		"size": size,
	}
	r, err := client.Search().Index("spnginxlog-" + date).Source(q).Do(context.Background())
	must.Must(err)

	fmt.Println(r.TotalHits())
	for _, v := range r.Hits.Hits {
		fmt.Println(string(v.Source))
		//j := simplejson.MustJSON(v.Source)
		//ts := j.Get("created_at").Int()
		//fmt.Println(ts)
		//j.Set("created_at", time.Unix(0, int64(ts*1000)).Format("2006-01-02 15:04:05"))
		//fmt.Println(j)
	}

	//client, err = elastic.NewSimpleClient(elastic.SetURL("http://10.70.11.217:9200"),
	//	elastic.SetBasicAuth("sp", "gaeamobile"))
	//must.Must(err)
	//q = m.M{
	//	"query": m.M{
	//		"bool": m.M{
	//			"must": m.M{"match_all": m.M{}},
	//			"filter": m.L{
	//				m.M{
	//					"term": m.M{
	//						"log_path": 1,
	//					},
	//				},
	//				m.M{
	//					"term": m.M{
	//						"env": "product",
	//					},
	//				},
	//				m.M{
	//					"term": m.M{
	//						"log_action": api,
	//					},
	//				},
	//				//m.M{
	//				//	"range": m.M{
	//				//		"created_at": m.M{"gte": timestamp},
	//				//	},
	//				//},
	//			},
	//		},
	//	},
	//	"sort": m.L{m.M{"created_at": orderBy}},
	//	"size": size,
	//}
	//r, err = client.Search().Index("ydlog-" + date).Source(q).Do(context.Background())
	//must.Must(err)
	//fmt.Println(r.TotalHits())
	//for _, v := range r.Hits.Hits {
	//	fmt.Println(string(v.Source))
	//}
}
