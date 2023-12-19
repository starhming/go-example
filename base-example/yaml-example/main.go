package yaml_example

import (
	"io"

	"gopkg.in/yaml.v3"
)

func main() {
	//sourceYaml, err := os.ReadFile("input.yaml")
	//if err != nil {
	//	panic(err)
	//}

}

func fetchYaml(sourceYaml []byte) (*yaml.Node, error) {
	rootNode := yaml.Node{}
	err := yaml.Unmarshal(sourceYaml, &rootNode)
	if err != nil {
		return nil, err
	}
	return &rootNode, nil
}

func streamYaml(writer io.Writer, indent *int, in *yaml.Node) error {
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(*indent)
	err := encoder.Encode(in)
	if err != nil {
		return err
	}
	return encoder.Close()
}
