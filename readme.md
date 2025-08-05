#
# Working in progress 
#

# RunExam

[![Docker Hub](https://img.shields.io/badge/docker-fpinna/runexam-blue?logo=docker)](https://hub.docker.com/r/fpinna/runexam)

**RunExam** is a modern self-hosted exam simulator for certification practice, with instant correction, explanations, PDF export, and ~~beautiful~~ UI.  
It works with custom JSON files defining your own set of questions.

*Disclaimer*

LLMs generated all files inside exam-data. I've been using them to achieve my certification objectives. 

---

## Features

- Random question order every run
- ~~Beautiful and responsive UI (Bootstrap)~~
- Per-question explanations and answer feedback
- Pass/fail result with minimum passing percentage
- Instant correction and review
- PDF export matching your exam report
- Easily pluggable JSON format for new exams
- [Exam json scheme definitions](exam-defs.md) - json schema


---

## Getting Started

### 1. Build Locally

```bash
git clone https://github.com/fpinna/runexam.git
cd runexam
go build -o runexam main.go
```

```
Run (example using a KCNA mock exam)
./runexam exam-data/kubernetes/kcna_simulator_60_questions_full.json
```

- Default port: 9171

Options:

-p, --port 
    Set custom port (default: 9171)

-l, --listen 
    Listen address (default: 0.0.0.0)

Example:

./runexam -p 8181 myexam.json

Now access: http://localhost:9171


# How-to use LLMs to help you generate questions data files for free?





# Run with Docker - working in progress
The easiest way:


```
docker run -it --rm -p 9171:9171 -v $PWD/myexam.json:/app/exam.json fpinna/runexam /app/exam.json
```

-p: Exposes port 9171 \
-v: Mounts your exam JSON file inside the container as /app/exam.json \
You can set a different port using -p and the flag --port inside the command.

Example with custom port
```
docker run -it --rm -p 8181:8181 -v $PWD/myexam.json:/app/exam.json fpinna/runexam /app/exam.json --port 8181
```

# inspired by https://github.com/thiago4go/kubernetes-security-kcsa-mock