package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type Feature struct {
	Name     string `json:"name"`
	Progress string `json:"progress"`
	Status   string `json:"status"`
	Tasks    []Task `json:"tasks"`
}

type Skill struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Stats struct {
	TotalSkills    int    `json:"totalSkills"`
	TotalFeatures  int    `json:"totalFeatures"`
	GlobalProgress string `json:"globalProgress"`
}

type DashboardData struct {
	Skills   []Skill   `json:"skills"`
	Features []Feature `json:"features"`
	Stats    Stats     `json:"stats"`
}

func ExportDashboardData(root, outputPath string) error {
	data := DashboardData{
		Skills:   getSkills(root),
		Features: getFeatures(root),
	}

	data.Stats.TotalSkills = len(data.Skills)
	data.Stats.TotalFeatures = len(data.Features)

	if len(data.Features) > 0 {
		totalProgress := 0
		for _, f := range data.Features {
			p := 0
			fmt.Sscanf(f.Progress, "%d%%", &p)
			totalProgress += p
		}
		avg := totalProgress / len(data.Features)
		data.Stats.GlobalProgress = fmt.Sprintf("%d%%", avg)
	} else {
		data.Stats.GlobalProgress = "0%"
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonData, 0644)
}

func getSkills(root string) []Skill {
	var skills []Skill
	files, _ := ioutil.ReadDir(root)
	for _, f := range files {
		if f.IsDir() {
			skillFile := filepath.Join(root, f.Name(), "SKILL.md")
			if _, err := os.Stat(skillFile); err == nil {
				content, _ := ioutil.ReadFile(skillFile)
				versionRegex := regexp.MustCompile(`version: (v?\d+\.\d+\.\d+)`)
				descRegex := regexp.MustCompile(`description: (.*)`)

				version := "v1.0.0"
				if m := versionRegex.FindStringSubmatch(string(content)); len(m) > 1 {
					version = m[1]
				}

				description := fmt.Sprintf("Advanced capability for %s.", f.Name())
				if m := descRegex.FindStringSubmatch(string(content)); len(m) > 1 {
					description = strings.TrimSpace(m[1])
				}

				skills = append(skills, Skill{
					Name:        f.Name(),
					Version:     version,
					Description: description,
					Status:      "Installed",
				})
			}
		}
	}
	return skills
}

func getFeatures(root string) []Feature {
	var features []Feature
	specsDir := filepath.Join(root, ".specs", "features")
	files, _ := ioutil.ReadDir(specsDir)
	for _, f := range files {
		if f.IsDir() {
			tasksFile := filepath.Join(specsDir, f.Name(), "tasks.md")
			var tasks []Task
			progress := "0%"
			status := "Pending"

			if _, err := os.Stat(tasksFile); err == nil {
				content, _ := ioutil.ReadFile(tasksFile)
				taskRegex := regexp.MustCompile(`- \[(.)\] (.*)`)
				matches := taskRegex.FindAllStringSubmatch(string(content), -1)

				doneCount := 0
				for _, m := range matches {
					done := strings.ToLower(m[1]) == "x"
					if done {
						doneCount++
					}
					tasks = append(tasks, Task{
						Name: strings.TrimSpace(m[2]),
						Done: done,
					})
				}

				if len(tasks) > 0 {
					p := (doneCount * 100) / len(tasks)
					progress = fmt.Sprintf("%d%%", p)
					if p == 100 {
						status = "Completed"
					} else if p > 0 {
						status = "In Progress"
					}
				}
			}

			features = append(features, Feature{
				Name:     f.Name(),
				Progress: progress,
				Status:   status,
				Tasks:    tasks,
			})
		}
	}
	return features
}
