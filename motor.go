package main

import (
	"aws-data-test-generator/motor"
	"sort"
	"strings"
	"sync"
)

var (
	wg       sync.WaitGroup
	columns  map[string][]string
	userData motor.RandomUserResponse
)

func (request *Request) Create() (Response, error) {
	req := (*request)
	res := Response{}

	columns = make(map[string][]string)

	userData = motor.RandomUserResponse{}
	userData.Load()

	sort.Slice(req.DataStruct, func(i, j int) bool {
		return req.DataStruct[i].ID < req.DataStruct[j].ID
	})

	wg.Add(len(req.DataStruct))

	for i, _ := range req.DataStruct {
		go req.DataStruct[i].generate(req.Size)
	}

	wg.Wait()

	if req.Headers {
		headers := []string{}

		for _, col := range req.DataStruct {
			headers = append(headers, col.Header)
		}

		res.Add(strings.Join(headers, ";"))
	}

	for i := 0; i < req.Size; i++ {
		line := []string{}

		for _, col := range req.DataStruct {
			line = append(line, columns[col.Type][i])
		}

		res.Add(strings.Join(line, ";"))
	}

	return res, nil
}

func (column *Column) generate(size int) {
	col := *column

	defer wg.Done()

	switch col.Type {
	case "code":
		values := []string{}

		for i := 0; i < size; i++ {
			values = append(values, motor.GenerateRandomNumericCode(col.Size))
		}

		columns[col.Type] = values
		break

	case "currency":
		values := []string{}

		for i := 0; i < size; i++ {
			value, _ := motor.GenerateRandomCurrencyValue(col.MaxValue, col.MinValue)
			values = append(values, value)
		}

		columns[col.Type] = values
		break

	case "first-name":
		values := []string{}

		for i := 0; i < size; i++ {
			values = append(values, userData.Get().Name.First)
		}

		columns[col.Type] = values
		break

	case "last-name":
		values := []string{}

		for i := 0; i < size; i++ {
			values = append(values, userData.Get().Name.Last)
		}

		columns[col.Type] = values
		break

	case "state":
		values := []string{}

		for i := 0; i < size; i++ {
			user := userData.Get()
			values = append(values, user.UF())
		}

		columns[col.Type] = values
		break

	case "city":
		values := []string{}

		for i := 0; i < size; i++ {
			values = append(values, userData.Get().Location.City)
		}

		columns[col.Type] = values
		break

	case "uuid":
		values := []string{}

		for i := 0; i < size; i++ {
			values = append(values, motor.GenerateRandomUUIDv4())
		}

		columns[col.Type] = values
		break
	}
}
