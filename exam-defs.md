# RunExam JSON Exam Format

## Overview

RunExam consumes exams defined in a single JSON file with the following schema.

---

## Top-Level Structure

```json
{
  "TestMetadata": { ... },
  "Questions": [ ... ]
}
```

---

## TestMetadata

| Field             | Type     | Required | Description                                                  |
|-------------------|----------|----------|--------------------------------------------------------------|
| TestName          | string   | Yes      | Name/title of the exam                                       |
| TestDescription   | string   | No       | Optional subtitle or description of the exam                 |
| PassingPercentage | float    | Yes      | Minimum percentage to pass (e.g. 70.0)                       |
| TestVersion       | string   | No       | Version tag for the test                                     |
| TestAuthor        | string   | No       | Name or group who created the test                          |
| TestDate          | string   | No       | Date in YYYY-MM-DD format                                    |
| TestDuration      | int      | Yes      | Duration of the exam in minutes                              |
| ExamDomains       | [string] | Yes      | List of high-level knowledge domains                         |
| ExamDescription   | string   | No       | (Optional) Full description or context of the exam           |

> **Note:** Fields like `TotalQuestions` or `PassingScore` are automatically derived and not required.

---

## Questions

Each entry in the `Questions` array represents one question.

| Field         | Type               | Required | Description                                                               |
|---------------|--------------------|----------|---------------------------------------------------------------------------|
| Title         | string             | Yes      | Short title or identifier for the question                                |
| Domain        | string             | Yes      | Category/domain for grouping and filtering                                |
| Question      | string             | Yes      | The question text                                                         |
| Type          | string             | Yes      | One of: `Single`, `Multiple`, `True`, `False`                             |
| Options       | object             | Conditional | Map of option keys to option text (required for Single/Multiple types)   |
| CorrectAnswer | [string]           | Yes      | Array of correct option keys (e.g., ["B"]) or ["True"] for TF types     |
| Explanation   | string             | Yes      | Explanation shown after test is submitted                                 |

### Supported Types

- `Single`: One correct answer from options (radio buttons)
- `Multiple`: Multiple correct answers from options (checkboxes)
- `True`: Boolean question where correct answer is True
- `False`: Boolean question where correct answer is False

### Example Question

```json
{
  "Title": "Control Plane Basics",
  "Domain": "Cluster Architecture",
  "Question": "Which component is responsible for maintaining node health?",
  "Options": {
    "A": "kubelet",
    "B": "kube-scheduler",
    "C": "etcd",
    "D": "kube-controller-manager"
  },
  "Type": "Single",
  "CorrectAnswer": ["D"],
  "Explanation": "The controller manager monitors and manages cluster state, including node health."
}
```

---

## Recommendations

- Use uppercase letters ("A", "B", "C", etc.) as option keys
- Ensure unique `Title` per question to ease identification
- Keep `Explanation` concise but informative
- `Options` are only required for `Single` and `Multiple` question types
- All other fields must be encoded in valid JSON types

---

## Full Example

See [`exam-data/kubernetes/kcna_simulator_60_questions_full.json`](./exam-data/kubernetes/kcna_simulator_60_questions_full.json) for a real-world example.

---

## License

MIT
