# RunExam JSON Exam Format

## Overview

RunExam consumes exams defined as a single JSON file with the following schema.

---

## Top-Level Structure

```json
{
  "TestMetadata": { ... },
  "Questions": [ ... ]
}
```

---

### TestMetadata

| Field             | Type     | Required | Description                                     |
|-------------------|----------|----------|-------------------------------------------------|
| TestName          | string   | Yes      | Name/title of the exam                          |
| PassingPercentage | float    | Yes      | Minimum percentage to pass (e.g. 70.0)          |
| ExamDomains       | [string] | Yes      | List of high-level domains or categories        |
| ExamDescription   | string   | No       | (Optional) exam or certification description    |

> **Note:** Fields like `TotalQuestions` or `PassingScore` are ignored by the latest RunExam version, as total questions and score are always calculated dynamically.

---

### Questions

Each entry in the `Questions` array represents one exam question.

| Field         | Type               | Required | Description                                                               |
|---------------|--------------------|----------|---------------------------------------------------------------------------|
| Title         | string             | Yes      | Display title, can include domain or difficulty                           |
| Domain        | string             | Yes      | Category/domain (for future stats/filtering)                              |
| Question      | string             | Yes      | The main question text                                                    |
| Options       | object             | Yes      | Map of option keys (e.g., "A", "B") to answer text                        |
| Type          | string             | Yes      | "One" (single-choice) or "Multi" (multi-select, not yet supported)        |
| CorrectAnswer | [string]           | Yes      | Array of correct option keys (e.g., ["C"])                                |
| Explanation   | string             | Yes      | Detailed explanation for the correct answer                               |

**Example:**

```json
{
  "Title": "Kubernetes Fundamentals (Q1)",
  "Domain": "Kubernetes Fundamentals",
  "Question": "Which resource ensures a specific number of pod replicas?",
  "Options": {
    "A": "DaemonSet",
    "B": "Job",
    "C": "ReplicaSet",
    "D": "StatefulSet"
  },
  "Type": "One",
  "CorrectAnswer": ["C"],
  "Explanation": "ReplicaSet ensures the desired number of pod replicas are always running by replacing failed pods."
}
```

---

### Recommendations

- **Randomization**: RunExam will automatically shuffle questions for each test attempt.
- **Option Keys**: Use capital letters ("A", "B", "C", etc.) for option keys for consistency.
- **Single Choice**: Only single-choice ("One") is supported in the UI currently.

---

### Full Example

See [`_testdata/kcna_simulator_60_questions_full.json`](./_testdata/kcna_simulator_60_questions_full.json) for a real-world example.

---

## License

MIT
