
# cs-job-crawler
Linkedin crawler that finds CS jobs.

## How to Use
Clone the repo and run `go build` in the `cs-job-crawler` directory. Run the resulting executable using:
```
./cs-job-crawler -n [number pages] -b [base url]
```
The base url should be a Linkedin job page with the format `https://www.linkedin.com/jobs/view/[page_id]/`

# Contributing
Feel free to fork the repository and make changes.

## References
Built using:
- [Go-Colly](http://go-colly.org/)
- [Goquery](https://godoc.org/github.com/PuerkitoBio/goquery)
