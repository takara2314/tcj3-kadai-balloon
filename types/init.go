package types

type ConfigYaml struct {
	Subjects  []string `yaml:"subjects"`
	Schedules struct {
		A []map[string]string `yaml:"A"`
		B []map[string]string `yaml:"B"`
	} `yaml:"schedules"`
	Timetables map[string]string `yaml:"timetables"`
}
