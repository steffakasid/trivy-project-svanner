package pkg

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aquasecurity/trivy/pkg/report"
	"github.com/xanzy/go-gitlab"
)

var git *gitlab.Client

type trivy struct {
	projName string
	state    string
	ignore   []string
	result   report.Results
}

type trivyResults []trivy

func init() {
	gitToken := os.Getenv("GITLAB_TOKEN")
	if gitToken == "" {
		log.Fatal("No GITLAB_TOKEN env var set!")
	}
	gitHost := os.Getenv("GITLAB_HOST")
	if gitHost == "" {
		gitHost = "https://gitlab.com"
	}

	var err error

	git, err = gitlab.NewClient(gitToken, gitlab.WithBaseURL(gitHost))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func ScanGroup(id, jobName, trivyArtifact string) (trivyResults, error) {
	if id == "" {
		return nil, errors.New("no group id set")
	}

	results := trivyResults{}
	options := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
		Archived:         gitlab.Bool(false),
		IncludeSubgroups: gitlab.Bool(true),
	}
	projs, _, err := git.Groups.ListGroupProjects(id, options)
	if err != nil {
		return nil, err
	}
	fmt.Println()
	for _, proj := range projs {
		fmt.Printf("Scan project %s for trivy results\n", proj.NameWithNamespace)
		projResult := trivy{projName: proj.Name}
		projResult.result, projResult.state, err = getTrivyResult(jobName, trivyArtifact, proj.ID, proj.DefaultBranch)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Result", projResult)
		}
		projResult.ignore, err = getTrivyIgnore(proj.ID, proj.DefaultBranch)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Ignore", projResult.ignore)
		}
		results = append(results, projResult)
		fmt.Println()
	}
	return results, nil
}

func getTrivyResult(jobName, trivyArtifact string, pid int, ref string) (report.Results, string, error) {
	jobs, _, err := git.Jobs.ListProjectJobs(pid, &gitlab.ListJobsOptions{IncludeRetried: *gitlab.Bool(false)})
	if err != nil {
		return nil, "", err
	}

	var state string
	for _, job := range jobs {
		if job.Name == jobName {
			state = job.Status
			break
		}
	}

	rdr, res, err := git.Jobs.DownloadArtifactsFile(pid, ref, &gitlab.DownloadArtifactsFileOptions{Job: gitlab.String(jobName)})
	if err != nil {
		if res.StatusCode == 404 {
			return nil, state, fmt.Errorf("no %s job result", jobName)
		} else {
			return nil, state, err
		}
	}
	unzip, err := zip.NewReader(rdr, rdr.Size())
	if err != nil {
		fmt.Println("Error unzip")
		return nil, state, err
	}

	for _, file := range unzip.File {
		if file.Name == trivyArtifact {
			rc, err := file.Open()

			if err != nil {
				fmt.Println("Error file open")
				return nil, state, err
			}

			bt, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, state, err
			}
			log.Printf("read %d byte", len(bt))
			rc.Close()

			jsonResult := &report.Results{}
			err = json.Unmarshal(bt, jsonResult)
			if err != nil {
				return nil, state, err
			}

			return *jsonResult, state, err
		}
	}
	return nil, state, fmt.Errorf("no %s file found", trivyArtifact)
}

func getTrivyIgnore(pid int, ref string) ([]string, error) {
	bt, res, err := git.RepositoryFiles.GetRawFile(pid, ".trivyignore", &gitlab.GetRawFileOptions{Ref: gitlab.String(ref)})
	if err != nil {
		if res.StatusCode == 404 {
			return nil, errors.New("no .trivyignore file found")
		} else {
			return nil, err
		}
	}
	var ignores []string
	for _, str := range strings.Split(string(bt), "\n") {
		if !strings.HasPrefix(str, "#") {
			ignores = append(ignores, str)
		}
	}
	return ignores, nil
}
