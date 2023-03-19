package elasticrepo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github/Abraxas-365/bitacora/internal"
	"github/Abraxas-365/bitacora/pkg/report/domain/models"
	"github/Abraxas-365/bitacora/pkg/report/domain/ports"
	"io"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type elastic struct {
	client  *elasticsearch.Client
	index   string
	timeout time.Duration
}

func ElasticRepositoryFactory(client *elasticsearch.Client, index string) ports.ReportSearchRepository {
	return &elastic{
		client:  client,
		index:   index,
		timeout: 10 * time.Second,
	}
}

func (e *elastic) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	req := esapi.DeleteRequest{
		Index:      e.index,
		DocumentID: id,
	}
	resp, err := req.Do(ctx, e.client)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "DeleteRequest.Do")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return internal.NewErrorf(internal.ErrorCodeUnknown, "DeleteRequest.Do %s", resp.StatusCode)
	}

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (e *elastic) Index(report models.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(report); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "json.NewEncoder.Encode")
	}

	req := esapi.IndexRequest{
		Index:      e.index,
		Body:       &buf,
		DocumentID: report.Id.String(),
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, e.client)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "IndexRequest.Do")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return internal.NewErrorf(internal.ErrorCodeUnknown, "IndexRequest.Do %s", resp.StatusCode)
	}

	io.Copy(io.Discard, resp.Body)

	return nil
}

func (e *elastic) Search(args models.SearchCriteria) (models.Reports, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	if args.IsZero() {
		return models.Reports{}, nil
	}

	should := make([]interface{}, 0, 5)

	if args.Title != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"title": *args.Title,
			},
		})
	}

	if args.Status != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"status": *args.Status,
			},
		})
	}

	if args.Tags != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"tags": args.Tags,
			},
		})
	}

	if args.Description != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"description": *args.Description,
			},
		})
	}

	if args.Error != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"error": *args.Error,
			},
		})
	}

	if args.Status != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"status": *args.Status,
			},
		})
	}

	if args.UserNickname != nil {
		should = append(should, map[string]interface{}{
			"match": map[string]interface{}{
				"nickname": *args.UserNickname,
			},
		})
	}

	var query map[string]interface{}

	if len(should) > 1 {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": should,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": should[0],
		}
	}

	query["sort"] = []interface{}{
		"_score",
		// 	map[string]interface{}{"id": "asc"},
	}

	query["from"] = args.From
	query["size"] = args.Size

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err)
		return models.Reports{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "json.NewEncoder.Encode")
	}

	req := esapi.SearchRequest{
		Index: []string{e.index},
		Body:  &buf,
	}

	resp, err := req.Do(ctx, e.client)
	if err != nil {
		fmt.Println(err)
		return models.Reports{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "SearchRequest.Do")
	}
	defer resp.Body.Close()

	if resp.IsError() {
		fmt.Println(resp.String())
		return models.Reports{}, internal.NewErrorf(internal.ErrorCodeUnknown, "SearchRequest.Do %d", resp.StatusCode)
	}

	var hits struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source models.Report `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&hits); err != nil {
		fmt.Println(err)
		return models.Reports{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "json.NewDecoder.Decode")
	}

	res := make(models.Reports, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		res[i].Id = hit.Source.Id
		res[i].Title = hit.Source.Title
		res[i].Tags = hit.Source.Tags
		res[i].Data = hit.Source.Data
		res[i].UserNickname = hit.Source.UserNickname
		res[i].Status = hit.Source.Status
		res[i].Error = hit.Source.Error
		res[i].Solution = hit.Source.Solution
	}

	return res, nil
}
