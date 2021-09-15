package pullrequest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"

	"ghub/pkg/rest"
)

const URL = "https://api.github.com/repos/%s/%s/pulls?state=open&sort=popularity&direction=asc&per_page=2&page=%s"

var relRegex = regexp.MustCompile(`&page=(\d+)>; rel="last"`)
var lastPageRegex = regexp.MustCompile(`\d+`)

type result struct {
	prs []PullRequest
	res *http.Response
}

func GetOpen(w http.ResponseWriter, r *http.Request) {
	log.Println("called get open pull request")

	var repoPRs []PullRequest

	vars := mux.Vars(r)
	user := vars["user"]
	repo := vars["repository"]

	log.Println("user: ", user)
	log.Println("repo: ", repo)

	resp := get(w, user, repo, 1)

	err := parse(resp, &repoPRs)

	if err != nil {
		log.Fatalln("Get error: ", err)
		rest.ErrorResp(w, http.StatusInternalServerError, "unable to parse github response")
	}

	last, err := lastPage(resp)

	if err != nil {
		log.Fatalln("lastPage error: ", err)
		rest.ErrorResp(w, http.StatusBadRequest, "unable to parse last page")
	}

	if last > 0 {
		ch := make(chan result)
		go getPaginated(ch, w, user, repo, last)

		for {
			v, ok := <-ch
			if ok == false {
				break
			}
			repoPRs = append(repoPRs, v.prs...)
		}
	}

	rest.JSONresp(w, http.StatusOK, repoPRs)
}

func getPaginated(ch chan result, w http.ResponseWriter, user string, repo string, last int) error {
	defer close(ch)

	for i := 2; i <= last; i++ {
		var res result
		var prs []PullRequest

		r := get(w, user, repo, i)

		res.res = r
		err := parse(r, &prs)

		if err != nil {
			log.Fatalln("error: ", err)
			return err
		}

		res.prs = prs

		ch <- res
	}

	return nil
}

func get(w http.ResponseWriter, user string, repo string, page int) *http.Response {
	parsedUrl, err := rest.NewURL(URL, user, repo, page)
	log.Println("parsedUrl: ", parsedUrl)

	if err != nil {
		log.Fatalln("NewURL error: ", err)
		rest.ErrorResp(w, http.StatusInternalServerError, "unable to build github URL")
	}

	resp, err := rest.Get(parsedUrl)

	if err != nil {
		log.Fatalln("Get error: ", err)
		rest.ErrorResp(w, http.StatusInternalServerError, "unable to fetch user repositories")
	}

	return resp
}

func parse(resp *http.Response, reposData *[]PullRequest) error {
	v, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(v, reposData)
}

func lastPage(r *http.Response) (int, error) {
	link := r.Header.Get("Link")

	rel := relRegex.FindString(link)
	lastPageStr := lastPageRegex.FindString(rel)

	if lastPageStr == "" {
		return 0, nil
	}

	lastPage, err := strconv.Atoi(lastPageStr)

	if err != nil {
		return 0, err
	}

	log.Println("lastPage: ", lastPage)

	return lastPage, nil
}
