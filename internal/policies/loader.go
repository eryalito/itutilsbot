package policies

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/open-policy-agent/opa/rego"
	"gopkg.in/yaml.v3"
)

type Authconfig struct {
	Modules []struct {
		Name string `yaml:"name"`
		File string `yaml:"file"`
	} `yaml:"modules"`
	Queries []rego.PreparedEvalQuery
	ConfigPath string
}

func (c *Authconfig) GetConfg(ConfigPath string) *Authconfig {
	c.ConfigPath = ConfigPath
	yamlFile, err := ioutil.ReadFile(c.ConfigPath + "/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func (c *Authconfig) Init() {
	for _, m2 := range c.Modules {
		file, err := ioutil.ReadFile(c.ConfigPath + "/" + m2.File)
		if err != nil {
			log.Fatalf("yamlFile.Get err   #%v ", err)
		}
		contents := string(file)
		ctx := context.Background()
		query, err := rego.New(rego.Module((m2.File), contents), rego.Query(fmt.Sprintf("allowed = data.%s.allow; error_message = data.%s.error_message", m2.Name, m2.Name))).PrepareForEval(ctx)
		if err != nil {
			log.Fatalf("Error processing rego file %v", err)
		}
		c.Queries = append(c.Queries, query)
	}
}

func (c *Authconfig) Validate(input interface{}) (string, bool) {
	message := "Not allowed"
	for _, q := range c.Queries {
		result, err := q.Eval(context.Background(), rego.EvalInput(input))
		if err != nil {
			log.Fatalf("Error evaluating rego file %v", err)
		}
		if len(result) > 0 {
			msg := result[0].Bindings.WithoutWildcards()["error_message"]
			if val, ok := msg.(string); ok {
				message = val
			}
			return message, result[0].Bindings.WithoutWildcards()["allowed"] == true
		}
	}
	return message, false
}
