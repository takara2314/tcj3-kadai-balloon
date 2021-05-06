package types

type ConfigYaml struct {
	Subjects  []string `yaml:"subjects"`
	Schedules struct {
		A [][][]string `yaml:"A"`
		B [][][]string `yaml:"B"`
	} `yaml:"schedules"`
	Timetables map[string][]interface{} `yaml:"timetables"`
}
