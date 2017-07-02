package main

import (
	"time"

	"golang.org/x/net/context"

	"fmt"

	"encoding/json"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Tweet is a structure used for serializing/deserializing
// data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		// Error handling
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("Elasticsearch version is: %s", esversion)

	exist, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	if !exist {
		// Create a new index
		createIndex, err := client.CreateIndex("twitter").Do(ctx)
		if err != nil {
			// Error handling
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Printf("[CREATE Error]Create index was not acknowledged")
		}
	}

	// Index(insert) a tweet
	tweet1 := Tweet{User: "sluongng", Message: "I am practicing Golang and Elasticsearch", Retweets: 0}
	put1, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index(insert) a second tweet
	tweet2 := Tweet{User: "sluongng", Message: "It is fun"}
	put2, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("2").
		BodyJson(tweet2).
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Type("tweet").
		Id("1").
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	_, err = client.Flush().
		Index("twitter").
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}

	termQuery := elastic.NewTermQuery("use.keyword", "sluongng")
	searchResult, err := client.Search().
		Index("twitter").
		Query(termQuery).
		Sort("user.keyword", true).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Here we interate through searchResult
	// The use of reflect and item.(Tweet) helps making our operation type safe
	// This ensure the only items we are operation upon are items of type (Tweet)
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// A second approach into iterate through results
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Actual interation
		for _, hit := range searchResult.Hits.Hits {

			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				fmt.Printf("[Type Error]Failed to deserialize tweet with index: %s\n", hit.Index)
			}

			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}

	} else {
		fmt.Print("Found no tweets\n")
	}

	update, err := client.Update().
		Index("twitter").
		Type("tweet").
		Id("1").
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").
			Lang("painless").
			Param("num", 1)).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)

	deleteIndex, err := client.DeleteIndex("twitter").Do(ctx)
	if err != nil {
		// Error handling
		panic(err)
	}
	if !deleteIndex.Acknowledged {
		fmt.Printf("[DELETE Error]Delete index was not acknowledged")
	}
}
